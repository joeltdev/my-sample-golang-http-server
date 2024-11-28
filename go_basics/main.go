package main

import (
	"errors"
	"fmt"
	"net/http"
)

const port = ":8080"

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! this is the home page")
}

func Aboutpage(w http.ResponseWriter, r *http.Request) {
	sum := addvalues(2, 3)
	fmt.Fprintf(w, "The sum is %d", sum)
}

func addvalues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := dividevalus(4, 0)

	if err != nil {
		fmt.Fprintf(w, "Can't divide by zero: %d", 0)
		return
	}

	fmt.Fprintf(w, "%d divided by %d is %d", 4, 2, f)
	return
}

func dividevalus(x, y int) (int, error) {
	if y <= 0 {
		err := errors.New("denominator must be greater than zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/about", Aboutpage)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting server on port %s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
