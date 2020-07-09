vanilla: *.go
	gofmt -w *.go
	go build -o vanilla *.go
