package cake

import (
	"net/http"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/user/:id"), []string{"user", ":id"})
	if !ok {
		t.Errorf("parsePattern fail")
	}
	ok = reflect.DeepEqual(parsePattern("/file/*img"), []string{"file", "*img"})
	if !ok {
		t.Errorf("parsePattern fail")
	}
}

func newTestRouter() *router {
	r := newRouter()
	r.addRoute(http.MethodGet, "/", nil)
	r.addRoute(http.MethodGet, "/user/:id", nil)
	r.addRoute(http.MethodGet, "/file/user/*.png", nil)
	return r
}

func TestRouter(t *testing.T) {
	r := newTestRouter()
	route, _ := r.getRoute(http.MethodGet, "/")
	if route == nil {
		t.Fatalf("getRoute fail")
	}
	if route.pattern != "/" {
		t.Errorf("getRoute fail")
	}

	route, m := r.getRoute(http.MethodGet, "/user/123")
	if route == nil {
		t.Fatalf("getRoute fail")
	}
	if route.pattern != "/user/:id" {
		t.Errorf("getRoute fail")
	}
	if m["id"] != "123" {
		t.Errorf("getRoute fail")
	}

	route, m = r.getRoute(http.MethodGet, "/file/user/custom-img/logo.png")
	if route == nil {
		t.Fatalf("getRoute fail")
	}
	if route.pattern != "/file/user/*.png" {
		t.Errorf("getRoute fail")
	}
	if m[".png"] != "custom-img/logo.png" {
		t.Errorf("getRoute fail")
	}
}
