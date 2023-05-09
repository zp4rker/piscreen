package main

import "piscreen/keys"

func main() {
	go keys.Listen("13")
	go keys.Listen("6")

	for {
	}
}
