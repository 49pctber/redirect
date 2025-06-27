package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/49pctber/redirect"
)

var default_redirect string

func RedirectDefault(w http.ResponseWriter, r *http.Request) {
	if default_redirect != "" {
		http.Redirect(w, r, default_redirect, http.StatusTemporaryRedirect)
	} else {
		fmt.Fprintf(w, "There was an error with the request.")
	}
}

func main() {
	// set default redirect for invalid labels
	default_redirect = ""
	if len(os.Args) > 1 {
		u, err := url.Parse(os.Args[1])
		if err == nil {
			default_redirect = u.String()
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RedirectDefault(w, r)
	})

	http.HandleFunc("GET /{label}", func(w http.ResponseWriter, r *http.Request) {
		rv, err := redirect.GetRedirect(r.PathValue("label"))
		if err != nil {
			RedirectDefault(w, r)
		} else {
			http.Redirect(w, r, rv.GetDestination(), http.StatusTemporaryRedirect)
		}
	})

	http.ListenAndServe(":8080", nil)
}
