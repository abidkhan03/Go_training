#!/bin/sh

go build -o bin/server main.go
go build -o bin/server mycsv/csv.go
