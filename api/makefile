build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/bootstrap main.go
	zip bin/bootstrap.zip ./bin/bootstrap
	rm ./bin/bootstrap

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose