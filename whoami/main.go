package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var StartTime time.Time

func init() {
	StartTime = time.Now()
}

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		info := map[string]interface{}{
			"StartTimeSecs":   StartTime.UTC().Unix(),
			"CurrentTimeSecs": now.UTC().Unix(),
			"Uptime":          now.Sub(StartTime),
		}
		header := make(map[string]interface{})
		for k, v := range r.Header {
			header[k] = v
		}
		info["Header"] = header

		data, _ := json.Marshal(info)
		w.Write(data)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm %s", r.Host)
	})
	http.ListenAndServe(":3000", nil)
}
