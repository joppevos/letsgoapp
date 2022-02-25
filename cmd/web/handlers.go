package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joppevos/letsgoapp/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, err)
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

func (app application) showSnippetForm(w http.ResponseWriter, r *http.Request)  {
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
		Snippet:  s,
	} )


}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	title := "my title"
	content := "my content"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
