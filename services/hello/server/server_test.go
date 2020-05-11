package server

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	srv := New(1123)
	req := httptest.NewRequest("GET", "http://example.com/hello", nil)
	w := httptest.NewRecorder()
	srv.sayHello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	require.EqualValues(t, "Hello!!", string(body))
}
