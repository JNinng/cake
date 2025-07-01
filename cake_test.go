package cake

import (
	"net/http"
	"reflect"
	"testing"
	"unsafe"
)

func TestEngine(t *testing.T) {
	engine := New()
	engine.GET("/v1/user/:id", func(c *Context) {})
	orderGroup := engine.Group("/v2")
	{
		orderGroup.GET("/user/:id", func(c *Context) {})
	}

	engineValue := reflect.ValueOf(engine).Elem()
	routerField := engineValue.FieldByName("router")
	if !routerField.IsValid() {
		t.Fatalf("routerField is not valid")
	}
	addr := unsafe.Pointer(routerField.UnsafeAddr())
	routerPtr := (**router)(addr)
	r := *routerPtr

	route, m := r.getRoute(http.MethodGet, "/v1/user/123")
	if route == nil {
		t.Fatalf("getRoute fail")
	}
	if route.pattern != "/v1/user/:id" {
		t.Errorf("getRoute fail")
	}
	if m["id"] != "123" {
		t.Errorf("getRoute fail")
	}

	route, m = r.getRoute(http.MethodGet, "/v2/user/1234")
	if route == nil {
		t.Fatalf("getRoute fail")
	}
	if route.pattern != "/v2/user/:id" {
		t.Errorf("getRoute fail")
	}
	if m["id"] != "1234" {
		t.Errorf("getRoute fail")
	}
}
