package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const homePageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Text Encoder/Decoder</title>
	<style>
		body {
			text-align: center;
			font-family: Arial, sans-serif;
			background-color: #f1f0f0;
			margin: 0;
			padding: 0;
		}
		.container {
			display: flex;
			justify-content: center;
			gap: 20px;
			flex-wrap: wrap;
			padding: 20px;
		}
		.form-container {
			background-color: #ffffff;
			border-radius: 8px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			padding: 20px;
			width: 45%;
			max-width: 500px;
			box-sizing: border-box;
		}
		textarea {
			width: 100%;
			box-sizing: border-box;
			margin-bottom: 10px;
			padding: 10px;
			border-radius: 4px;
			border: 1px solid #ddd;
			font-family: Arial, sans-serif;
		}
		input[type="submit"] {
			padding: 10px 20px;
			border: none;
			border-radius: 4px;
			background-color: #007BFF;
			color: #ffffff;
			cursor: pointer;
			font-size: 16px;
		}
		input[type="submit"]:hover {
			background-color: #0056b3;
		}
		h1, h2 {
			margin: 0 0 10px;
		}
	</style>
</head>
<body>
	<h1>Text Encoder/Decoder</h1>
	<div class="container">
		<div class="form-container">
			<form action="/encode" method="post">
				<h2>Encode Text</h2>
				<textarea name="text" rows="10" placeholder="Enter text to encode..."></textarea><br>
				<input type="submit" value="Encode">
			</form>
		</div>
		<div class="form-container">
			<form action="/decode" method="post">
				<h2>Decode Text</h2>
				<textarea name="text" rows="10" placeholder="Enter encoded text to decode..."></textarea><br>
				<input type="submit" value="Decode">
			</form>
		</div>
	</div>
</body>
</html>`

const resultPageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Art Encoder/Decoder</title>
	<style>
		body {
			text-align: center;
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
			margin: 0;
			padding: 0;
		}
		h1 {
			margin-top: 20px;
		}
		pre {
			text-align: left;
			display: inline-block;
			white-space: pre-wrap; /* Preserve whitespace and line breaks */
			max-width: 90%;
			margin: 20px auto;
			padding: 20px;
			background-color: #ffffff;
			border-radius: 8px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		}
		a {
			display: inline-block;
			margin-top: 20px;
			padding: 10px 20px;
			border-radius: 4px;
			background-color: #007BFF;
			color: #ffffff;
			text-decoration: none;
			font-size: 16px;
		}
		a:hover {
			background-color: #0056b3;
		}
	</style>
</head>
<body>
	<h1>{{.Title}}</h1>
	<pre>{{.Result}}</pre>
	<a href="/">Return to Homepage</a>
</body>
</html>`

var resultTemplate = template.Must(template.New("result").Parse(resultPageTemplate))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homePageTemplate)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("text")
		encoded := Encoder(input)
		data := map[string]string{
			"Title":  "Encoded Art",
			"Result": encoded,
		}
		if err := resultTemplate.Execute(w, data); err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		}
	}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("text")
		decoded, err := Decoder(input)
		data := map[string]string{
			"Title":  "Decoded Art",
			"Result": decoded,
		}
		if err != nil {
			data["Title"] = "Error"
			data["Result"] = "Error: " + err.Error()
		}
		if err := resultTemplate.Execute(w, data); err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		}
	}
}

// Start the HTTP server
func startServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
