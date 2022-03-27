package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", HomeHandler)
	var port = os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}
