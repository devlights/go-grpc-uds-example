usage:
	@echo '-----------------------------------------------------------------------------'
	@echo '以下のどれかのタスクを指定します.'
	@echo '  - install-requirements -- grpcを実行するのに必要なものをインストールします.'
	@echo '                              - protoc は プロジェクトディレクトリ直下の bin にインストールされます.'
	@echo '                              - protoc-gen-go は $(go env GOPATH)/bin にインストールされます.'
	@echo '                              - protoc-gen-doc は $(go env GOPATH)/bin にインストールされます.'
	@echo '  - protoc               -- protocを実行します.'
	@echo '                              - protoファイルは protoディレクトリ の下に存在しているとします.'
	@echo '                              - 生成されたgoファイルは internal ディレクトリの下に配置されます.'
	@echo '  - run                  -- サンプルを実行します.'
	@echo '                              - サーバのサンプル は、      cmd/server/main.go に存在しているとします.'
	@echo '                              - クライアントのサンプル は、 cmd/client/main.go に存在しているとします.'
	@echo '-----------------------------------------------------------------------------'
	@echo '[REFERENCES]'
	@echo '  - https://developers.google.com/protocol-buffers/docs/gotutorial'
	@echo '  - https://devlights.hatenablog.com/entry/2020/08/26/130037'
	@echo '  - https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3'
	@echo '-----------------------------------------------------------------------------'
	
install-requirements: _download-protoc _unzip-protoc _locate-protoc _cleanup-tmp _goget-grpc

_download-protoc:
	mkdir -p tmp && \
	cd tmp && \
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.4/protoc-21.4-linux-x86_64.zip --output protoc.zip

_unzip-protoc:
	cd tmp && \
	unzip ./protoc.zip -d protoc

_locate-protoc:
	mkdir -p bin && \
	rm -rf bin/protoc && \
	cd tmp && \
	mv -f ./protoc/ ../bin

_cleanup-tmp:
	rm -rf ./tmp

_goget-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc: _gen-go-out

_gen-go-out:
	mkdir -p internal
	bin/protoc/bin/protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/*.proto

run:
	go run cmd/server/server.go &
	sleep 1
	go run cmd/client/client.go
	sleep 1
	pkill server
	true
