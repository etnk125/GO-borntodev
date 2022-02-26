package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func init() {
	userJSON := `[
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
	err := json.Unmarshal([]byte(userJSON), &users)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, user := range users {
		if highestID < user.ID {
			highestID = user.ID
		}
	}
	return highestID + 1
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	JSON, err := json.Marshal(users)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(JSON)

	case http.MethodPost:
		var newUser User
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

		} else {

			err = json.Unmarshal(Bodybyte, newUser)

			if err != nil || newUser.ID != 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {

				newUser.ID = getNextID()
				users = append(users, newUser)
				w.WriteHeader(http.StatusCreated)
			}
			return
		}
	}
}

func findID(ID int) (*User, int) {
	for i, user := range users {
		if user.ID == ID {
			return &user, i
		}
	}
	return nil, 0
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	URLPathSegment := strings.Split(r.URL.Path, "user/")
	ID, err := strconv.Atoi(URLPathSegment[len(URLPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user, index := findID(ID)
	if user == nil {
		http.Error(w, fmt.Sprint("not found user ID :", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	case http.MethodPut:
		var updatedUser User
		Bytebody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bytebody, &updatedUser)
		if err != nil || updatedUser.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user = &updatedUser
		users[index] = *user
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before mw start")
		handler.ServeHTTP(w, r)
		fmt.Println("mw ended")
	})
}
func enableCORSmiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before mw start")

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Method", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Add("Access-Control-Allow-Header", "Accept, Content-Type, Content-Length, Authorization")

		handler.ServeHTTP(w, r)
		fmt.Println("mw ended")
	})
}

func main() {
	userItemHandler := http.HandlerFunc(userHandler)
	userListHandler := http.HandlerFunc(usersHandler)
	http.Handle("/user/", enableCORSmiddlewareHandler(userItemHandler))
	http.Handle("/user", enableCORSmiddlewareHandler(userListHandler))
	http.ListenAndServe(":8080", nil)

}
