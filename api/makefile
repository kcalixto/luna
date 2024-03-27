build:
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/bootstrap main.go
	zip -j bin/bootstrap.zip bin/bootstrap

	sam build --config-file samconfig.toml

clean:
	rm -rf ./bin

deploy: clean build
	sam deploy --debug --config-file samconfig.toml