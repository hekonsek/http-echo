#!/usr/bin/env bash

go build http-echo.go

docker build . -t hekonsek/http-echo
docker push hekonsek/http-echo