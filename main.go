package main

import (
	"net/http"
	"uwu/uwu"
)

const html = `
<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>UwUify</title>
<meta name="description" content="UwUify your text for the fun of it!">
<meta name="author" content="Helba The AI">
</head>
<body>
<h3>UwUify<h3>
	<strong>https://uwuify.helba.ai/?uwu=<code>Your text here</code><strong>
	<br />
	<br />
	<a href="https://github.com/Toyz/uwu-server" target="_blank">Source code on Github</a>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["uwu"]

	if !ok || len(keys[0]) < 1 {
		w.Header().Set("Content-Type", "text/html")

		w.Write([]byte(html))
		return
	}

	key := keys[0]

	w.Write([]byte(uwu.UwUify([]rune(key))))
}
