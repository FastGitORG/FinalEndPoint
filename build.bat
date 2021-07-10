SET GOARCH=amd64
SET GOOS=linux
go build -ldflags="-s -w" -o fgep .\src