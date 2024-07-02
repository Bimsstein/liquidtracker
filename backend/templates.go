package main

import (
	"html/template"
)

var tmpl = template.Must(template.New("index").Parse(`
<html>
<head>
    <title>Feedback</title>
</head>
<body>
    <form action="/" method="post">
        <label for="brand">Marke:</label><br>
        <select id="brand" name="brand">
            {{range .Brands}}
            <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
        </select><br>
        <label for="variety">Sorte:</label><br>
        <input type="text" id="variety" name="variety"><br>
        <label for="rating">Bewertung (1-5):</label><br>
        <input type="number" id="rating" name="rating" min="1" max="5"><br>
        <input type="submit" value="Submit">
    </form>

    <h2>Add New Brand</h2>
    <form action="/add" method="post">
        <label for="brand_name">Brand Name:</label>
        <input type="text" id="brand_name" name="brand_name">
        <input type="submit" value="Add Brand">
    </form>
</body>
</html>
`))
