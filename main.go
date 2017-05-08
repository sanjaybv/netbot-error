package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	go func() {
		time.Sleep(2 * time.Second)
		panic("This is a purposeful panic.")
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello! This is netbot-error service."))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
