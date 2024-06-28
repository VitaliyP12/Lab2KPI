default: out/example

clean:
 rm -rf out

test: *.go
 go test ./...

build: implementation.go cmd/example/main.go
 mkdir -p out
 go build -o out/example ./cmd/example