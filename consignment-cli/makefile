build:
	GOOS=linux GOARCH=amd64 go build -o consignment-cli
	docker build -t consignment-cli .

run:
	docker run consignment-cli