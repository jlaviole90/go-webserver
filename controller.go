package main

import (
	"fmt"
	"io"
	"net/http"
)

func getKeyServerAddr() Key {
	return "serverAddr"
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		fmt.Printf("could not read body: %s\n", err1)
	}

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
		ctx.Value(getKeyServerAddr()),
		hasFirst, first,
		hasSecond, second, body)

	_, err2 := io.WriteString(w, "Hello world!\n")
	if err2 != nil {
		fmt.Printf("error getting server root test\n")
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /hello request\n", ctx.Value(getKeyServerAddr()))

	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		fmt.Printf("error getting test endpoint\n")
	}
}
