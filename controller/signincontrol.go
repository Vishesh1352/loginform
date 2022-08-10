package controller

import (
	"fmt"
	"net/http"

	"example.com/loginform/conector"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var result conector.LoginData

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "we got here to signin")

	password := r.FormValue("password")
	Name := r.FormValue("name")

	err := conector.DB.Where("name=?", Name).First(&result).Error
	if err != nil {
		fmt.Fprintf(w, "ParseForm()    	 err: %v \n", err)
		return
	}

	if password != result.Password {
		fmt.Fprintf(w, "\n Password Authorization Error\n")
		return
	}

	fmt.Fprintf(w, "Address: %v , Image: %v \n", result.ID, result.Image)

	fmt.Fprintln(w, "<html><img src="+result.Image+"></html>")

}
