package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func main() {
	result := os.Getenv("PORT")
	if result == "" {
		result = fmt.Sprintf("Port doesn't set %s", result)
		fmt.Println(result)
	}
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+result, nil)
	if err != nil {
		panic(err)
	}
}
