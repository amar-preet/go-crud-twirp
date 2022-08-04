proto:
	protoc --twirp_out=. --go_out=rpc/twirpAPI/ rpc/twirpAPI/twirp.proto

run:
	go run cmd/server/main.go