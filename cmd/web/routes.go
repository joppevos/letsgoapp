package main

import (
	"github.com/justinas/alice"
	"net/http"
	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	r := mux.NewRouter()
	r.HandleFunc("/", app.home)
	r.HandleFunc("/snippet", app.showSnippet).Methods(http.MethodGet)
	r.HandleFunc("/snippet/create", app.showSnippetForm).Methods(http.MethodGet)
	r.HandleFunc("/snippet/create", app.createSnippet).Methods(http.MethodPost)
	r.HandleFunc("/snippet/{id}", app.showSnippet).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(r)
}
