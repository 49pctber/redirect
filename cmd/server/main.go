package main

import (
	"fmt"
	"net/http"

	"github.com/49pctber/redirect"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is your redirect server.")
	})

	http.HandleFunc("GET /{label}", func(w http.ResponseWriter, r *http.Request) {
		rv, err := redirect.GetRedirect(r.PathValue("label"))
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}

		http.Redirect(w, r, rv.GetDestination(), http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":8080", nil)
}
