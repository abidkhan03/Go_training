package change_route

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()

	data, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	// print the json data
	fmt.Println(string(data))
	//convert into unmarshal
	data1 := &Request{}
	err = json.Unmarshal(data, data1)
	if err != nil {
		panic(err)
	}
	fmt.Println(data1.Name)
	
	response := &Response{
		Code:      200,
		Message:   "Welcome " + req.Name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	// convert into unmarshal json
	json.NewEncoder(w).Encode(response)
}

