package main

import (
	"strconv"
	"strings"
	"time"
)

func messageGetTopID() int64 {
	if len(messageStore) < 1 {
		return 0
	}

	for i := 0; i < len(messageStore); i++ {
		if index := getIDfromindex(i); index != 0 {
			return index
		}
	}

	return 0
}

func getIDfromindex(i int) int64 {
	if len(messageStore) <= i || i < 0 {
		return 0
	}

	// magic number lo
	if len(messageStore[i]) < 23 {
		return 0
	}

	switch {
	case messageStore[i][1] != '[':
		return 0
	case messageStore[i][2] != '"':
		return 0
	}

	var (
		idRune strings.Builder
		msg    = messageStore[i]
	)

	for i := 3; i < len(msg); i++ {
		if msg[i] == '"' {
			break
		}

		idRune.WriteByte(msg[i])
	}

	r, _ := strconv.ParseInt(idRune.String(), 10, 64)
	return r
}

var loading bool

func loadMore() {
	if d == nil || Channel == nil {
		return
	}

	if loading {
		return
	}

	beforeID := messageGetTopID()
	if beforeID == 0 {
		return
	}

	loading = true
	input.SetPlaceholder("Loading more...")

	defer func() {
		input.SetPlaceholder(cfg.Prop.DefaultStatus)
		loading = false
	}()

	c, err := d.State.Channel(Channel.ID)
	if err != nil {
		Warn(err.Error())
		return
	}

	msgs, err := d.ChannelMessages(Channel.ID, 35, beforeID, 0, 0)
	if err != nil {
		return
	}

	if len(msgs) == 0 {
		// Drop out early if no messages
		return
	}

	var reversed = make([]string, 0, len(msgs)*2)

	for i := len(msgs) - 1; i >= 0; i-- {
		m := msgs[i]

		if rstore.Check(m.Author, RelationshipBlocked) && cfg.Prop.HideBlocked {
			continue
		}

		sentTime, err := m.Timestamp.Parse()
		if err != nil {
			sentTime = time.Now()
		}

		if i < len(msgs)-1 && (msgs[i+1].Author.ID != m.Author.ID || messageisOld(m, msgs[i+1])) {
			username, color := us.DiscordThis(m)

			reversed = append(reversed,
				authorTmpl.ExecuteString(map[string]interface{}{
					"color": fmtHex(color),
					"name":  username,
					"time":  sentTime.Format(time.Stamp),
				}),
			)
		}

		reversed = append(reversed,
			messageTmpl.ExecuteString(map[string]interface{}{
				"ID":      strconv.FormatInt(m.ID, 10),
				"content": fmtMessage(m),
			}),
		)
	}

	//wg.Wait()

	messageStore = append(reversed, messageStore...)
	messagesView.SetText(strings.Join(messageStore, ""))

	input.SetPlaceholder("Done.")

	for i, j := 0, len(msgs)-1; i < j; i, j = i+1, j-1 {
		msgs[i], msgs[j] = msgs[j], msgs[i]
	}

	d.State.Lock()
	c.Messages = append(msgs, c.Messages...)
	d.State.Unlock()

	messagesView.Highlight(strconv.FormatInt(beforeID, 10))
	messagesView.ScrollToHighlight()

	time.Sleep(time.Second * 5)
}
