package main

import (
	"net/http"
	"uwu/uwu"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["uwu"]

	if !ok || len(keys[0]) < 1 {
		w.Write([]byte("'uwu' param is missing"))
		return
	}

	key := keys[0]

	w.Write([]byte(uwu.UwUify([]rune(key))))
}
