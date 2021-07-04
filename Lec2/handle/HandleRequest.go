package handle

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/mul", mul)
	http.HandleFunc("/div", div)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func add(w http.ResponseWriter, r *http.Request){
	num1 := r.FormValue("a")
	num2 := r.FormValue("b")
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	result := strconv.Itoa(a  + b)
	fmt.Fprintf(w, "Answer: " + result )
}

func sub (w http.ResponseWriter, r *http.Request){
	num1 := r.FormValue("a")
	num2 := r.FormValue("b")
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	result := strconv.Itoa(a  - b)
	fmt.Fprintf(w, "Answer: " + result )
}
func mul (w http.ResponseWriter, r *http.Request){
	num1 := r.FormValue("a")
	num2 := r.FormValue("b")
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	result := strconv.Itoa(a  * b)
	fmt.Fprintf(w, "Answer: " + result )
}

func div (w http.ResponseWriter, r *http.Request){
	num1 := r.FormValue("a")
	num2 := r.FormValue("b")
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	if b  == 0{
		fmt.Fprintf(w, "Cannot be divided by zero")
	}else{
		result := float64(a) / float64(b)
		fmt.Fprintf(w, "Answer: " + strconv.FormatFloat(result, 'f', 2, 32) )
	}
}