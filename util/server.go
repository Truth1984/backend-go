package util

import (
	"fmt"
	"net/http"

	mm "github.com/Truth1984/mux-middleware"
	mux "github.com/gorilla/mux"
)

func ServerQuick(port int, content string) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	})
	Info("Server starting, Listening on port " + fmt.Sprint(port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func ServerInit() *mux.Router {
	return mux.NewRouter()
}

func ServerHtml(router *mux.Router, routerPath string, filepath string) {
	router.PathPrefix(routerPath).Handler(http.FileServer(http.Dir(filepath)))
}

// mux old global middleware
func ServerMiddleware(router *mux.Router, mw mux.MiddlewareFunc) {
	router.Use(mw)
}

func ServerMiddlewareCompile(entry []func(input mm.HttpPkg), middleware []func(input mm.HttpPkg), propeller []func(input mm.HttpPkg)) func(w http.ResponseWriter, req *http.Request) {
	return mm.Compile(entry, middleware, propeller)
}

func ServerAddPath(router *mux.Router, path string, method []string, f func(http.ResponseWriter, *http.Request)) {
	router.Methods(method...).Path(path).HandlerFunc(f)
}

func ServerGet(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("GET")
}

func ServerPost(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("POST")
}

func ServerPut(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("PUT")
}

func ServerDelete(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("DELETE")
}

func ServerHead(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("HEAD")
}

func ServerOptions(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("OPTIONS")
}

func ServerPatch(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("PATCH")
}

func ServerTrace(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("TRACE")
}

func ServerAny(router *mux.Router, action string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(action, f).Methods("GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH", "TRACE")
}

func ServerStart(router *mux.Router, port int) {
	Info("Server starting, Listening on port " + fmt.Sprint(port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
