package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/ ", Handler)


	http.ListenAndServe(":8080", server)
}
