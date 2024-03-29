package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {

	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, loyoutFiles()...)
	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

//Returns a slice of strings representing the layout files in our app
func loyoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

//Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//addTemplatePath takes slices strings representing files paths for templates
//adn prepends the templateDir directory to each string in the slice
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

//addTemplateExt takes slices strings representing files paths for templates
//adn prepends the TemplateExt extension to each string in the slice
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
