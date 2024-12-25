package object

import "fmt"

type ObjectType string

const (
	INTERGER_OBJ = "INTEGER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTERGER_OBJ }
