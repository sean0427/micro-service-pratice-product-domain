package mongodb

import (
	"fmt"
	"strings"
)

const msg = "mongo error"

type MongoError struct {
	err error
}

func NewMongoError(err error) error {
	return &MongoError{err: err}
}

func (e *MongoError) Error() string {
	return fmt.Sprintf("%s %v", msg, e.err)
}

func (m *MongoError) Is(target error) bool { return strings.HasPrefix(target.Error(), msg) }

func (m *MongoError) Unwrap() error { return m.err }
