package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/StephanUllmann/bookings-go/pkg/config"
	"github.com/StephanUllmann/bookings-go/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData (templateData *models.TemplateData) *models.TemplateData {
	return templateData
}


func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	// create template cache
	var templateCache map[string]*template.Template
	var err error
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}

	}

	// get requested template from cache
	templ, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not read template from template cache")
	}

	buffer := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	err = templ.Execute(buffer, templateData)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all files *.page.templ
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range throug *.page...
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}
	return myCache, nil

}

// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := templateCache[t]

// 	if !inMap{
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		log.Println("created template in cache")
// 		} else {
// 			log.Println("using cached template")
// 		}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	templateCache[t] = tmpl
// 	return nil
// }
