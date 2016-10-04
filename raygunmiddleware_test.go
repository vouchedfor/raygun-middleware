package raygunmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vouchedfor/raygun-middleware"
)

func TestHandleRequest(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Response from upstream"))
	})

	handler := raygunmiddleware.Handler{DevMode: true}
	raygunHandler := handler.HandleRequest(mockHandler)
	raygunHandler.ServeHTTP(w, r)
}
