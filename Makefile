vanilla: *.go
	gofmt -w *.go
	go build -o vanilla *.go

test: *.go
	gofmt -w *.go
	go test -v
