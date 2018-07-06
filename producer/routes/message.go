package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puper/qman/producer/components/producer"

	"github.com/puper/qman/producer/app"
)

type Message struct {
}

func (this *Message) Configure(r *gin.Engine) {
	r.POST("/message/put", this.Put)
}

func (this *Message) Put(c *gin.Context) {
	msg := &producer.Message{
		Topic:      c.PostForm("topic"),
		Tag:        c.PostForm("tag"),
		Key:        c.PostForm("key"),
		BusinessID: c.PostForm("business_id"),
		Value:      c.PostForm("value"),
	}
	err := app.GetProducer().Put(msg)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    0,
				"message": err.Error(),
			},
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"result": nil,
		})
	}
}
