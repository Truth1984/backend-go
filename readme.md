## setup

```Go
import (
	"io"

	u "github.com/Truth1984/awadau-go"
	un "github.com/Truth1984/backend-go/util"
)

func main() {
	un.Setup("", u.Map("logging", un.ConfigLogger{Level: 0}))
	un.Trace("Logging msg")

	r := un.ServerInit()

	mw1 := func(input un.HttpPkg) {
		input.Local["body"] = input.Req.Body
		input.Next()
	}

	mw2 := func(input un.HttpPkg) {
		bodyBytes, err := io.ReadAll(input.Local["body"].(io.Reader))
		un.ErrorEH(err)
		un.Warn("middleware body-parser test", string(bodyBytes))
		input.Next()
	}

	mws := []func(input un.HttpPkg){mw1, mw2}

	un.ServerPost(r, "/", un.ServerMiddlewareCompile(mws, nil, nil))
	un.ServerStart(r, 3000)

	// send post req. to localhost:3000 to test the result
}
```
