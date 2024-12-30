package v1

import (
	"fmt"
	"net/http"
)

func NewAPI() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	return mux
}
