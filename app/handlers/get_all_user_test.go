package handlers_test

import (
	"errors"
	"go-ranking/app/handlers"
	"go-ranking/app/repository"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-ranking/tests/repositorytest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	setUp := func(t *testing.T) (
		repoMock *repositorytest.MockRankingRepository,
		c *http.Client,
		url string,
	) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockRankingRepository(ctrl)
		router := chi.NewRouter()
		router.Get("/users", handlers.GetAllUser(repoMock))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/users"
		return
	}

	t.Run("it should successfully return all users, return 200", func(t *testing.T) {
		repoMock, client, url := setUp(t)

		stubUsers := []*repository.User{
			{
				ID:    1,
				Name:  "John",
				Email: "john@gmail.com",
				Score: 50,
			},
			{
				ID:    2,
				Name:  "Daniel",
				Email: "daniel@gmail.com",
				Score: 60,
			},
			{
				ID:    3,
				Name:  "Poh",
				Email: "poh@gmail.com",
				Score: 70,
			},
		}

		repoMock.EXPECT().GetAllUser(false, 0).Times(1).Return(stubUsers, nil)

		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should fail due to invalid minimum score param, return 400", func(t *testing.T) {
		_, client, url := setUp(t)
		invalidUrl := url + "?minScore=invalid"
		req, _ := http.NewRequest("GET", invalidUrl, nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("it should fail due to internal server error, return 500", func(t *testing.T) {
		repoMock, client, url := setUp(t)

		repoMock.EXPECT().GetAllUser(false, 0).Times(1).Return(nil, errors.New("database error"))
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})
}
