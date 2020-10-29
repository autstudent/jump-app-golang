package api

import (
    "fmt"
	"net/http"
	"log"
	"encoding/json"
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


func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "/ - Greetings from GoLand!")
}

func jump(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Go-Lang-modifier", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, React-modifier")

	// Test POST method
	if r.Method == "GET" {
		getResponse := AppResponse{Code: http.StatusOK, Message: "/jump - Greetings from GoLand!"}
		getData, err := json.Marshal(getResponse) 
		if err != nil { 
		  panic("Error in Marshal") 
		}
		fmt.Fprint(w, string(getData))
		return
	}

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

	var j Jump

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

	// Make calls
	if len(i) == 0 {
		errorHandler(w, r, http.StatusBadRequest)
		return
	} 

	if len(i) == 1 { 

		var url = i[0] + j.Last_path

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

			mes = res.Message
			cod = res.Code
		}

	}

	if len(i) > 1 { 
		mes = "More than one jump is not supported right"
		cod = http.StatusOK
	}
	
	// Generate the response
	response := AppResponse{Code: cod, Message: mes}
	data, err := json.Marshal(response) 
	if err != nil { 
	  panic("Error in Marshal") 
	} 

	log.Println("Sending... " + string(data))
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
