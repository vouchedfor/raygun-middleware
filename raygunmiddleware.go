package raygunmiddleware

import (
	"log"
	"net/http"

	"github.com/MindscapeHQ/raygun4go"
)

type Handler struct {
	AppName string
	ApiKey  string
	DevMode bool
}

func (h *Handler) HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.DevMode != true {
			raygun, err := raygun4go.New(h.AppName, h.ApiKey)
			if err != nil {
				log.Println("Unable to create Raygun client:", err.Error())
			}

			defer raygun.HandleError()
		}

		next.ServeHTTP(w, r)
	})
}
