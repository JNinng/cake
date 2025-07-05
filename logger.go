package cake

import "log"

func Logger() HandlerFunc {
	return func(c *Context) {
		log.Printf("%s\n", c.Path)
		c.Next()
	}
}
