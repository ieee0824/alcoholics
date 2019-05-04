# alcoholics

alcoholics is golang gin Web Framework Middleware.  
This is a resiliency tool that helps applications tolerate random instance failures.

```Go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ieee0824/alcoholics"
)

func main() {
	r := gin.Default()
	r.Use(alcoholics.New([]alcoholics.Option{
		{
			Probability: 5,
			StatusCode:  200,
		},
		{
			Probability:  1,
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "なんか無いよ",
		},
		{
			Probability:  1,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "なんかおかしい",
		},
	}).Drunk())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "うごいた")
	})

	r.Run()
}
```

```
$ curl -X GET -IL localhost:8080
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:08 GMT
Content-Length: 20

$ curl -X GET -IL localhost:8080
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:10 GMT
Content-Length: 23

$ curl -X GET -IL localhost:8080
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:11 GMT
Content-Length: 14

$ curl -X GET -IL localhost:8080
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:11 GMT
Content-Length: 14

$ curl -X GET -IL localhost:8080
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:12 GMT
Content-Length: 23

$ curl -X GET -IL localhost:8080
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 04 May 2019 04:16:14 GMT
Content-Length: 14
```