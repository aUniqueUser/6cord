package main

import (
	"strings"

	"github.com/diamondburned/discordgo"
)

func messageDelete(s *discordgo.Session, rm *discordgo.MessageDelete) {
	if d == nil || Channel == nil {
		return
	}

	if rm.ChannelID != Channel.ID {
		return
	}

	for i := len(messageStore) - 1; i >= 0; i-- {
		if ID := getIDfromindex(i); ID != 0 {
			if rm.ID != ID {
				continue
			}

			prev := 0

			if i > 0 {
				if strings.HasPrefix(messageStore[i-1], "\n\n[#") {
					prev = 1
					setLastAuthor(0)
				}
			}

			messageStore = append(
				messageStore[:i-prev],
				messageStore[i+1:]...,
			)

			app.QueueUpdate(func() {
				messagesView.SetText(strings.Join(messageStore, ""))
			})

			scrollChat()

			return
		}
	}
}

func messageDeleteBulk(s *discordgo.Session, rmb *discordgo.MessageDeleteBulk) {
	for _, m := range rmb.Messages {
		messageDelete(s, &discordgo.MessageDelete{
			Message: &discordgo.Message{
				ChannelID: rmb.ChannelID,
				ID:        m,
			},
		})
	}
}
