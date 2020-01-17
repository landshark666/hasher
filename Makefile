all: hasher hasher.exe

hasher.exe: hasher.go
	GOOS=windows GOARCH=amd64 go build github.com/landshark666/hasher

hasher: hasher.go
	GOOS=linux GOARCH=amd64 go build github.com/landshark666/hasher

