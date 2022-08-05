## Generate CRUD API interface, client and server

```
make proto
```

## Run the Server
```
make run
```

## Run the Client
```
make client
```


## Testing the endpoints via JSON content-type

### GetAlbums
```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbums
```


### GetAlbumByID
```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "5"}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbumByID
```

### PostAlbums
```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"title": "Sunrise","artist": "Tiesto","price": 9}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/PostAlbums
```

### DeleteAlbumByID
```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "1"}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/DeleteAlbumByID
```

### UpdateAlbumByID
```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "5", "title": "Legend","artist": "Sidhu Moosewala","price": 79}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/UpdateAlbumByID
```
