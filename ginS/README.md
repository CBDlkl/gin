# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/CBDlkl/gin"
	"github.com/CBDlkl/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
