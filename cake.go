package cake

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

// ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
		return
	}
	_, _ = fmt.Fprint(w, "404 page not found")
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router[method+"-"+pattern] = handler
}

// GET 添加 GET 请求处理器
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodGet, pattern, handler)
}

// POST 添加 POST 请求处理器
func (engine Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPost, pattern, handler)
}

// Run 运行
func (engine Engine) Run(addr string) error {
	return http.ListenAndServe(addr, &engine)
}
