package main

import "net/http"

func (a *application) loadSession(next http.Handler) http.Handler {
	return a.session.LoadAndSave(next)
}
