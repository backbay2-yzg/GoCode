package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloWorld"))
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8006", nil)
}
