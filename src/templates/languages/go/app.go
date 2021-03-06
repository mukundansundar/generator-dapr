package main
 
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
    "log"
	"net/http"
	"fmt"
	"math/rand"
	"os"
	"github.com/gorilla/mux"
)

type body struct {
    Number int `json:"number,int"`
}

// daprState represents a state in Dapr.
type daprState struct {
	Key   string    `json:"key,omitempty"`
	Value int `json:"value,omitempty"`
}

var daprPort = os.Getenv("DAPR_HTTP_PORT")
var stateURL = fmt.Sprintf("http://localhost:%s/v1.0/state", daprPort)

// Routes - note that these are case sensitive
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/randomNumber", randomNumber).Methods("GET", "OPTIONS")
	router.HandleFunc("/savedNumber", getSavedNumber).Methods("GET", "OPTIONS")
	router.HandleFunc("/saveNumber", saveNumber).Methods("POST", "OPTIONS")
	router.HandleFunc("/dapr/subscribe", subscribe).Methods("GET", "OPTIONS")
	router.HandleFunc("/A", topicAHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/B", topicBHandler).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6000", router))
}

func randomNumber(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rand.Intn(100))
}

func getSavedNumber(w http.ResponseWriter, r *http.Request) {
	savedNumberURL := fmt.Sprintf("%s/savedNumber", stateURL)
	fmt.Print(savedNumberURL)
	response, err := http.Get(savedNumberURL)
	if err != nil {
        fmt.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
        } else {
			json.NewEncoder(w).Encode(string(contents))
		}
	}
}

func saveNumber(w http.ResponseWriter, r *http.Request) {
	var body body
	json.NewDecoder(r.Body).Decode(&body)
	state := daprState{
		Key:	"savedNumber",
		Value:	body.Number}

	states := []daprState{state}

	jsonValue, err := json.Marshal(states)

	if err != nil {
		log.Printf("Could not save states in Dapr: %s", err.Error())
	}

	log.Printf("Posting state to %s with '%s'", stateURL, jsonValue)
	res, err := http.Post(stateURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Could not save states: %s", err.Error())
	} else {
		res.Body.Close()
	}
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	topics := [2]string{"A", "B"}
	json.NewEncoder(w).Encode(topics)
}

func topicAHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Got message of topic 'A'");
	var jsonValue map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonValue)
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println(jsonValue);
	}
}

func topicBHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Got message of topic 'B'");
	var jsonValue map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonValue)
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println(jsonValue);
	}
}
