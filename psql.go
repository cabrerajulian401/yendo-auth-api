package main
import (

	"database/sql" // allows file to interact with psql database
	"net/http" // For Client Server Communication
	"github.com/lib/pq" // Golang Postegre SQL Driver Package 
	"log" // for logging
	"fmt" // formatting package

)


// code to retrive a PostegreSQL using GET Request 

func retrieveRecord(w, http.ResponseWriter, r *http.Request) {

	if r.method != {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed) // sends a 405 "Method not Allowed"
		return
	}

	// assigning the results to a 'placeholder' for now 

	rowsRs, err := db.Query("SELECT * FROM transactions") // This queries the Server to get all of the rows from the Transaction Database 
	
	if err != nil { // condition for getting an error 

		http.Error(w, http.StatusText(500),http.StatusInternalServerError) // Sends a 500 Request Error "Internal Server Error"
		return

	}


}


