package cake

import (
	"net/http"
	"sync"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	groups     []*RouterGroup
	router     *router
	ctxFactory ContextFactory
}

type poolContextFactor struct {
	pool *sync.Pool
}

func (factor *poolContextFactor) Get(w http.ResponseWriter, req *http.Request) *Context {
	c := factor.pool.Get().(*Context)
	updateContext(c, w, req)
	return c
}

func (factor *poolContextFactor) Put(c *Context) {
	factor.pool.Put(c)
}

func New() *Engine {
	pool := &sync.Pool{
		New: func() any {
			return new(Context)
		},
	}
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.ctxFactory = &poolContextFactor{pool: pool}
	return engine
}

// Run 运行
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.ctxFactory.Get(w, req)
	defer engine.ctxFactory.Put(c)
	engine.router.handle(c)
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		engine: engine,
		parent: g,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodGet, pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPost, pattern, handler)
}

// addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加 GET 请求处理器
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodGet, pattern, handler)
}

// POST 添加 POST 请求处理器
func (engine Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPost, pattern, handler)
}
