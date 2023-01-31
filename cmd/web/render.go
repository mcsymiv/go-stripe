package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// templateData holds default data structures
// that can be passed to template pages
type templateData struct {
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

// functions holds custom template functions
// that can be passed to template pages
var functions = template.FuncMap{}

// Go comment directive
// embeds template directory to binaries
// syntax: //directive (no space between `//` and directive)
//
//go:embed templates
var templateFileSystem embed.FS

// default template directory name
// contains .partial.tmpl, .base.tmpl, page.tmpl files
var tempDirectory string = "templates"

// add DefaultTemplateData populates templateData struct
// with default data structures and/or values in tempalate pages
func (a *app) addDefaultTemplateData(td *templateData, r *http.Request) *templateData {
	return td
}

func (a *app) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	// template to render from templates directory
	toRender := fmt.Sprintf("%s/%s.page.tmpl", tempDirectory, page)

	// checks if app already has template to render
	// in cached templateCache map
	// if no template is cached
	// renders new template
	_, cached := a.templageCache[toRender]
	if !cached {
		a.infoLog.Println(fmt.Sprintf("parsing %s template", toRender))
		t, err = a.parseTemplate(toRender, page, partials)
		if err != nil {
			a.errorLog.Println("unable to parse template")

			return err
		}
	}

	// Gets template from application config cache
	// if page has already been parsed and created
	t = a.templageCache[toRender]

	// Add default template data if none was provided
	// to the page template
	if td == nil {
		td = &templateData{}
	}

	// Add default template data to provided in render
	td = a.addDefaultTemplateData(td, r)

	// Execute template
	err = t.Execute(w, td)
	if err != nil {
		a.errorLog.Println("unable to execute template")

		return err
	}

	return nil
}

// parseTemplate builds pages from template partials
// (base, defined pages, layouts)
func (a *app) parseTemplate(tmplToTender, page string, partials []string) (*template.Template, error) {
	var t *template.Template
	var err error

	for i, p := range partials {
		partials[i] = fmt.Sprintf("%s/%s.partial.tmpl", tempDirectory, p)
	}

	t, err = template.
		New(fmt.Sprintf("%s.page.tmpl", page)).
		Funcs(functions).
		ParseFS(
			templateFileSystem,
			fmt.Sprintf("%s/base.layout.tmpl", tempDirectory),
			// strings.Join(partials, ","),
			tmplToTender,
		)
	if err != nil {
		a.errorLog.Println("unable to create new template", err)

		return nil, err
	}

	// Puts parsed template to application config cache
	a.templageCache[tmplToTender] = t

	return t, nil
}
