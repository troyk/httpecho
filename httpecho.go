package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var openConnections int64 = 0

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliction/json")
	r.ParseMultipartForm(16384)
	atomic.AddInt64(&openConnections, 1)
	myConnNum := atomic.LoadInt64(&openConnections)
	delay, err := strconv.ParseInt(r.Form.Get("delay"), 10, 64)
	if err != nil {
		delay = 0
	}
	time.Sleep(time.Duration(delay+myConnNum) * time.Millisecond)
	r.Form.Set("connections", strconv.FormatInt(myConnNum, 10))
	out, err := json.Marshal(r.Form)
	if err != nil {
		out = []byte(err.Error())
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(out)))
	w.Write(out)
	atomic.AddInt64(&openConnections, -1)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6969", nil)
}
