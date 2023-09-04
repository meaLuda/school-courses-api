# help us deploy to AWS

tidy:
	go mod tidy

# build on current file
build: tidy
	env GOOS=linux go build -ldflags="-s -w" -o .

# run build before deploy
deploy_prod: build
	serverless deploy --stage prod 