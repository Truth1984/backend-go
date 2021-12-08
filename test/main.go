package main

import (
	"net/http"

	u "github.com/Truth1984/awadau-go"
	un "github.com/Truth1984/backend-go/util"
)

func pm(line string, proceed bool) func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{}) {
	return func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{}) {
		u.Print(line)
		if proceed {
			next()
		}
	}
}

func main() {
	un.Setup("", u.Map("loglevel", 0, "server", true))
	un.ServerAddPath("/", []string{"GET"}, pm("1", true), pm("2", false), pm("3", true))
	un.ServerPropeller(pm("4", true))
	un.ServerPropeller(pm("5", false))
	un.ServerPropeller(pm("6", true))

	un.ServerStart()
}
