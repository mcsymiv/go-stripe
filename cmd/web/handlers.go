package main

import "net/http"

func (a *app) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	a.infoLog.Println("hit virtual terminal")
}
