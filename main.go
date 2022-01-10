package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		fmt.Fprint(rw, "Hello World\n")
		d, err := ioutil.ReadAll(r.Body)
		if err != null {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw,"Hello %s ", d)
	})
	http.ListenAndServe(":9090",nil)
}
