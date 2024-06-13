gen:
	protoc \
	  --go_out=. \
	  --proto_path=api \
	  --go-grpc_out=. \
	  api/*.proto
preBuild:
	go mod tidy
clean:
	rm -rf ./build/

build: clean preBuild
	go build -o ./build/ ./cmd/node
	go build -o ./build/ ./cmd/cli



cli:
	./build/cli --host localhost --port 8301

ex1:
	./build/node --port 8301

ex2:
	./build/node --port 8302 --chost localhost --cport 8301 --repl 2

ex3:
	./build/node --port 8303 --chost localhost --cport 8302 --repl 2

ex4:
	./build/node --port 8304 --chost localhost --cport 8301