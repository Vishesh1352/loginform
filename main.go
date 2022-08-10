package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/loginform/conector"
	"example.com/loginform/controller"
)

func main() {

	//r := gin.Default()
	//r.Static("/", "loginform/")
	//r.SetFuncMap(template.FuncMap{
	//	"upper": strings.ToUpper,
	//})
	//r.LoadHTMLGlob("loginform/*.html")
	conector.ConnectToDatabase()

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)

	//r.GET("/signin/:Name/:Password", controller.Signin)
	http.HandleFunc("/signin", controller.Signin)
	http.HandleFunc("/login", controller.Login)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//r.Run("localhost:8080")
}
