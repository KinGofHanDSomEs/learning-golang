package main

import (
	"fmt"
	"sync"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}
type SessionManager struct {
	id       int
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {

	return &SessionManager{
		id:       1,
		sessions: make(map[string]Session),
	}
}

func (sm *SessionManager) StartSession(userID string) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	id := fmt.Sprint(sm.id)
	sm.sessions[id] = Session{
		ID:        id,
		UserID:    userID,
		ExpiresAt: time.Now().Add(120 * time.Second),
	}
	sm.id++
	return id
}

func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	session, ok := sm.sessions[sessionID]
	if !ok || time.Now().After(session.ExpiresAt) {
		return Session{}, false
	}
	return session, ok
}
