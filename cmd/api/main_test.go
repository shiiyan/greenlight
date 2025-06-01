package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHealthCheckHandler(b *testing.B) {
	app := new(application)

	w := httptest.NewRecorder()
	r := new(http.Request)
	for n := 0; n < b.N; n++ {
		app.healthcheckHandler(w, r)
	}
}
