package main

import (
	"fmt"
	"go-crud-twirp/db"
	twirpapi "go-crud-twirp/internal/twirpAPI"
	"go-crud-twirp/rpc/twirpAPI"
	"net/http"
	"os"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Println(db)
	server := &twirpapi.Server{DB: db}
	twirpHandler := twirpAPI.NewTwirpAPIServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
