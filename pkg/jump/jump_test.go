package jump

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "io/ioutil"
  "encoding/json"
  "log"
  "bytes"
)

func Test_home(t *testing.T) {
  r := http.NewServeMux()
  r.HandleFunc("/", home)
  ts := httptest.NewServer(r)
  defer ts.Close()

  res, err := http.Get(ts.URL + "/")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
  }
  bodyString := string(b)
  expectString := "/ - Greetings from Golang!"
  if bodyString != expectString {
    t.Errorf("Expected %s received %s", bodyString, expectString)
  }
}

func Test_jump_get(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("/jump", jump)
	ts := httptest.NewServer(r)
  defer ts.Close()
  
  res, err := http.Get(ts.URL + "/jump")
	if err != nil {
	   t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
	   t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }
  
  b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
  }

  var bodyJson AppResponse
  err = json.Unmarshal(b, &bodyJson)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}
  var expectJson AppResponse
  expectJson = AppResponse{Code: 200, Message: "/jump - Greetings from Golang!"}
  log.Println(bodyJson)
  log.Println(expectJson)
  if bodyJson != expectJson {
    t.Errorf("Expected %v received %v", bodyJson, expectJson)
  }
}

func Test_jump_post(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
  
  var a = []string{srv.URL}
  var jump Jump
  jump = Jump{Message: "hi", Last_path: "/last", Jump_path: "/jump", Jumps: a}
  requestBody, err := json.Marshal(jump)
  res, err := http.Post(a[0] + "/jump", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
	   t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
	   t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }
  
  b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
  }

  var bodyJson AppResponse
  err = json.Unmarshal(b, &bodyJson)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}
  var expectJson AppResponse
  expectJson = AppResponse{Code: 200, Message: "/last - Greetings from Golang!"}
  if bodyJson != expectJson {
    t.Errorf("Expected %v received %v", bodyJson, expectJson)
  }
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/last", jumpMock)
	handler.HandleFunc("/jump", jump)
 
	srv := httptest.NewServer(handler)

	return srv
}
 
func jumpMock(w http.ResponseWriter, r *http.Request) {
  var res AppResponse
  res = AppResponse{Code: 200, Message: "/last - Greetings from Golang!"}
  data, err := json.Marshal(res) 
	if err != nil { 
	  panic("Error in Marshal") 
  } 
	_, _ = w.Write(data)
}