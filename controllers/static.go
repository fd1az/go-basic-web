package controllers

import (
	"github.com/fdiaz7/go-basic-web/views"
)

//NewStatic -> return a pointer to a Static struct
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

//Static struct
type Static struct {
	Home    *views.View
	Contact *views.View
}
