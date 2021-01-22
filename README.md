# Golang Demo

## Introduction

Golang Demo is one of a set of microservices, named Jumps, developed to generate a microservice communication test tool in order to support multi hands-on and webinars around microservices in Kubernetes.

## Quick Start Golang Demo

Once the demo project has been uploaded, it is required to execute the following process:

- Install Golang

```$bash
# MacOS
$ brew install golang

# Fedora
$ dnf install golang
```

- Execute the App from the root folder

```$bash
$ make build
go build -o bin/golang-demo -race cmd/main.go
$ make run
bin/golang-demo
2020/11/30 22:07:25 Starting server on :8442
```

## Golang Test

Regarding test, it is required execute next command:

```$bash
$ make test
go test ./...
?       github.com/acidonper/golang-demo/cmd    [no test files]
ok      github.com/acidonper/golang-demo/pkg/jump       3.095s
```

## Test Demo App API Locally

- GET method to reach /

```$bash
$ curl -X GET localhost:8442/
/ - Greetings from Golang!
```

- GET method to reach /jump

```$bash
$ curl -X GET localhost:8442/jump
{"code":200,"message":"/jump - Greetings from Golang!"}
```

- POST method with JUMP Object in the body to make multi jumps through Golang Demo

```$bash
$ curl -XPOST -H "Content-type: application/json" -d '{
    "message": "Hello",
    "last_path": "/jump",
    "jump_path": "/jump",
    "jumps": [
        "http://localhost:8442",
        "http://localhost:8442",
        "http://localhost:8442",
        "http://localhost:8442",
        "http://localhost:8442",
        "http://localhost:8442"
    ]
}' 'localhost:8442/jump'
{"code":200,"message":"/jump - Greetings from Golang!"}
```

- POST to index page

```$bash
$ curl -XPOST -H "Content-type: application/json" -d '{
    "message": "Hello",
    "last_path": "/users",
    "jump_path": "/jump",
    "jumps": [
        "http://localhost:8443",
    ]
}' 'localhost:8442/jump'
```

## Author Information

Asier Cidon @Red Hat

asier.cidon@gmail.com
