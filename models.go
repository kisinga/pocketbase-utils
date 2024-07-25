package main

import "context"

type DB interface {
	Connect() error
	Close() error
	InsertOne(collection string, data interface{}) error
	FindOne(collection string, query interface{}, result interface{}) error
	FindAll(collection string, query interface{}, result interface{}) error
	UpdateOne(collection string, query interface{}, update interface{}) error
	DeleteOne(collection string, query interface{}) error
	DeleteAll(collection string, query interface{}) error
	NewTransation(ctx context.Context, validate bool) (Transaction, error)
}

type Transaction interface {
	Add(collection string, data interface{}) error
	Read(collection string, data interface{}) error
	Update(collection string, data interface{}) error
	Delete(collection string, id int) error
	Commit() (IDs []string, err error)
	Rollback() error
}
