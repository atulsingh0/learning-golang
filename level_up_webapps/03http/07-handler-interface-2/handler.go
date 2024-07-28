package main

import (
	"fmt"
	"net/http"
	"time"
)

type UptimeHandler struct {
	started time.Time
}

func (u UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current uptime is: %s", time.Since(u.started))
}

func main() {
	// creating custom server
	mux := http.NewServeMux()
	u := UptimeHandler{time.Now()}

	mux.Handle("/", u)

	http.ListenAndServe(":3000", mux)

}
