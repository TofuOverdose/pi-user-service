#!/bin/bash

cd "./internal/users"
# todo move to test.env and parse it here
export MONGO_URI="mongodb://localhost:27017" MONGO_DATABASE="test-user-service"
go test ./...