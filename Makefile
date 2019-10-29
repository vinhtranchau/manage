all: client_build  server_build

build: server_build client_build

start: client_start server_start envoy_start

server_build:
							protoc -I proto/ proto/$(input)/$(input).proto   --go_out=plugins=grpc:backend/internal/
							@echo "Server has been built"

server_start:
							go run backend/cmd/management/main.go

client_build:
							protoc --proto_path=proto  --js_out=import_style=commonjs,binary:frontend/src/proto/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:frontend/src/proto/	 proto/$(input)/$(input).proto 
							@echo "Client has been built"

client_start:
							cd frontend; yarn start

envoy_build:
							sudo -E docker build -t envoy:v1 .
							@echo "Docker has been built"

envoy_start:
							sudo docker run  -p 8080:8080 --net=host  envoy:v1
							@echo "Envoy started"
