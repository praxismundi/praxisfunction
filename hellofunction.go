package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprintf(w, "Hello, world!")
		return
	}
	if d.Name == "" {
		fmt.Fprintf(w, "Hello, world!")
		return
	}
	fmt.Fprintf(w, "Hello, %s", html.EscapeString(d.Name))
	return
}
