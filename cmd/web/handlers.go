package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joppevos/letsgoapp/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, err)
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		}
		app.serveError(w, err)
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.Form.Get("expires")

	errors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors:  errors,
			FormData:    r.PostForm,
		})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
