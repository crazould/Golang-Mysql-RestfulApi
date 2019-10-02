package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllUser(w http.ResponseWriter, r *http.Request) {

	db := new(DbHandler)
	rows, _ := db.Query("SELECT * FROM User")

	var users []User

	for rows.Next() {
		var user User

		rows.Scan(&user.ID, &user.Username, &user.Password)

		users = append(users, user)
	}

	bytes, _ := json.Marshal(users)
	fmt.Fprintf(w, string(bytes))
}

func getUser(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	id := Params["id"]

	var users []User
	db := new(DbHandler)
	rows, _ := db.Query(fmt.Sprintf("SELECT * FROM User WHERE ID ='%s'", id))

	for rows.Next() {
		var user User

		rows.Scan(&user.ID, &user.Username, &user.Password)

		users = append(users, user)
	}

	bytes, _ := json.Marshal(users)
	fmt.Fprintf(w, string(bytes))
}

func createUser(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(requestBody, &user)
	db := new(DbHandler)
	db.Query(fmt.Sprintf("INSERT INTO User (Username, Password) VALUES('%s','%s')",
		user.Username,
		user.Password,
	))

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	id := Params["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(requestBody, &user)

	db := new(DbHandler)
	db.Query(fmt.Sprintf("UPDATE User SET Username='%s', Password ='%s' WHERE ID ='%s'",
		user.Username,
		user.Password,
		id,
	))

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	id := Params["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(requestBody, &user)

	db := new(DbHandler)
	db.Query(fmt.Sprintf("DELETE FROM User WHERE ID ='%s'", id))
}

// HandlerRequest ....
func HandlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/users", getAllUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")

	log.Panic(http.ListenAndServe(":2020", router))

}

func main() {
	fmt.Println("Application Start")
	HandlerRequest()
}
