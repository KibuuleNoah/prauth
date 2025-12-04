package entities

import (
	"github.com/gorilla/sessions"
)

type AppCtx struct{
	CookieStore sessions.CookieStore
}

