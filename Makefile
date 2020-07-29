.PHONEY: clean get
default: build
build: get
	 if [ ! -d "./bin/" ]; then mkdir ./bin/; fi
	 env GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o ./bin/memberlist.so ./src/
buildosx: get
	 if [ ! -d "./bin/" ]; then mkdir ./bin/; fi
	 env GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o ./bin/memberlist.so ./src/
get:
	 go mod download
install:
	 cp ./bin/memberlist.so /usr/local/lib/pulseha
clean:
	 go clean -modcache