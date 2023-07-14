package main

import (
	"database/sql"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func setup(db *sql.DB) *discordgo.Session {
	s, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	s.AddHandler(func(s *discordgo.Session, react *discordgo.MessageReactionAdd) {
		if react.Emoji.Name != "white_check_mark" {
			return
		}
		msg, err := s.ChannelMessage(react.ChannelID, react.MessageID)
		if err != nil {
			log.Printf("Failed to obtain Discord message")
			return
		}
		if msg.Author.ID == react.UserID {
			return
		}
	})
	s.AddHandler(func(s *discordgo.Session, react *discordgo.MessageReactionRemove) {
		if react.Emoji.Name != "white_check_mark" {
			return
		}
		msg, err := s.ChannelMessage(react.ChannelID, react.MessageID)
		if err != nil {
			log.Printf("Failed to obtain Discord message")
			return
		}
		if msg.Author.ID == react.UserID {
			return
		}
	})
	s.AddHandler(func(s *discordgo.Session, c *discordgo.MessageCreate) {
		if c.ChannelID != viper.GetString("submissions-channel") {
			return
		}
		_, err := db.Query("SELECT COUNT(1) DiscordMsg FROM Submissions WHERE User='%s'", c.Author.ID)
		if err != nil {
			log.Printf("Failed to execute query: %r", err)
			return
		}
		s.ChannelMessageDelete(c.ChannelID, c.Message.ID)
	})
	s.AddHandler(func(s *discordgo.Session, d *discordgo.MessageDelete) {
		if d.ChannelID != viper.GetString("submissions-channel") {
			return
		}
		id, err := db.Query("SELECT COUNT(1) DiscordMsg FROM Submissions WHERE User='%s'", d.Author.ID)
		if err != nil {
			log.Printf("Failed to execute query: %r", err)
			return
		}
		loggedId := ""
		id.Scan(&loggedId)
		if loggedId == d.Message.ID {
			db.Exec("DELETE FROM Submissions WHERE User='%s'", d.Author.ID)
		}
	})
	err = s.Open()
	if err != nil {
		db.Close()
		log.Fatalf("Cannot open the session: %v", err)
	}
	return s
}
