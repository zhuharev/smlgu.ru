package models

import (
	"fmt"
)

type FieldMismatch struct {
	GotFileds int
	NeedFiled int
}

func (fm *FieldMismatch) Error() string {
	return fmt.Sprintf("[Field mismatch] got %d fields, need %d", fm.GotFileds, fm.NeedFiled)
}

type UnsupportedType struct {
	Type string
}

func (ut *UnsupportedType) Error() string {
	return fmt.Sprintf("[Unsupported type] got %q", ut.Type)
}
