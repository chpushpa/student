package students

import (
	"github.com/chpushpa/student/domain/httperrors"
)

type Student struct {
	Id     int64    `json:"id"`
	Name   string   `json:"name"`
	Grade  string   `json:"grade"`
	Detail *Details `json:"detail"`
}
type Details struct {
	Class int    `json:"class"`
	Batch string `json:"batch"`
}

func (student Student) Validate() *httperrors.HttpError {
	if student.Name == "" {
		return httperrors.NewBadRequestError("invalid name")
	}
	if student.Grade == "" {
		return httperrors.NewBadRequestError("invalid grade address")
	}
	return nil
}
