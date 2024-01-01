package main

import "github.com/bwmarrin/discordgo"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { return }
	// chanid := m.ChannelID
}

