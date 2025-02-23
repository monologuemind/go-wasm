package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	handler := http.FileServer(http.Dir(dir))
	http.Handle("/", handler)

	fmt.Println("Serving", dir, "on port", port)
	fmt.Println("Open in your browser: http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
