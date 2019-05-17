all: build

build:
	go build http-echo.go

docker: build
	docker build . -t hekonsek/http-echo

docker-push: docker
	docker push hekonsek/http-echo