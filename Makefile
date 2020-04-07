GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BIN=bin
CMD=$(BIN)/blackmanor

$(CMD): clean
	go $(GO_BUILD_ENV) build -v -o $(CMD)

clean:
	rm -rf $(BIN)

heroku: $(CMD)
	heroku container:push web