package main

import (
	"embed"
	"html/template"
)

type templateDate struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CRSFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	Api             string
	CSSVersion      string
}

var functions = template.FuncMap{}

// go:embed templates
var templateFileSystem embed.FS
