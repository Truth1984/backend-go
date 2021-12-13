package main

import (
	un "github.com/Truth1984/backend-go/util"
)

func main() {
	// un.Setup("", u.Map("logging", un.ConfigLogger{Level: 0}))

	r := un.ServerInit()

	un.ServerStart(r, 3000)
}
