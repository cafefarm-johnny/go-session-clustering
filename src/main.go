package main

import (
	"Go-Session-Clustering/src/session"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/find", session.Find)
	http.HandleFunc("/login", session.Login)
	http.HandleFunc("/logout", session.Logout)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
