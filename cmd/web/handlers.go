package main

import "net/http"

func (a *app) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	a.infoLog.Println("hit virtual terminal")
	if err := a.renderTemplate(w, r, "terminal", nil); err != nil {
		a.errorLog.Println("unable to render terminal page template", err)
	}
}
