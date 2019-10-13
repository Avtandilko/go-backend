package main

import (
	"os"
)

// DbConn ...
type DbConn struct {
	DbHost     string
	DbName     string
	DbUsername string
	DbPassword string
}

// NewDbConn ...
func NewDbConn() *DbConn {
	return &DbConn{
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}
