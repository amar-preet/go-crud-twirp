proto:
	protoc --twirp_out=. --go_out=rpc/twirpAPI/ rpc/twirpAPI/twirp.proto

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

sql:
	sqlc generate

