package mongodb

import "fmt"

type MongoError struct {
	err error
}

func NewMongoError(err error) error {
	return &MongoError{err: err}
}

func (e *MongoError) Error() string {
	return fmt.Sprintf("mongo error: %v", e.err)
}
