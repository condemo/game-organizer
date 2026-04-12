binary-name=game-organizer

build:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./services/rest/cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./services/rest/cmd/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./services/rest/cmd/main.go

run: build
	@./bin/${binary-name}-linux

arm-build:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./services/rest/cmd/main.go

arm-run: arm-build
	@./bin/${binary-name}-arm64

test:
	@go test ./... -v -count=1./services/rest/cmd/main.go

clean:
	@rm -rf ./bin/*
	@go clean



goose-up:
	@goose up
goose-down:
	@goose down

