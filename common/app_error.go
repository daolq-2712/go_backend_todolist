package common

import (
	"errors"
	"fmt"
	"strings"
)

type AppError struct {
	RootErr error  `json:"_"`
	Message string `json:"message"`
	Log     string `json:"log"`
	Key     string `json:"key"`
}

func NewFullErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{RootErr: root, Message: msg, Log: log, Key: key}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		RootErr: root,
		Message: msg,
		Log:     log,
		Key:     key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(err, "Some thing went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "ErrInvalid")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}
