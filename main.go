package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serv := &http.Server{
		Addr:    ":1112",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, getKeyServerAddr(), l.Addr().String())
			return ctx
		},
	}

	err := serv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed unexpectedly\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
	cancelCtx()
}
