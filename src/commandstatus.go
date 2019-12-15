package main

import (
	"strings"

	"github.com/diamondburned/discordgo"
)

func setStatus(input []string) {
	if d.State.Settings == nil {
		Message("Settings are uninitialized")
		return
	}

	s := d.State.Settings.Status

	if len(input) < 2 {
		switch s {
		case discordgo.StatusOnline:
			Message("Status: Online")
		case discordgo.StatusIdle:
			Message("Status: Idle")
		case discordgo.StatusDoNotDisturb:
			Message("Status: Do not disturb")
		case discordgo.StatusInvisible:
			Message("Status: Invisible")
		default:
			Message(string(s))
		}

		return
	}

	switch strings.Join(input[1:], " ") {
	case string(discordgo.StatusOnline), "Online":
		s = discordgo.StatusOnline

	case string(discordgo.StatusIdle), "Idle",
		"Away", "away":
		s = discordgo.StatusIdle

	case string(discordgo.StatusDoNotDisturb),
		"do not disturb", "Do not disturb", "Do Not Disturb",
		"Busy", "busy":
		s = discordgo.StatusDoNotDisturb

	case string(discordgo.StatusInvisible), "invis", "Invisible":
		s = discordgo.StatusInvisible

	default:
		Message("Unknown status to set, check description")
		return
	}

	if _, err := d.UserUpdateStatus(s); err != nil {
		Warn(err.Error())
		return
	}

	Message("Set status to " + string(s))
}

func setListen(text []string) {
	if len(text) < 2 {
		Message("Missing string!")
		return
	}

	s := strings.Join(text[1:], " ")

	if err := d.UpdateListeningStatus(s); err != nil {
		Message(err.Error())
	} else {
		Message("Set listening status to " + s)
	}
}

func setGame(text []string) {
	var s string
	if len(text) > 1 {
		s = strings.Join(text[1:], " ")
	}

	var (
		msg      string
		gametype = discordgo.GameTypeGame
	)

	switch {
	case strings.HasPrefix(strings.ToLower(s), "listening to "):
		s = s[13:]
		gametype = discordgo.GameTypeListening
		msg = "Set listening to "
	case strings.HasPrefix(strings.ToLower(s), "watching "):
		s = s[9:]
		gametype = discordgo.GameTypeWatching
		msg = "Set watching "
	default:
		msg = "Set game to "
	}

	usd := discordgo.UpdateStatusData{
		Status: string(d.State.Settings.Status),
		Game: &discordgo.Game{
			Name: s,
			Type: gametype,
		},
	}

	if err := d.UpdateStatusComplex(usd); err != nil {
		Message(err.Error())
		return
	}

	if s != "" {
		Message(msg + s + ".")
	} else {
		Message("Reset presence successfully.")
	}
}
