# go-grpc-uds-example
gRPC with Unix domain socket (UDS) example by golang

It can be executed using [go-task](https://taskfile.dev/#/).

```sh
$ task --list
task: Available tasks for this project:
* install-requirements: install requirements
* protoc:               gen protoc
* run:                  run
```

## Install gRPC and Go libraries

```sh
task install-requirements
```

## Run protoc

```sh
task protoc
```

## Run Server and Client

```sh
task run
```

## Example

```sh
$ task install-requirements
task: [install-requirements] mkdir tmp
task: [_download-protoc] curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip --output protoc.zip
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 1548k  100 1548k    0     0  4254k      0 --:--:-- --:--:-- --:--:-- 4254k
task: [_unzip-protoc] unzip ./protoc.zip -d protoc
Archive:  ./protoc.zip
  inflating: protoc/bin/protoc       
  inflating: protoc/include/google/protobuf/any.proto  
  inflating: protoc/include/google/protobuf/api.proto  
  inflating: protoc/include/google/protobuf/compiler/plugin.proto  
  inflating: protoc/include/google/protobuf/descriptor.proto  
  inflating: protoc/include/google/protobuf/duration.proto  
  inflating: protoc/include/google/protobuf/empty.proto  
  inflating: protoc/include/google/protobuf/field_mask.proto  
  inflating: protoc/include/google/protobuf/source_context.proto  
  inflating: protoc/include/google/protobuf/struct.proto  
  inflating: protoc/include/google/protobuf/timestamp.proto  
  inflating: protoc/include/google/protobuf/type.proto  
  inflating: protoc/include/google/protobuf/wrappers.proto  
  inflating: protoc/readme.txt       
task: [install-requirements] mkdir -p bin
task: [install-requirements] rm -rf bin/protoc
task: [_locate-protoc] mv -f ./protoc/ ../bin
task: [install-requirements] rm -rf ./tmp
task: [install-requirements] go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
task: [install-requirements] go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
task: [install-requirements] rm -rf ./text-files; git clone --quiet https://github.com/devlights/text-files


$ task protoc
task: [protoc] mkdir -p internal
task: [protoc] bin/protoc/bin/protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/*.proto


$ task run
task: [run] go run cmd/server/server.go &
task: [run] sleep 1
task: [run] go run cmd/client/client.go
data:"HELLO WORLD"
data:"GOLANG"
data:"GOROUTINE"
data:"THIS PROGRAM RUNS ON CROSTINI"
task: [run] sleep 1
task: [run] pkill server
signal: terminated
```