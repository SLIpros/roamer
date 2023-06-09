# gorilla mux router extension

## Install
```go
go get -u github.com/SLIpros/roamer/pkg/gorilla@latest
```

## Example
```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/SLIpros/roamer"
	"github.com/SLIpros/roamer/parser"
	roamerGorilla "github.com/SLIpros/roamer/pkg/gorilla"
	"github.com/gorilla/mux"
)

type Body struct {
	UserID string `path:"user_id"`
}

func main() {
	router := mux.NewRouter()

	r := roamer.NewRoamer(
		roamer.WithParsers(
			parser.NewPath(roamerGorilla.Path),
		),
	)

	router.Use(roamer.Middleware[Body](r))
	router.HandleFunc("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		var body Body
		if err := roamer.ParsedDataFromContext(r.Context(), &body); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods(http.MethodPost)
	http.ListenAndServe(":3000", router)
}
```