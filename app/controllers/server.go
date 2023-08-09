package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"regexp"
	"strconv"

	"app/config"
	"app/models"
)

func generateHTML(w http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/templates/%s.html", file))
	}

	// `template.Must()` wraps (*Template, error) returning function, and panics on error
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (s models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		s = models.Session{UUID: cookie.Value}
		if ok, _ := s.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}

	return s, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}

		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartServer() error {
	files := http.FileServer(http.Dir(config.Config.Srv.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":" + config.Config.Srv.Port, nil)
}
