test:
	gotest -v 

int-test:
	gotest -v --tags=integration

docker-test: 
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

get-gotest:
	go get -u -v github.com/rakyll/gotest
