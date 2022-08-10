package controller

import (
	"fmt"
	"net/http"
	"net/url"

	"example.com/loginform/conector"
)

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("index is parse here \n")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	password := r.FormValue("password")
	address := r.FormValue("address")
	Image := r.FormValue("Image")
	if len(name) <= 2 || len(password) <= 2 || len(address) <= 2 || len(Image) <= 2 {
		fmt.Printf("invald name orpasswd or address length")
		fmt.Fprintf(w, "invald name or passwd or address length")
	} else {
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
		fmt.Fprintf(w, "Password = %s\n", password)
		fmt.Fprintf(w, "image=%s\n", Image)
		v, err := url.ParseQuery(Image)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%+v\n", v)

		login := conector.LoginData{
			Name:     string(name),
			Password: string(password),
			Address:  string(address),
			Image:    string(Image),
		}
		conector.DB.Create(&login)
	}

}
