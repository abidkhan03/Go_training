#!/bin/sh

# if a user write -b or --build, then it build the project
if [ "$1" = "-b" ] || [ "$1" = "--build" ]; then
    go build -o bin/server main.go
fi

./bin/server
