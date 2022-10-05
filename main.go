package main

import (
	"fmt"
	"net/http"
	"web_booking_byGO/handler"
)

const portNumb = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/devide", handler.Devide)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumb))
	_ = http.ListenAndServe(portNumb, nil)
}
