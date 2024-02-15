package optype

import (
	"database/sql/driver"
	"fmt"
	"io"
)

type OpType string

const (
	OpTypeInsert OpType = "INSERT"
	OpTypeUpdate OpType = "UPDATE"
	OpTypeDelete OpType = "DELETE"
)

var opTypes = []string{
	OpTypeInsert.String(),
	OpTypeUpdate.String(),
	OpTypeDelete.String(),
}

// Values provides list valid values for Enum.
func (OpType) Values() (kinds []string) {
	kinds = append(kinds, opTypes...)
	return
}

func (op OpType) Value() (driver.Value, error) {
	return op.String(), nil
}

func (op OpType) String() string {
	return string(op)
}

func (op OpType) MarshalGQL(w io.Writer) {
	io.WriteString(w, op.String())
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (op *OpType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*op = OpType(str)

	switch *op {
	case OpTypeInsert, OpTypeUpdate, OpTypeDelete:
		return nil
	default:
		return fmt.Errorf("%s is not a valid history operation type", str)
	}
}