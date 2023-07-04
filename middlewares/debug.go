package middlewares

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type debugBody struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
	Body     string `json:"body"`
}

func Debug() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

		body := &debugBody{Action: c.Request.Method, Resource: c.Request.RequestURI, Body: string(data)}
		fmt.Println("############## DEBUG Message ##############")
		fmt.Println(body)
		fmt.Println("############## DEBUG Message ##############")

		c.Next()
	}
}
