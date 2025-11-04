# What is this repository

This repository holds simple REST API for todos. Todos are saved in MYSQL database

# Prerequisities

You need to have go installed on your system and have running mysql instance with db named `todos` and root password `password`.

**Hint**: Easiest way to spin up mysql instance is by using docker. If you have it installed on your system simply run: `docker run -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql`.

# How to run

In order to run the api execute the following command (assuming you are in the repo root dir): `go build ./cmd/server/main.go` and then `go run ./cmd/server/main.go`.

**Hint**: If you are using vscode you can install REST Client extension and open up [requests.rest](./requests.rest) file which has sample http requests to try out the api
