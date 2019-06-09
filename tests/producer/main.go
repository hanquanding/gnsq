package main

import "gnsq"

func main() {
	gnsq.Publish("127.0.0.1:4150", "test", []byte("hello world"))
}
