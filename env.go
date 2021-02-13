package main

import "os"

func readVars() {
	var ok bool

	bindAddress, ok = os.LookupEnv("RSS_SERVER_BIND_ADDRESS")
	if !ok {
		bindAddress = "localhost:26009"
	}
}
