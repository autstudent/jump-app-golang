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

- Execute de App from de root folder

```$bash
$ make run
2020/11/21 20:13:59 Starting server on :8442
```

## Golang Test

Regarding test, it is required execute next command:

```$bash
$ make test-cover
?       github.com/acidonper/golang-demo/cmd/golang-demo        [no test files]
ok      github.com/acidonper/golang-demo/internal/api   (cached)        coverage: 94.3% of statements
```

## Test Demo App API Locally

- GET method to reach /

```$bash
$ curl -X GET localhost:8442/
/ - Greetings from GoLand!
```

- GET method to reach /jump

```$bash
$ curl -X GET localhost:8442/jump
{"code":200,"message":"/jump - Greetings from GoLand!"}
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
{"code":200,"message":"/jump - Greetings from GoLand!"}
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
