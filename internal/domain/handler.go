package domain

import (
	"fmt"
	"net/http"
)

func AHandler(w http.ResponseWriter, r *http.Request) {
	data := `{
		"message": "ok"
	}`
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, data)
}
