// +build !aws_lambda

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := Main()

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(":8080", handler)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("terminated", <-errs)
}
