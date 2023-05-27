package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	DiscordBotToken          string
	DiscordFeedbackChannelID string
}

func NewFeedbackHandler(discordBotToken, discordFeedbackChannelID string) *FeedbackHandler {
	return &FeedbackHandler{
		DiscordBotToken:          discordBotToken,
		DiscordFeedbackChannelID: discordFeedbackChannelID,
	}
}

type submitFeedbackRequest struct {
	Text string `json:"text"`
}

func (r *submitFeedbackRequest) validate() error {
	if r.Text == "" {
		return errors.New("field 'text' is missing")
	}
	return nil
}

func (h *FeedbackHandler) SubmitFeedback(c *gin.Context) error {
	var req submitFeedbackRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	if err := req.validate(); err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	discord, err := discordgo.New("Bot " + h.DiscordBotToken)
	if err != nil {
		return err
	}

	_, err = discord.ChannelMessageSend(h.DiscordFeedbackChannelID, "Received feedback: "+req.Text)
	if err != nil {
		return err
	}

	return nil
}
