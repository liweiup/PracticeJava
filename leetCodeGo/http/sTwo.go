package main

import "net/http"

func httpr() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go"))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err.Error())
	}
}
