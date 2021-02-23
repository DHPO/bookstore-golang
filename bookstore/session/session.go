package session

import "github.com/gorilla/sessions"

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("kjlngalkjfnasl;kfna"))
}
