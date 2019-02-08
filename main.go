package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //Takes two params, the root and the call back function.
		w.Write([]byte("Hello Go!"))
	})
	http.ListenAndServe(":8000", nil)
}
