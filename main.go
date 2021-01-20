package main

import (
	"net/http"
	"uwu/uwu"

	strip "github.com/grokify/html-strip-tags-go" // => strip
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
	<h4 style="margin-bototm: 1px">Nightbot</h4>
	<code>!commands add !uwu $(urlfetch json https://uwuify.helba.ai/?uwu=$(querystring));</code>
	<br/>
	<h4 syle="margin-bottom: 1px">Try it out</h4>
	<input id="myInput" placeholder="Uwuify input">
	<button id="myBtn" onclick="uwuify()">Uwuify</button>
	<div id="result"></div>

	<script>
	var input = document.getElementById("myInput");
	input.addEventListener("keyup", function(event) {
	if (event.keyCode === 13) {
		uwuify();
	}
	});

	function uwuify() {
		fetch("/?uwu=" + input.value).then(function(response) {
			return response.text().then(function(text) {
				document.getElementById("result").innerHTML=text;
			});
		  });
	}
	</script>
	<br />
	<a href="https://github.com/Toyz/uwu-server" target="_blank">Source code on Github</a>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["uwu"]

	if !ok || len(keys[0]) < 1 {
		w.Header().Set("Content-Type", "text/html")

		w.Write([]byte(html))
		return
	}

	key := keys[0]

	setupResponse(&w, r)

	if len(key) > 256 {
		w.Write([]byte("'uwu' max legnth is 256 characters"))
		return
	}

	key = strip.StripTags(key)
	w.Write([]byte(uwu.UwUify([]rune(key))))
}
