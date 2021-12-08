package util

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

func ServerQuick(port int, content string) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	})
	LIP("Server starting, Listening on port " + fmt.Sprint(port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

// mux old global middleware
func ServerMiddleware(mw mux.MiddlewareFunc) {
	r := singleton["router"].(*mux.Router)
	r.Use(mw)
}

func ServerPropeller(f func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{})) {
	p := singleton["propeller"].([]func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{}))
	p = append(p, f)
	singleton["propeller"] = p
}

func ServerAddPath(path string, method []string, handler ...func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{})) {
	r := singleton["router"].(*mux.Router)
	r.Methods(method...).Path(path).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		NEXT := false
		local := make(map[string]interface{})

		setnext := func() {
			NEXT = true
		}

		for _, h := range handler {
			NEXT = false
			h(req, res, setnext, local)
			if !NEXT {
				break
			}
		}

		for _, h := range singleton["propeller"].([]func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{})) {
			NEXT = false
			h(req, res, setnext, local)
			if !NEXT {
				break
			}
		}
	})
}

func setServer() {
	singleton["router"] = mux.NewRouter()
	singleton["propeller"] = []func(req *http.Request, res http.ResponseWriter, next func(), local map[string]interface{}){}
}

func ServerStart() {
	r := singleton["router"].(*mux.Router)
	LIP("Server starting, Listening on port " + fmt.Sprint(ConfigGet("port")))
	http.ListenAndServe(fmt.Sprintf(":%d", ConfigGet("port")), r)
}
