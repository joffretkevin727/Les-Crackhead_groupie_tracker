package main

import (
	"Les-Crackhead_groupie_tracker/router"
	"fmt"
	"net/http"
)

// FONCTION PRINCIPAL
func main() {
	r := router.New()
	fmt.Println("http://localhost:8080/ressource")
	http.ListenAndServe(":8080", r)

}
