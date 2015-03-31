package assingle

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// EncodeMsg implements msgp.Encodable
func (z *DataTest) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(5)
	if err != nil {
		return
	}
	err = en.WriteString("n")
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	err = en.WriteString("e")
	if err != nil {
		return
	}
	err = en.WriteString(z.Email)
	if err != nil {
		return
	}
	err = en.WriteString("a")
	if err != nil {
		return
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		return
	}
	err = en.WriteString("s")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Scope)))
	if err != nil {
		return
	}
	for xvk := range z.Scope {
		err = en.WriteString(z.Scope[xvk])
		if err != nil {
			return
		}
	}
	err = en.WriteString("t")
	if err != nil {
		return
	}
	err = en.WriteTime(z.CreatedAt)
	if err != nil {
		return
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *DataTest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "n":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "e":
			z.Email, err = dc.ReadString()
			if err != nil {
				return
			}
		case "a":
			z.Age, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "s":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Scope) >= int(xsz) {
				z.Scope = z.Scope[:xsz]
			} else {
				z.Scope = make([]string, xsz)
			}
			for xvk := range z.Scope {
				z.Scope[xvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "t":
			z.CreatedAt, err = dc.ReadTime()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DataTest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 5)
	o = msgp.AppendString(o, "n")
	o = msgp.AppendString(o, z.Name)
	o = msgp.AppendString(o, "e")
	o = msgp.AppendString(o, z.Email)
	o = msgp.AppendString(o, "a")
	o = msgp.AppendInt(o, z.Age)
	o = msgp.AppendString(o, "s")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Scope)))
	for xvk := range z.Scope {
		o = msgp.AppendString(o, z.Scope[xvk])
	}
	o = msgp.AppendString(o, "t")
	o = msgp.AppendTime(o, z.CreatedAt)
	return
}

//UnmarshalMsg implements msgp.Unmarshaler
func (z *DataTest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "n":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "e":
			z.Email, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "a":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "s":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Scope) >= int(xsz) {
				z.Scope = z.Scope[:xsz]
			} else {
				z.Scope = make([]string, xsz)
			}
			for xvk := range z.Scope {
				z.Scope[xvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "t":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *DataTest) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 1 + msgp.StringPrefixSize + len(z.Name) + msgp.StringPrefixSize + 1 + msgp.StringPrefixSize + len(z.Email) + msgp.StringPrefixSize + 1 + msgp.IntSize + msgp.StringPrefixSize + 1 + msgp.ArrayHeaderSize
	for xvk := range z.Scope {
		s += msgp.StringPrefixSize + len(z.Scope[xvk])
	}
	s += msgp.StringPrefixSize + 1 + msgp.TimeSize
	return
}
