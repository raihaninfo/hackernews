package main

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

type Template struct {
	URL            string
	IsAuthenticate bool
	AuthUser       string
	Flash          string
	Error          string
	CSRFToken      string
}

func (a *application) defaultData(td *Template, r *http.Request) *Template {
	td.URL = a.server.url
	return td
}

func (a *application) render(w http.ResponseWriter, r *http.Request, view string, vars jet.VarMap) error {
	td := &Template{}
	td = a.defaultData(td, r)
	tp, err := a.view.GetTemplate(fmt.Sprintf("%s.html", view))
	if err != nil {
		return err
	}
	err = tp.Execute(w, vars, td)
	if err != nil {
		return err
	}
	return nil
}
