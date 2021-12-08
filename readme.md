## setup

```Go
import un "github.com/Truth1984/backend-go/util"

func main() {
	un.Setup("",nil)
	un.LWP("warning test")
}
```

## server setup

```Go
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
	// listening on port 3000, visiting localhost:3000 println:
	// 1 \n 2 \n 4 \n 5
}
```

## process

1. use `un.ServerMiddleware` for old `MiddlewareFunc`, which will run first for all routes

2. use `un.ServerAddPath` to add path, and define path specific middleware

3. use `un.ServerPropeller` for global propeller
