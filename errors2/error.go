package errors2

const (
	Internal = iota
	BadReq
	NotFound
	Unauthorized
	Forbidden
)

type Error interface {
	Code() int
	Error() string
	Type() int
}

type Errors2 struct {
	code    int    `json:"-"`
	message string `json:"error"`
	ty      int    `json:"-"`
}

func (err *Errors2) Error() string {
	return err.message
}

func (err *Errors2) Code() int {
	return err.code
}

func (err *Errors2) Type() int {
	return err.ty
}

func NewError(code int, msg string, ty int) *Errors2 {
	return &Errors2{code, msg, ty}
}

func NewInternal(msg string) *Errors2 {
	return &Errors2{500, msg, Internal}
}

func NewBadReq(msg string) *Errors2 {
	return &Errors2{400, msg, BadReq}
}

func NewNotFound(msg string) *Errors2 {
	return &Errors2{404, msg, NotFound}
}

func NewUnauthorized(msg string) *Errors2 {
	return &Errors2{401, msg, Unauthorized}
}

func NewForbidden(msg string) *Errors2 {
	return &Errors2{403, msg, Forbidden}
}
