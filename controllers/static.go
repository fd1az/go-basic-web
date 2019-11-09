package controllers

import (
	"github.com/fdiaz7/go-basic-web/views"
)

//NewStatic -> return a pointer to a Static struct
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

//Static struct
type Static struct {
	Home    *views.View
	Contact *views.View
}
