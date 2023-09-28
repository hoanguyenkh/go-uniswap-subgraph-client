package unigraphclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	testUrl := "test"
	t.Run("when opts is nil", func(t *testing.T) {
		client := NewClient(testUrl, nil)

		assert.NotNil(t, client)
	})

	t.Run("when opts is not nil", func(t *testing.T) {
		opts := &ClientOptions{
			HttpClient: &http.Client{},
			CloseReq:   true,
		}
		client := NewClient(testUrl, opts)

		assert.NotNil(t, client)
	})
}

func TestGetFactoryById(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "factory")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.GetFactoryById(context.Background(), id, nil)
		assert.Nil(t, err)
		assert.Equal(t, id, resp.Factory.ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "factory")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.GetFactoryById(context.Background(), id, opts)
		assert.NotNil(t, err)
	})
}

func TestListFactories(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "factory")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.ListFactories(context.Background(), nil)
		assert.Nil(t, err)
		assert.Len(t, resp.Factories, 1)
		assert.Equal(t, id, resp.Factories[0].ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "factory")
		defer server.Close()

		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.ListFactories(context.Background(), opts)
		assert.NotNil(t, err)
	})
}

func TestGetPoolById(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "pool")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.GetPoolById(context.Background(), id, nil)
		assert.Nil(t, err)
		assert.Equal(t, id, resp.Pool.ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "pool")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.GetPoolById(context.Background(), id, opts)
		assert.NotNil(t, err)
	})
}

func TestListPools(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "pool")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.ListPools(context.Background(), nil)
		assert.Nil(t, err)
		assert.Len(t, resp.Pools, 1)
		assert.Equal(t, id, resp.Pools[0].ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "pool")
		defer server.Close()

		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.ListPools(context.Background(), opts)
		assert.NotNil(t, err)
	})
}

func TestGetTokenById(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "token")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.GetTokenById(context.Background(), id, nil)
		assert.Nil(t, err)
		assert.Equal(t, id, resp.Token.ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessById, "token")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.GetTokenById(context.Background(), id, opts)
		assert.NotNil(t, err)
	})
}

func TestListTokens(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "token")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		resp, err := client.ListTokens(context.Background(), nil)
		assert.Nil(t, err)
		assert.Len(t, resp.Tokens, 1)
		assert.Equal(t, id, resp.Tokens[0].ID)
	})

	t.Run("when query construction fails", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "token")
		defer server.Close()

		client := NewClient(server.URL, nil)

		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := client.ListTokens(context.Background(), opts)
		assert.NotNil(t, err)
	})
}

func TestExecuteRequestAndConvert(t *testing.T) {
	t.Run("when successful", func(t *testing.T) {
		server := getTestServer(t, SuccessList, "factory")
		defer server.Close()

		id := "test"
		client := NewClient(server.URL, nil)

		req, err := constructListQuery(FactoryFields, nil)
		assert.Nil(t, err)

		resp, err := executeRequestAndConvert(context.Background(), req, ListFactoriesResponse{}, client)
		assert.Nil(t, err)
		assert.Len(t, resp.Factories, 1)
		assert.Equal(t, id, resp.Factories[0].ID)
	})

	t.Run("when server returns error", func(t *testing.T) {
		server := getTestServer(t, ServerError, "factory")
		defer server.Close()

		client := NewClient(server.URL, nil)

		req, err := constructListQuery(FactoryFields, nil)
		assert.Nil(t, err)

		_, err = executeRequestAndConvert(context.Background(), req, ListFactoriesResponse{}, client)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "server returned a non-200 status code")
	})
}

type Case int

const (
	SuccessById Case = iota
	SuccessList
	ServerError
)

func getTestServer(t *testing.T, testCase Case, model string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		_, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		switch testCase {
		case SuccessById:
			io.WriteString(w, fmt.Sprintf(`{
				"data": {
					"%s": {
						"id": "test"
					}
				}
			}`, model))
		case SuccessList:
			io.WriteString(w, fmt.Sprintf(`{
				"data": {
					"%s": [
						{
							"id": "test"
						}
					]
				}
			}`, pluralizeModelName(model)))
		case ServerError:
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `Internal Server Error`)
		}
	}))
}
