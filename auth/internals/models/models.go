package models

import (
	"auth/internals/utils"
	"database/sql"
	"net/http"
	"sync"
	"time"
)

type EndPoint struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

type Message struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type Session struct {
	Token    string
	UserID   int
	Username string
	Profil   string
	TimeOut  time.Time
}

type SessionManager struct {
	lock     *sync.RWMutex
	sessions map[string]*Session
}

type App struct {
	Database       *sql.DB
	SessionHandler SessionManager
}

var (
	Validators = map[string]utils.Validator{
		"username":   utils.UsernameValidator{},
		"age":        utils.AgeValidator{},
		"gender":     utils.GenderValidator{},
		"first_name": utils.FirstNameValidator{},
		"last_name":  utils.LastNameValidator{},
		"password":   utils.PasswordValidator{},
		"email":      utils.EmailValidator{},
	}
)
