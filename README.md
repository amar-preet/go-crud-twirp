# go-crud-twirp

A simple CRUD API in `go` using [twirp](https://twitchtv.github.io/twirp/docs/intro.html) with storage on postgreSQL

## Install Dependencies

Make sure local PostgreSQL database is running. 
Then Run,
```
go mod tidy
```


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


## Testing the endpoints


### Using Protobuf

<details>
<summary>
GET Albums
</summary>

```
echo  \
    | protoc --encode twirpAPI.GetAlbumsReq ./rpc/twirpAPI/twirp.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbums \
    | protoc --decode twirpAPI.GetAlbumsResp  ./rpc/twirpAPI/twirp.proto
```
</details>

<details>
<summary>
GetAlbumByID
</summary>

```
echo  'id:2' \
    | protoc --encode twirpAPI.GetAlbumByIDReq ./rpc/twirpAPI/twirp.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbumByID \
    | protoc --decode twirpAPI.GetAlbumByIDResp  ./rpc/twirpAPI/twirp.proto
```
</details>


<details>
<summary>
Post Album
</summary>

```
echo  'title:"Invinsible",artist:"Paul Oakenfold",price:19' \
    | protoc --encode twirpAPI.PostAlbumsReq ./rpc/twirpAPI/twirp.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/twirp/twirpAPI.TwirpAPI/PostAlbums \
    | protoc --decode twirpAPI.PostAlbumsResp  ./rpc/twirpAPI/twirp.proto
```
</details>

<details>
<summary>
Delete Album
</summary>

```
echo 'id:9' \
    | protoc --encode twirpAPI.DeleteAlbumByIDReq ./rpc/twirpAPI/twirp.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/twirp/twirpAPI.TwirpAPI/DeleteAlbumByID \
    | protoc --decode twirpAPI.DeleteAlbumByIDResp  ./rpc/twirpAPI/twirp.proto
```
</details>

<details>
<summary>
Update Album By ID
</summary>

```
echo  'id:12,title:"Going Home",artist:"Drake",price:45' \
    | protoc --encode twirpAPI.UpdateAlbumByIDReq ./rpc/twirpAPI/twirp.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/twirp/twirpAPI.TwirpAPI/UpdateAlbumByID \
    | protoc --decode twirpAPI.UpdateAlbumByIDResp  ./rpc/twirpAPI/twirp.proto
```
</details>

### Using JSON


<details>
<summary>
GET Albums
</summary>

```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbums
```
</details>

<details>
<summary>
GET Album By ID
</summary>

```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "5"}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/GetAlbumByID
```
</details>

<details>
<summary>
POST an Album
</summary>

```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"title": "Sunrise","artist": "Tiesto","price": 9}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/PostAlbums
```
</details>

<details>
<summary>
DELETE Album By ID
</summary>

```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "1"}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/DeleteAlbumByID
```
</details>

<details>
<summary>
UPDATE Album By ID
</summary>

```
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"id": "5", "title": "Legend","artist": "Sidhu Moosewala","price": 79}' \
    http://localhost:8080/twirp/twirpAPI.TwirpAPI/UpdateAlbumByID
```
</details>
