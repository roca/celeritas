package session

import "github.com/alexedwards/scs/v2"

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
}

func (c *Session) InitSession() *scs.SessionManager {}
