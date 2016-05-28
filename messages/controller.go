package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MainController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		FetchMessages(w, req)
	case "POST":
		PostMessages(w, req)
	}
}

//database connection

/*
FetchMessages ...
Fetch messages from the database
*/
func FetchMessages(w http.ResponseWriter, r *http.Request) {

	db, err := getDbConnection()
	checkErr(err)
	//QUERY DATA

	rows, err := db.Query("SELECT email, msg_body, pub_date FROM Messages ORDER BY pub_date DESC")
	checkErr(err)

	messages := Messages{}

	for rows.Next() {

		msg := Message{}

		err = rows.Scan(&msg.Email, &msg.MsgBody, &pubDate)
		checkErr(err)

		msg.PubDate = formatTime(pubDate)

		messages = append(messages, msg)

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		checkErr(err)
	}
}

/*
PostMessages ...
Inserts messages to the database
*/
func PostMessages(w http.ResponseWriter, req *http.Request) {

	email := req.FormValue("email")
	msgBody := req.FormValue("msg")

	db, err := getDbConnection()
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO Messages(email, msg_body) VALUES(?,?)")
	checkErr(err)

	_, err = stmt.Exec(email, msgBody)
	checkErr(err)

	fmt.Println(email)
	fmt.Println(msgBody)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
