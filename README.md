# url-shortener
URL shortener

go mod init url-shortener

go get -u github.com/gorilla/mux
go get -u github.com/satori/go.uuid

go run main.go

curl -X POST -H "Content-Type: application/json" -d @url.json http://localhost:8080/shorten