# raygun-middleware
Go middleware for Raygun

## Installing

````bash
go get github.com/vouchedfor/raygun-middleware
````
## Usage

````go

package main

import (
  "fmt"
  "net/http"
  
  "github.com/vouchedfor/raygun-middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi there!")
}

func main() {
  raygunHandler := raygunmiddleware.NewHandler("app_name", "api_key", false)

  http.Handle("/", raygunHandler.HandleRequest(http.HandlerFunc(handler)))
  http.ListenAndServe(":8080", nil)
}

````