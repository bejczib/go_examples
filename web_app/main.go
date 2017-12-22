package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", worker)
	http.ListenAndServe(":8080", nil)

}

func worker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "kecskefa")
	time.Sleep(time.Millisecond * 1)
}
