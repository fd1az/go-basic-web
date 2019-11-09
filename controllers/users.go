package controllers

import (
	"fmt"
	"net/http"

	"github.com/fdiaz7/go-basic-web/views"
)

//NewUsers is used to create a new Users controllers
//This function will panic if the templates are not parse correctly
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//Users struct
type Users struct {
	NewView *views.View
}

//SignupForm struc
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//New GET -> This is user tu render the form where a user can create a new user acount
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

//Create POST message
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)

}
