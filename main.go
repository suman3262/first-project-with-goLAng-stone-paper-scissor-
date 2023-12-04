package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rock-paper/rps"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

//enable CORS
func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

//all controler recive a request & responce like express server
func homePage(w http.ResponseWriter , r *http.Request){
	//we server the html page g=from backend

  renderTemplate(w,"index.html")
	
}

func playRound(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("call playRound function ")

	vars := mux.Vars(r)
    choiceNumberStr := vars["choiceNumber"]

    choiceNumber, err := strconv.Atoi(choiceNumberStr)
    if err != nil {
        // Handle error: Invalid input provided in the URL for choiceNumber
        http.Error(w, "Invalid choiceNumber provided", http.StatusBadRequest)
        return
    }

	result := rps.PlayRound(choiceNumber)

	//to send result in json format******
    out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}
	//to set responce type in json , that browser identify that a json has recived
	w.Header().Set("Content-Type", "application/json")

	//to return json responce
	w.Write(out)
}

func main(){
	const port = "4000"

	r := mux.NewRouter()

	r.HandleFunc("/", homePage)
	r.HandleFunc("/play/{choiceNumber}", playRound).Methods("GET")

	log.Println("Server started on port:", port)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

//this function server the HTML page , diffrent diffrent
func renderTemplate(w http.ResponseWriter , page string){

	// recived template and error

	t,err:=template.ParseFiles(page);

	if err !=nil{
		log.Println(err);
	}

	//templae has a built function *Execute() that takes 2params 1>responce , 2>nil
	err=t.Execute(w,nil)

	if err !=nil{
		log.Println(err);
	}
}