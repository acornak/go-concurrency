package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session *scs.SessionManager
	DB      *sql.DB
	InfoLog *log.Logger
	Error   *log.Logger
	Wait    *sync.WaitGroup
}
