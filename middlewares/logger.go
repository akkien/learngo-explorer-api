package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"log"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.MustGet("requestId").(string)
		logger := c.MustGet("logger").(*log.Logger)

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		logger.Println(requestId+"-header:", c.Request.Header)

		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, body); err != nil {
			fmt.Println(err)
		}
		logger.Println(requestId+"-body:", buffer)

		c.Next()
	}
}
