# Primality Test

This is a relatively simple client-server application that takes integers sent by the client and confirms their primality on the server, then returns the result to the client.

  - Spin up the server
  - Use the client to send values to the server
  - Observe the result in glorious CLI-o-vision, presented as a Boolean value (prime = true or false)

### Technology

* [golang] - the client, server and their associated testing helpers are written in Go.

### Building from source

You can build the binaries from source by the standard means, assuming you have Go installed:

```sh
$ cd client
$ go build
$ cd ../server
$ go build
```

Prebuilt binaries are also included as part of the distribution.

### Testing

Each directory has a testing helper file (server_test.go and client_test.go) that is used to validate the core functions in each component. The following are covered:

##### Server

prime_check - the function used to check whether a given input is prime. 

##### Client

addr_check - the function that checks for valid server IP addresses to connect to
input_check - the function that checks the user-entered input is actually valid

Core functionality testing can be performed as such:

```sh
$ cd client
$ go test
...
$ cd ../server
$ go test
...
```

Code coverage can be roughly estimated by running the following commands:
```sh
$ cd client
$ go test -cover
...
$ cd ../server
$ go test -cover
...
```

PLEASE NOTE: the code coverage scores are low here as I'd have to write fake helpers and handlers for a lot of the existing imported libraries. I'm new to Go - I only started learning it this week, in fact - and am considering this particular task to be outside the scope of the challenge.

### Usage

##### Server 

The server can either be deployed as a Docker container:
```sh
$ cd server
$ docker build -t prim_server:v1 -f ./Dockerfile .
$ docker run -it --name prim_server -p 8000:8000 prim_server:v1
```

or you can run the binary from the command line:
```sh
$ ./server
```

The server has been designed to accept multiple connections and deal with their requests concurrently. Basic information on requests is shown in the terminal window, or in the Docker container logs.

##### Client

The client takes an IP address as an argument, e.g:

```sh
$ ./client 127.0.0.1:8000
```

Once a connection is made, enter non-negative integers above 0 to be tested for primality.

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [golang]: <https://golang.org>

