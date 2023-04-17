package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "localhost:1313"
	fmt.Printf("> Server running on http://%s\n", addr)
	http.ListenAndServe(addr, http.FileServer(http.Dir("./static")))
}