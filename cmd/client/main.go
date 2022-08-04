package main

import (
	"context"
	"fmt"
	"go-crud-twirp/rpc/twirpAPI"
	"net/http"
	"os"
)

func main() {
	client := twirpAPI.NewTwirpAPIProtobufClient("http://localhost:8080", &http.Client{})

	albums, err := client.GetAlbums(context.Background(), &twirpAPI.GetAlbumsReq{})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Albums: %+v", albums)

	albumById, err := client.GetAlbumByID(context.Background(), &twirpAPI.GetAlbumByIDReq{Id: 1})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Albums: %+v", albumById)
}
