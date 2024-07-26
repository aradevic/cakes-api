package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"alexco.waracle.com/cakes/repo"
)

var service = Service{D: &repo.TestDBRepo{}}

func TestFindCakesById(t *testing.T) {
	t.Run("Test find by Id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/cakes/", nil)
		request.SetPathValue("id", "1")
		response := httptest.NewRecorder()

		service.FindCakesById(response, request)
		got := response.Body.String()
		want := "{1 lemon chesse cake the best  5}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
