
all: test build

test:
	gotest -v ./...

build:
	go build -o app -v

intTest:
	go test -v --tags=integration ./service/...

lint:
	golangci-lint run

jwt-mac:
	go build -o tokenGen ./cmd/jwt

jwt-CI:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tokenGenCI ./cmd/jwt

cpenv:
	cp .env ./cmd/jwt/.

emg1Deploy:
	cp ./builds/cloudbuild_1.yml ./cloudbuild.yml
	git add --all
	git commit -m 'emergency build #1'
	git push
	cp ./builds/cloudbuild.yml ./cloudbuild.yml

emg2Deploy:
	cp ./builds/cloudbuild_2.yml ./cloudbuild.yml
	git add --all
	git commit -m 'emergency build #2'
	git push
	cp ./builds/cloudbuild.yml ./cloudbuild.yml
