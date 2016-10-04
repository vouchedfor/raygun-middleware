package raygunmiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockErrorHandler struct {
	t           *testing.T
	panicThrown bool
}

func (mockErrorHandler mockErrorHandler) HandleError() error {
	if r := recover(); r == nil && mockErrorHandler.panicThrown == true {
		mockErrorHandler.t.Error("Error handler should pick this panic")
	}

	return nil
}

func TestHandleRequestWithoutPanic(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Response from upstream"))
	})

	handler := NewHandler("test", "test", false)
	handler.raygunClient = mockErrorHandler{t: t, panicThrown: false}

	raygunHandler := handler.HandleRequest(mockHandler)
	raygunHandler.ServeHTTP(w, r)
}

func TestHandleRequestWithPanic(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("oooooops")
	})

	handler := NewHandler("test", "test", false)
	handler.raygunClient = mockErrorHandler{t: t, panicThrown: true}

	raygunHandler := handler.HandleRequest(mockHandler)
	raygunHandler.ServeHTTP(w, r)
}
