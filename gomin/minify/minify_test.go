package minify

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()

	RootHandler(w, req)

	assert.Equal(t, w.Body.String(), "Hi there, I love foo!")
}
