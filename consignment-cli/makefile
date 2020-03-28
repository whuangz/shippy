build:
	GOOS=linux GOARCH=amd64 go build -o shippy-consignment-cli
	docker build -t shippy-consignment-cli .

run:
	docker run shippy-consignment-cli