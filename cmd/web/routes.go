package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddlewares := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	dynamicMiddlewares := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddlewares.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddlewares.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddlewares.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddlewares.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddlewares.Then(mux)
}