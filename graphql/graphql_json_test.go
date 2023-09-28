package graphql

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoJSON(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		assert.Equal(t, r.Method, http.MethodPost)
		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(b), `{"query":"query {}","variables":null}`+"\n")
		io.WriteString(w, `{
			"data": {
				"something": "yes"
			}
		}`)
	}))
	defer srv.Close()

	ctx := context.Background()
	client := NewClient(srv.URL)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var responseData map[string]interface{}
	err := client.Run(ctx, &Request{q: "query {}"}, &responseData)
	assert.Nil(t, err)
	assert.Equal(t, calls, 1) // calls
	assert.Equal(t, responseData["something"], "yes")
}

func TestDoJSONServerError(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		assert.Equal(t, r.Method, http.MethodPost)
		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(b), `{"query":"query {}","variables":null}`+"\n")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `Internal Server Error`)
	}))
	defer srv.Close()

	ctx := context.Background()
	client := NewClient(srv.URL)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var responseData map[string]interface{}
	err := client.Run(ctx, &Request{q: "query {}"}, &responseData)
	assert.Equal(t, calls, 1) // calls
	assert.Equal(t, err.Error(), "graphql: server returned a non-200 status code: 500")
}

func TestDoJSONBadRequestErr(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		assert.Equal(t, r.Method, http.MethodPost)
		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(b), `{"query":"query {}","variables":null}`+"\n")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"errors": [{
				"message": "miscellaneous message as to why the the request was bad"
			}]
		}`)
	}))
	defer srv.Close()

	ctx := context.Background()
	client := NewClient(srv.URL)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var responseData map[string]interface{}
	err := client.Run(ctx, &Request{q: "query {}"}, &responseData)
	assert.Equal(t, calls, 1) // calls
	assert.Equal(t, err.Error(), "graphql: miscellaneous message as to why the the request was bad")
}

func TestQueryJSON(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(b), `{"query":"query {}","variables":{"username":"matryer"}}`+"\n")
		_, err = io.WriteString(w, `{"data":{"value":"some data"}}`)
		assert.Nil(t, err)
	}))
	defer srv.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := NewClient(srv.URL)

	req := NewRequest("query {}")
	req.Var("username", "matryer")

	// check variables
	assert.True(t, req != nil)
	assert.Equal(t, req.vars["username"], "matryer")

	var resp struct {
		Value string
	}
	err := client.Run(ctx, req, &resp)
	assert.Nil(t, err)
	assert.Equal(t, calls, 1)

	assert.Equal(t, resp.Value, "some data")
}

func TestHeader(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		assert.Equal(t, r.Header.Get("X-Custom-Header"), "123")

		_, err := io.WriteString(w, `{"data":{"value":"some data"}}`)
		assert.Nil(t, err)
	}))
	defer srv.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := NewClient(srv.URL)

	req := NewRequest("query {}")
	req.Header.Set("X-Custom-Header", "123")

	var resp struct {
		Value string
	}
	err := client.Run(ctx, req, &resp)
	assert.Nil(t, err)
	assert.Equal(t, calls, 1)

	assert.Equal(t, resp.Value, "some data")
}
