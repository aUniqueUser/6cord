package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/discordgo"
	"github.com/diamondburned/tview/v2"
	"gitlab.com/diamondburned/6cord/antitele"
	"gitlab.com/diamondburned/6cord/shortener"
)

const zeroes = "000000"

func fmtHex(hex int) string {
	h := zeroes + strconv.FormatInt(int64(hex), 16)
	return h[len(h)-len(zeroes):]
}

func fmtMessage(m *discordgo.Message) string {
	ct, emojiMap := parseMessageContent(m)
	ct = strings.Map(func(r rune) rune {
		for _, z := range antitele.ZeroWidthRunes {
			if z == r {
				return -1
			}
		}

		return r
	}, ct)

	if m.EditedTimestamp != "" {
		ct += " " + readChannelColorPrefix + "(edited)[-::-]"

		// Prevent cases where the message is empty
		// " (edited)"
		ct = strings.TrimPrefix(ct, " ")
	}

	var (
		c strings.Builder
		l = strings.Split(ct, "\n")

		attachments         = m.Attachments
		reactions, reactMap = formatReactions(m.Reactions)
	)

	if ct != "" {
		for i := 0; i < len(l); i++ {
			if !cfg.Prop.CompactMode || i != 0 {
				c.WriteString(chatPadding)
			}

			c.WriteString(l[i])

			if i != len(l)-1 {
				c.WriteByte('\n')
			}
		}
	}

	for k, v := range reactMap {
		emojiMap[k] = v
	}

	for _, arr := range emojiMap {
		attachments = append(
			attachments,
			&discordgo.MessageAttachment{
				Filename: arr[0],
				URL:      arr[1],
			},
		)
	}

	for _, e := range m.Embeds {
		var embed = make([]string, 0, 5)

		if e.URL != "" {
			attachments = append(
				attachments,
				&discordgo.MessageAttachment{
					Filename: "EmbedURL",
					URL:      e.URL,
				},
			)
		}

		if e.Author != nil {
			embed = append(
				embed,
				"[::du]"+e.Author.Name+"[::-]",
			)

			if e.Author.IconURL != "" {
				attachments = append(
					m.Attachments,
					&discordgo.MessageAttachment{
						Filename: "AuthorIcon",
						URL:      e.Author.IconURL,
					},
				)
			}

			if e.Author.URL != "" {
				attachments = append(
					m.Attachments,
					&discordgo.MessageAttachment{
						Filename: "AuthorURL",
						URL:      e.Author.URL,
					},
				)
			}
		}

		if e.Title != "" {
			embed = append(
				embed,
				splitEmbedLine(e.Title, "[::b]", "[#0096cf]")...,
			)
		}

		if e.Description != "" {
			var desc, emojis = parseEmojis(e.Description)

			embed = append(embed, splitEmbedLine(desc)...)

			for _, arr := range emojis {
				attachments = append(
					m.Attachments,
					&discordgo.MessageAttachment{
						Filename: arr[0],
						URL:      arr[1],
					},
				)
			}
		}

		if len(e.Fields) > 0 {
			embed = append(embed, "")

			for _, f := range e.Fields {
				embed = append(embed,
					splitEmbedLine(f.Name, " [::b]")...)
				embed = append(embed,
					splitEmbedLine(f.Value, " [::d]")...)
				embed = append(embed, "")
			}
		}

		var footer []string
		if e.Footer != nil {
			footer = append(
				footer,
				"[::d]"+tview.Escape(e.Footer.Text)+"[::-]",
			)

			if e.Footer.IconURL != "" {
				attachments = append(
					m.Attachments,
					&discordgo.MessageAttachment{
						Filename: "FooterIcon",
						URL:      e.Footer.IconURL,
					},
				)
			}
		}

		if e.Timestamp != "" {
			footer = append(
				footer,
				"[::d]"+e.Timestamp+"[::-]",
			)
		}

		if len(footer) > 0 {
			embed = append(
				embed,
				strings.Join(footer, " - "),
			)
		}

		//if e.Thumbnail != nil {
		//attachments = append(
		//m.Attachments,
		//&discordgo.MessageAttachment{
		//Filename: "Thumbnail",
		//URL:      e.Thumbnail.URL,
		//},
		//)
		//}

		if e.Image != nil {
			attachments = append(
				m.Attachments,
				&discordgo.MessageAttachment{
					Filename: "Image",
					URL:      e.Image.URL,
				},
			)
		}

		if e.Video != nil {
			attachments = append(
				m.Attachments,
				&discordgo.MessageAttachment{
					Filename: "Video",
					URL:      e.Video.URL,
				},
			)
		}

		var embedPadding = chatPadding
		if len(embedPadding) > 2 {
			embedPadding = chatPadding[:len(chatPadding)-2]
		}

		c.WriteByte('\n')

		for i, l := range embed {
			c.WriteString(embedPadding + fmt.Sprintf("[#%06X]", e.Color) + "┃[-::] " + l)

			if i != len(embed)-1 {
				c.WriteByte('\n')
			}
		}
	}

	if len(m.Reactions) > 0 { // Reactions
		c.WriteString("\n" + chatPadding + chatPadding + reactions)
	}

	if len(attachments) > 0 {
		for _, a := range attachments {
			c.WriteString(fmt.Sprintf(
				"\n%s[::d][%s[]: %s[::-]",
				chatPadding,
				tview.Escape(a.Filename),
				shortener.ShortenURL(a.URL),
			))
		}
	}

	if cfg.Prop.CompactMode {
		// If the message begins with a code block,
		// we don't want the first line of the code
		// block to warp in the first line.
		if len(m.Content) > 3 && m.Content[:3] == "```" {
			return "\n" + chatPadding + c.String()
		}
	}

	return c.String()
}
