gen:
	protoc \
	  --go_out=. \
	  --proto_path=api \
	  --go-grpc_out=. \
	  api/*.proto

clean:
	rm -rf ./build/

build: clean
	go build -o ./build/ ./cmd/node

ex1:
	./build/node --port 8301

ex2:
	./build/node --port 8302 --chost localhost --cport 8301

ex3:
	./build/node --port 8303 --chost localhost --cport 8302

ex4:
	./build/node --port 8304 --chost localhost --cport 8301