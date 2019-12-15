package main

import (
	"github.com/diamondburned/discordgo"
)

func guildMemberAdd(s *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	guildMemberDoSomething(gma.Member)
}

func guildMemberRemove(s *discordgo.Session, gmr *discordgo.GuildMemberRemove) {
	if gmr.User == nil {
		return
	}

	us.RemoveUser(gmr.GuildID, gmr.User.ID)
}

func guildMemberUpdate(s *discordgo.Session, gma *discordgo.GuildMemberUpdate) {
	guildMemberDoSomething(gma.Member)
}

func guildMemberDoSomething(gm *discordgo.Member) {
	if gm.User == nil {
		return
	}

	us.UpdateUser(
		gm.GuildID,
		gm.User.ID,
		gm.User.Username,
		gm.Nick,
		gm.User.Discriminator,
		getUserColor(gm.GuildID, gm.Roles),
	)
}
