package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/user"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving homepage")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "<h1>Hello</h1>")

		//ctx := r.Context()
		ctx := appengine.NewContext(r)
		u := user.Current(ctx)
		fmt.Fprintf(w, "User: %v <br>\n", u)
		if u != nil {
			fmt.Fprintf(w, "Admin? %v <br>\n", u.Admin)
		}

		if u == nil {
			loginURL, err := user.LoginURL(ctx, "/")
			if err == nil {
				fmt.Fprintf(w, `<a href="%s">Sign in</a> <br>`, loginURL)
			} else {
				fmt.Fprintf(w, `LoginURL error=%v`, err)
			}
		} else {
			logoutURL, err := user.LogoutURL(ctx, "/")
			if err == nil {
				fmt.Fprintf(w, `<a href="%s">Sign out</a> <br>`, logoutURL)
			} else {
				fmt.Fprintf(w, `LogoutURL error=%v <br>`, err)
			}
		}
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving Admin section")
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "ADMIN SECTION")

		ctx := r.Context()
		u := user.Current(ctx)
		fmt.Fprintf(w, "User: %v \n", u)
		if u != nil {
			fmt.Fprintf(w, "Admin? %v \n", u.Admin)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}
