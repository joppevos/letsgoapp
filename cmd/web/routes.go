package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	r := mux.NewRouter()
	r.Handle("/", dynamicMiddleware.ThenFunc(app.home))
	r.Handle("/snippet", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods(http.MethodGet)
	r.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm)).Methods(http.MethodGet)
	r.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet)).Methods(http.MethodPost)
	r.Handle("/snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(r)
}
