package models

import (
	"auth/internals/models/user"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

func CreateSessionManager() SessionManager {
	return SessionManager{
		sessions: make(map[string]*Session),
		lock:     &sync.RWMutex{},
	}
}

func generateUUID() string {
	id, _ := uuid.NewV7()
	return id.String()
}

func (sm *SessionManager) CreateSession(userID int, username string, profil string) *Session {
	s := &Session{
		Token:    generateUUID(),
		UserID:   userID,
		Username: username,
		Profil:   profil,
		TimeOut:  time.Now().Add(30 * time.Minute),
	}
	sm.lock.Lock()
	sm.sessions[s.Token] = s
	sm.lock.Unlock()
	return s
}

func (sm *SessionManager) DeleteSession(token string) {
	defer sm.lock.Unlock()
	sm.lock.Lock()
	delete(sm.sessions, token)
}

func (sm *SessionManager) GetSession(token string) (*Session, bool) {
	sm.lock.Lock()
	session, ok := sm.sessions[token]
	sm.lock.Unlock()
	if !ok {
		return nil, false
	}
	return session, true
}

func (sm *SessionManager) DeleteSessionWithUserID(userId int) {
	for _, session := range sm.sessions {
		if session.UserID == userId {
			sm.DeleteSession(session.Token)
		}
	}
}

func (s *Session) IsExpired() bool {
	return s.TimeOut.Before(time.Now())
}

func (app *App) OpenSession(w http.ResponseWriter, user user.User) {
	app.SessionHandler.DeleteSessionWithUserID(user.ID)
	session := app.SessionHandler.CreateSession(user.ID, user.Username, user.Email)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   session.Token,
		Expires: session.TimeOut,
	})
}
