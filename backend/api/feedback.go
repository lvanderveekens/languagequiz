package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type FeedbackHandler struct {
}

func NewFeedbackHandler() *FeedbackHandler {
	return &FeedbackHandler{}
}

func (h *FeedbackHandler) ReceiveFeedback(c *gin.Context) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "JsAWTVz58UF6SaFf@gmail.com")
	m.SetHeader("To", "lvanderveekens@proton.me")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "lucianovdveekens@gmail.com", "dlczbnnhepjlrrku")

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
