# Simple http server in Golang

### Description
The main.go file to run the httpserver reqeust with a simple hello route to listen on port 8080
and run the server with the command

```go run main.go``` 

If Go is not installed then

```sudo apt-get install golang-go```


In scripts folder build.sh and run.sh to build and run the program with the command

#### Build

```./scripts/build.sh``` 

If permission is required then

```chmod +x scripts/build.sh```


#### Run

```./scripts/run.sh```

If permission is required then

```chmod +x scripts/run.sh```

or 

```go run main.go```

http://localhost:8000/parse

Copy the url and paste it into Postman, while sending url select the POST method and then select the ```Body >> raw```
after selecting raw body give a path to the cvs file like 
```{
        "path": "mycsv/record.csv"
   }
```
or 
open the docs folder to run that file in Postman.

In Postman 


### Author 
* **Abdullah Abid** 