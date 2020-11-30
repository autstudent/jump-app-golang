package jump

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
	"bytes"
	"strconv"
)

type Jump struct {
	Message string
    Last_path string
    Jump_path string
    Jumps []string
}

type AppResponse struct {
	Code int  `json:"code"`
    Message string  `json:"message"`
}


// Index function
func home(w http.ResponseWriter, r *http.Request) {

	// Test only / accepted
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
	}
	
	// return message
    fmt.Fprintf(w, "/ - Greetings from GoLand!")
}

// Jump Function
func jump(w http.ResponseWriter, r *http.Request) {

	// Define custom header to avoid CORS
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Go-Lang-modifier", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, React-modifier")

	// GET Method return a direct Response
	if r.Method == "GET" {
		getResponse := AppResponse{Code: http.StatusOK, Message: "/jump - Greetings from GoLand!"}
		getData, err := json.Marshal(getResponse) 
		if err != nil { 
		  panic("Error in Marshal") 
		}
		fmt.Fprint(w, string(getData))
		return
	}

	// OTION Methods write header
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Logs request
	log.Println(r.Method, "/", r.URL.Path[1:])
	log.Println("Headers ->", r.Header)
	log.Println("Host -> ", r.Host)

	// Parse JSON body
	dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

	// Define a jump variable
	var j Jump

	// Decode object in POST body
	err := dec.Decode(&j)
	if err != nil {
		log.Println("Error decoding Jump Object")
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Define response variables
	var mes string
	var cod int
	var i = j.Jumps

	// Return Error when receive a POST and jumps are not defined
	if len(i) == 0 {
		errorHandler(w, r, http.StatusBadRequest)
		return
	} 

	// Add jump to headers
	var jumpheader = "jump" + strconv.Itoa(len(i))
	w.Header().Add(jumpheader, "Golang")

	// When there is 1 jump 
	if len(i) == 1 { 

		// Define Last URL
		var url = i[0] + j.Last_path

		// Sent GET request to the last jump
		log.Println("GET Calling", url)
		req, err := http.Get(url)
		if err != nil {
			mes = "/jump - Farewell from GoLand! Error jumping " + url
			cod = http.StatusBadGateway
		} else {
			respdec := json.NewDecoder(req.Body)
			respdec.DisallowUnknownFields()
			var res AppResponse
			errdec := respdec.Decode(&res)
			if errdec != nil {
				log.Println("Error decoding Response Object")
				errorHandler(w, r, http.StatusBadRequest)
				return
			}

			// Generate response
			mes = res.Message
			cod = res.Code
		}

	}

	// When there are more than 1 jumps
	if len(i) > 1 { 

		// Define URL and Body
		var url = i[0] + j.Jump_path
		var body = j
		body.Jumps = i[1:]

		// Sent POST request to the last jump
		log.Println("POST Calling", url, "-> Body: ", body)
		requestBody, err := json.Marshal(body)
		req, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			mes = "/jump - Farewell from GoLand! Error jumping " + url
			cod = http.StatusBadGateway
		} else {
			respdec := json.NewDecoder(req.Body)
			respdec.DisallowUnknownFields()
			var res AppResponse
			errdec := respdec.Decode(&res)
			if errdec != nil {
				log.Println("Error decoding Response Object")
				errorHandler(w, r, http.StatusBadRequest)
				return
			}

			// Generate response
			mes = res.Message
			cod = res.Code
		}
	}
	
	// Generate the final response
	response := AppResponse{Code: cod, Message: mes}
	data, err := json.Marshal(response) 
	if err != nil { 
	  panic("Error in Marshal") 
	} 

	log.Println("Sending Response... " + string(data))

	// Sent the final Response
	fmt.Fprint(w, string(data)) 
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, status, " - Not found")
	}
	if status == http.StatusBadRequest {
        fmt.Fprint(w, status, " - Bad Resquest")
    }
}

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/jump", jump)

	port := ":8442"
	log.Println("Starting server on", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)

}
