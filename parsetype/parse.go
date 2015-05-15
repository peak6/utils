package parsetype

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

type Parser struct {
	Packages map[string]*Package
	Types    map[string]*Type
}

type Package struct {
	Package *ast.Package
	Files   map[string]*File
}

type File struct {
	File      *ast.File
	Imports   map[string]string
	Alias     map[string]string
	IsBuiltIn bool
}

type Type struct {
	Name        string
	Type        string
	Doc         *ast.CommentGroup
	Comment     *ast.CommentGroup
	FromPackage string
	Properties  map[string]*Type
	ArrayType   *Type
	MapKey      string
	MapType     *Type
	Tags        map[string]string
	RefType     *Type
	RefPackage  string
	RefTypeName string
	IsPointer   bool
}

func NewParser() *Parser {
	return &Parser{
		Packages: make(map[string]*Package),
		Types:    make(map[string]*Type),
	}
}

func (p *Parser) ParseDir(dir string) {
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	p.parse(dir, packages)
}

func (p *Parser) trimGoPath(s string) string {
	paths := os.Getenv("GOPATH")
	goPaths := strings.Split(paths, ":")
	for i := 0; i < len(goPaths); i++ {
		if strings.Contains(s, goPaths[i]) {
			return strings.TrimPrefix(s, path.Join(goPaths[i], "src")+"/")
		}
	}

	return s
}

func (p *Parser) parseFromGoPath(dir string) {
	paths := os.Getenv("GOPATH")
	goPaths := strings.Split(paths, ":")
	for i := 0; i < len(goPaths); i++ {
		p.ParseDir(path.Join(goPaths[i], "src", dir))
	}
}

func (p *Parser) parse(dir string, packages map[string]*ast.Package) {
	for _, astPackage := range packages {
		if p.Packages[dir] != nil {
			return
		}

		pack := &Package{
			Files:   make(map[string]*File),
			Package: astPackage,
		}

		p.Packages[dir] = pack

		for _, astFile := range astPackage.Files {
			file := &File{
				File:    astFile,
				Imports: make(map[string]string),
				Alias:   make(map[string]string),
			}

			pack.Files[astFile.Name.String()] = file

			for _, astDeclaration := range astFile.Decls {
				if generalDeclaration, ok := astDeclaration.(*ast.GenDecl); ok {
					for _, astSpec := range generalDeclaration.Specs {

						if typeSpec, ok := astSpec.(*ast.ImportSpec); ok {
							importPath := strings.Replace(typeSpec.Path.Value, "\"", "", -1)
							packName := importPath
							packNameIndex := strings.LastIndex(importPath, "/")
							if packNameIndex > 0 {
								packName = importPath[packNameIndex+1:]
							}

							if typeSpec.Name != nil {
								file.Alias[typeSpec.Name.String()] = packName
							}

							file.Imports[packName] = importPath
						}

						if typeSpec, ok := astSpec.(*ast.TypeSpec); ok {
							p.parseTypeSpec(p.trimGoPath(dir), file, typeSpec)
						}
					}

				}
			}
		}
	}
}

func (p *Parser) parseTypeSpec(packname string, file *File, astType *ast.TypeSpec) *Type {
	name := packname + "." + astType.Name.String()
	if _, ok := p.Types[name]; ok {
		return p.Types[name]
	}

	t := &Type{}
	t.Name = name
	t.Comment = astType.Comment
	t.Doc = astType.Doc
	p.parseType(packname, file, t, astType.Type)
	p.Types[name] = t

	return p.Types[name]
}

func (p *Parser) parseType(packname string, file *File, t *Type, astType ast.Expr) {
	switch expr := astType.(type) {
	case *ast.Ident:
		if basicTypes[expr.String()] {
			t.Type = expr.String()
		} else {
			// composite in pkg
			t.Type = "ref"
			t.RefPackage = packname
			t.RefTypeName = expr.String()
			t.RefType = p.Types[t.RefPackage+"."+t.RefTypeName]
		}
	case *ast.ArrayType:
		t.Type = "array"
		t.ArrayType = &Type{}
		p.parseType(packname, file, t.ArrayType, expr.Elt)
	case *ast.MapType:
		t.Type = "map"
		t.MapType = &Type{}
		t.MapKey = fmt.Sprint(expr.Key)
		p.parseType(packname, file, t.MapType, expr.Value)
	case *ast.SelectorExpr:
		t.Type = "ref"
		t.RefTypeName = expr.Sel.Name
		t.RefPackage = fmt.Sprint(expr.X)

		if t.RefPackage == "time" {
			t.Type = "time"
			break
		}

		//check is alias
		for key, val := range file.Alias {
			if key == t.RefPackage {
				t.RefPackage = val
			}
		}
		p.parseFromGoPath(file.Imports[t.RefPackage])
		if ref, ok := p.Types[file.Imports[t.RefPackage]+"."+t.RefTypeName]; ok {
			t.RefType = ref
		}
	case *ast.StarExpr:
		t.IsPointer = true
		p.parseType(packname, file, t, expr.X)
	case *ast.StructType:
		t.Type = "struct"
		t.Properties = make(map[string]*Type)
		for i := 0; i < expr.Fields.NumFields(); i++ {
			fieldType := &Type{}
			isComposite := false
			fieldType.Name, isComposite = p.readPropName(expr.Fields.List[i])
			p.parseType(packname, file, fieldType, expr.Fields.List[i].Type)

			// composite out side pkg
			if isComposite && fieldType.RefType != nil && fieldType.RefType.Properties != nil {
				for key, val := range fieldType.RefType.Properties {
					t.Properties[key] = val
				}
			} else {
				t.Properties[fieldType.Name] = fieldType
			}
		}
	default:
		t.Type = "unknown"
	}
}

func (p *Parser) readPropName(field *ast.Field) (string, bool) {
	if len(field.Names) == 0 {
		return "", true
	}

	return field.Names[0].String(), false
}

var basicTypes = map[string]bool{
	"bool":       true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"float32":    true,
	"float64":    true,
	"string":     true,
	"complex64":  true,
	"complex128": true,
	"byte":       true,
	"rune":       true,
	"uintptr":    true,
	"error":      true,
	"Time":       true,
}
