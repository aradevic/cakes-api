package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"alexco.waracle.com/cakes/repo"
)

var service = NewService(&repo.TestDBRepo{})

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

func TestGetCakes(t *testing.T) {
	t.Run("Test get all cakes", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/cakes/", nil)
		response := httptest.NewRecorder()

		service.GetCakes(response, request)

		got := response.Body.String()
		want := "[]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Test get cakes by name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/cakes?name=lemon", nil)
		response := httptest.NewRecorder()

		service.GetCakes(response, request)

		got := response.Body.String()
		want := "[]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func TestAddCakes(t *testing.T) {
	t.Run("Test add a new cake", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost,
			"/cakes",
			strings.NewReader(`{"id":1,"name":"NY cheesecak","YumFactor":2}`))

		response := httptest.NewRecorder()

		service.AddCake(response, request)
	})
}
func TestDeleteCakes(t *testing.T) {
	t.Run("Test delete cake", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, "/cakes", nil)
		request.SetPathValue("id", "1")

		response := httptest.NewRecorder()

		service.DeleteCakes(response, request)
	})
}
