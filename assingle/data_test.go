package assingle

import (
	"time"
)

//go:generate msgp

type DataTest struct {
	Name      string    `msg:"n"`
	Email     string    `msg:"e"`
	Age       int       `msg:"a"`
	Scope     []string  `msg:"s"`
	CreatedAt time.Time `msg:"t"`
}
