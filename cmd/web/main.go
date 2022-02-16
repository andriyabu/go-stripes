package main

import (
	"log"
	"text/template"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int    // port number
	env  string // environtment : development or production
	api  string // api URL
	db   struct {
		dsn string //database source name
	}

	stripe struct {
		secret string
		key    string
	}
}

//receiver for a various application
type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}
