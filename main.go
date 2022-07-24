package main

import (
	"fmt"
	"net/http"
)

const (
	LISTEN_PORT = 1212
)

func main() {

	LogCode("Starting web server at Port", fmt.Sprintf("http://127.0.0.1:%d/", LISTEN_PORT))
	http.ListenAndServe(fmt.Sprintf(`:%d`, LISTEN_PORT), nil)

}
