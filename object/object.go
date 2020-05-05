package object

import "fmt"

// Type indicates the type of the object.
type Type string

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
)

// Object represents any internal value in our language
type Object interface {
	Type() Type
	Inspect() string
}

// Integer represents integer values
type Integer struct {
	Value int64
}

// Inspect returns the value as a string
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the Integer type name
func (i *Integer) Type() Type { return IntegerObj }

// Boolean represents boolean values
type Boolean struct {
	Value bool
}

// Inspect returns a string representation of the boolean value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns the Boolean type name
func (b *Boolean) Type() Type { return BooleanObj }

// Null represent a null value
type Null struct{}

// Inspect returns a string representation of null
func (n *Null) Inspect() string { return "null" }

// Type returns the Null type name
func (n *Null) Type() Type { return NullObj }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type      { return ReturnValueObj }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() Type      { return ErrorObj }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
