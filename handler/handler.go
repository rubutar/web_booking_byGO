package handler

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
func Devide(w http.ResponseWriter, r *http.Request) {
	x, y := float32(100.0), float32(10.0)
	f, err := devideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "cannot devide by zero")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f devided by %f is %f", x, y, f))
}

func addValues(x, y int) int {
	return x + y
}
func devideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot devided by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template : ", err)
		return
	}
}
