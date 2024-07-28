package main

import (
	"fmt"
	"net/http"
	"time"
)

type UptimeHandler struct {
	started time.Time
}

// SecretTokenHandler chained handlers
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

func (s SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Secret-Token") == s.secret {
		s.next.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
	}
}

func (u UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current uptime is: %s", time.Since(u.started))
}

func main() {
	// creating custom server
	mux := http.NewServeMux()
	u := UptimeHandler{time.Now()}
	s := SecretTokenHandler{next: u, secret: "1248"}

	mux.Handle("/", s)

	http.ListenAndServe(":3000", mux)

}
