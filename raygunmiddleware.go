package raygunmiddleware

import (
	"log"
	"net/http"

	"github.com/MindscapeHQ/raygun4go"
)

type ErrorHandler interface {
	HandleError() error
}

type Handler struct {
	raygunClient ErrorHandler
}

func NewHandler(appName, apiKey string, silentMode bool) Handler {
	raygunClient, err := raygun4go.New(appName, apiKey)
	if err != nil {
		log.Println("Unable to create Raygun client:", err.Error())
	}
	raygunClient.Silent(silentMode)

	return Handler{
		raygunClient: raygunClient,
	}
}

func (h *Handler) HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer h.raygunClient.HandleError()

		next.ServeHTTP(w, r)
	})
}
