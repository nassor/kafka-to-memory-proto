all:

proto:
	@protoc -I/usr/local/include --go_out=plugins=grpc:./api --proto_path=./api/proto/ ./api/proto/*.proto

clean:
	@rm -f ./build/event-receiver
	@rm -f ./build/device

build: clean
	@mkdir -p ./build
	@go build -o ./build/event-receiver ./services/event-receiver/cmd
	@go build -o ./build/device ./services/device/cmd