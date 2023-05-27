package api

import (
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

func (h *FeedbackHandler) ReceiveFeedback(c *gin.Context) error {
	discord, err := discordgo.New("Bot " + h.DiscordBotToken)
	if err != nil {
		return err
	}

	_, err = discord.ChannelMessageSend(h.DiscordFeedbackChannelID, "hallo")
	if err != nil {
		return err
	}

	return nil
}
