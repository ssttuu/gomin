package minify

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/foo", nil)
    	w := httptest.NewRecorder()

	RootHandler(w, req)

	assert.Equal(t, w.Body.String(), "Hi there, I love foo!")
}