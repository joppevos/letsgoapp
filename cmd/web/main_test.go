package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowSnippetNegativeNumber(t *testing.T) {

}

func TestShowSnippet(t *testing.T) {
	tests := []struct {
		name   string
		target string
		want   int
	}{
		{name: "Negative numbers should return StatusNotFuund",
			target: "/snippet?id=-1",
			want:   http.StatusNotFound},
		{name: "postive numbers should return oke",
			target: "/snippet?id=1",
			want:   http.StatusOK},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.target, nil)
			w := httptest.NewRecorder()

			showSnippet(w, req)

			resp := w.Result()
			got := resp.StatusCode

			if resp.StatusCode != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}

		})
	}
}
