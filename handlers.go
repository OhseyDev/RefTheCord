package main

import "github.com/bwmarrin/discordgo"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isTournamentChannel(m.ChannelID) {
		return
	}
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !isValidEntry(m) || isTournamentOpen(m.ChannelID) {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		return
	}
	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	s.MessageReactionAdd(m.ChannelID, m.ID, "❎")
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if !isTournamentChannel(r.ChannelID) {
		return
	}
	id := r.Emoji.APIName()
	if id != "✅" && id != "❎" {
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.Member.User.ID)
	}
}
