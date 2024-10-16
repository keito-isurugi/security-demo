package view

import (
	"html/template"
	"net/http"
)

func SessionLoginPage(w http.ResponseWriter, r *http.Request) {
	tpl := template.HTML(`
        <form method="post" action="http://localhost:8080/session_login">
            <label for="id">ID</label>
            <input type="text" name="id">
            <br>
            <label for="password">Password</label>
            <input type="test" name="password">
            <br>
            <input type="submit" label="submit">
        </form>
    `)
	t := template.Must(template.New("a").Parse(`<div>{{ . }}</div>`))
	t.Execute(w, tpl)
}
