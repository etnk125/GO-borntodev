package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type St struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Sts []St

func init() {
	JSON := `[
		{
			"id":1,
			"name":"c"
		},
		{
			"id":2,
			"name":"py"
		},
		{
			"id":3,
			"name":"goe"
		}
	]`
	err := json.Unmarshal([]byte(JSON), &Sts)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, st := range Sts {
		if highestID < st.ID {
			highestID = st.ID
		}
	}
	return highestID + 1
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	JSON, err := json.Marshal(Sts)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(JSON)
	case http.MethodPost:
		var newSt St
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, newSt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newSt.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newSt.ID = getNextID()
		Sts = append(Sts, newSt)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
func main() {
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)

}
