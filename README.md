# SimpleKV: A simple key-value store

SimpleKV is a simple key-value store written in Goland. It is a simple implementation of a key-value store that supports the following operations:

- `Put(key, value)`: Set the value of a key.
- `Get(key)`: Get the value of a key.
- `Delete(key)`: Delete a key.
- `Update(key, value)`: Update the value of a key.

The key-value store is implemented using a hash map. The key-value pairs are stored in a map where the key is the key and the value is the value. The key-value store is thread-safe and supports concurrent access.

## Server

The project comes with a simple server that exposes the key-value store over HTTP. The server supports the following endpoints:

- `POST /`: Set the value of a key. The key and value should be passed as query parameters.
- `GET /`: Get the value of a key. The key should be passed as a query parameter. If the key does not exist, the server will return a `404 Not Found` response.
- `DELETE /`: Delete a key. The key should be passed as a query parameter. If the key does not exist, the server will return a `404 Not Found` response.
- `PUT /`: Update the value of a key. The key and value should be passed as query parameters. If the key does not exist, the server will return a `404 Not Found` response.

## Usage

To run the server, use the makefile. Just run `make build` to build the server and `make run` to run it. The server will start on port 8080 by default.

## Motivation

This is a simple project and my first contact with Go. The goal here is to pave
the way for building a production-ready Distributed Cache solution with Go. This
project was crutial to solidify the basics of Go, such as concurrency, error handling, HTTP servers and testing. The next project will be a more complex solution, with more features and a better architecture, and the last project will be a distributed solution.
