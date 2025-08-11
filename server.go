// Yendo Mini-Project



package main

import (
	
	"encoding/json" // Importing JSON package to marshal Go Data Structures
	"log" 		    // Importing Log package 
	"net/http"		// HTTP package to create both clients and servers
	"time"
	"github.com/google/uuid" // Useful for measuring and displaying time. 
	"fmt" // FMT Package that provides printing functions 

)
ff 

// This is our data structure that serves for storing the attributes of a Credit Card Authorization

// JSON Tages will tell Go how to convert this struct to and from JSON

type AuthorizationRequest struct {

	ID string `json:"id"` 
	CardNumber string `json:"cardNumber"`
	Amount float64 `json:"amount"`
	Timestamp int64 `json:"timestamp"`

}


// Now to initiate our in-memory databse that is dynamic using our Struct

var authorizations = []AuthorizationRequest{}

// Note that the slice authorization is a slice being a collection of the Struct we declared above



// Creating our function for GET route /authorizations

func getAuthorizations( w http.ResponseWriter, r *http.Request) {

	// control logic to only handle GET request here
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Handling getAuthorization Request")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorizations)

}



// Creating our function for the POST route /

func handleAuthorize( w http.ResponseWriter, r *http.Request) {

	// control logic to only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Handling handleAuthorize request")

	var newAuth AuthorizationRequest // referencing Authorization Request Struct to put into new strcut

	if err := json.NewDecoder(r.Body).Decode(&newAuth); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}


	newAuth.ID = uuid.New().String()
	newAuth.Timestamp = time.Now().Unix()

	authorizations = append(authorizations,newAuth) // appending elements of another slice, note that a slice is one specific instance of an underlying array

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // sends in a 200 for Successin POST Request

	json.NewEncoder(w).Encode(newAuth) // returns full jSON object with Customer Data




}





// Create main function which is our entry to application

func main() {
 
	http.HandleFunc("/authorize",handleAuthorize)
	http.HandleFunc("/authorizations", getAuthorizations)
	log.Println("Starting on :8080") // log on console that server has been initialized 
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}




}





