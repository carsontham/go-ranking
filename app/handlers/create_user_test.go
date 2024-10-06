package handlers_test

import (
	"bytes"
	"errors"
	"go-ranking/app/handlers"
	"go-ranking/app/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-ranking/tests/repositorytest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	setUp := func(t *testing.T) (
		repoMock *repositorytest.MockRankingRepository,
		c *http.Client,
		url string,
	) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockRankingRepository(ctrl)
		v := validator.New()

		router := chi.NewRouter()
		router.Post("/users", handlers.CreateNewUser(repoMock, v))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/users"
		return
	}

	t.Run("it should successfully create a new account, return 201", func(t *testing.T) {
		repoMock, client, url := setUp(t)
		stubEmail := "newUser@hello.com"
		reqBody := `
			{
				"name": "newUser",
				"email": "newUser@hello.com",
				"score": 200
			}`
		stubUser := &repository.User{
			Name:  "newUser",
			Email: "newUser@hello.com",
			Score: 200,
		}
		repoMock.EXPECT().CheckUniqueEmail(stubEmail).Times(1).Return(true, nil)
		repoMock.EXPECT().CreateNewUser(stubUser).Times(1).Return(nil)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("it should fail as email is already in use, return 409", func(t *testing.T) {
		repoMock, client, url := setUp(t)
		stubEmail := "newUser@hello.com"
		reqBody := `
			{
				"name": "newUser",
				"email": "newUser@hello.com",
				"score": 200
			}`
		repoMock.EXPECT().CheckUniqueEmail(stubEmail).Times(1).Return(false, nil)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusConflict, resp.StatusCode)
		}
	})

	t.Run("it should fail as request body is malformed, return 400", func(t *testing.T) {
		_, client, url := setUp(t)
		//stubEmail := "newUser@hello.com"
		reqBody := `
			{
				"invalid_body": newUser@hello.com
				,
			}`
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("it should fail due to validation checks when invalid email format, return 422", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
				"name": "newUser",
				"email": "newUser@hello",
				"score": 1
			}`
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		}
	})

	t.Run("it should fail due to validation checks when score is less than 0, return 422", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
				"name": "newUser",
				"email": "newUser@hello.com",
				"score": -1
			}`
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		}
	})

	t.Run("it should fail due to internal server error, return 500", func(t *testing.T) {
		repoMock, client, url := setUp(t)
		stubEmail := "newUser@hello.com"
		reqBody := `
			{
				"name": "newUser",
				"email": "newUser@hello.com",
				"score": 200
			}`
		stubUser := &repository.User{
			Name:  "newUser",
			Email: "newUser@hello.com",
			Score: 200,
		}
		repoMock.EXPECT().CheckUniqueEmail(stubEmail).Times(1).Return(true, nil)
		repoMock.EXPECT().CreateNewUser(stubUser).Times(1).Return(errors.New("db error"))
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})
}
