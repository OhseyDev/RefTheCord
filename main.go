package main

import (
	"log"
	_ "time"

	"github.com/OhseyDev/RefTheCord/lib"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	//"flag"
)

type TournamentType int

const (
	INVALID_TOURNAMENT TournamentType = 0
	BASIC_TOURNAMENT   TournamentType = 1
)

type Tournament struct {
	channel string
	variant TournamentType
}

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "stop",
		Description: "Shutdown Discord bot; development only",
	},
}

var cfg = lib.NewConfig()
var db = lib.PrepareDB(cfg)

func main() {
	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Error on Discord login: %v", err)
		return
	}
	running := true
	discord.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildScheduledEvents
	discord.AddHandler(messageCreate)
	discord.AddHandler(messageReactionAdd)
	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"stop": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User.ID == "890285201287151656" {
				running = false
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{Content: "Shutting down bot"},
				})
			}
		},
	}
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	if err = discord.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}
	defer discord.Close()
	defer db.Close()

	log.Print("Bot is now running. Press CTRL-C to exit.")
	for running {
	}
}

func isValidEntry(m *discordgo.MessageCreate) bool {
	valid := false
	for _, _ = range m.Attachments {
		valid = true
	}
	for _, _ = range m.Embeds {
		valid = true
	}
	return valid
}

func isTournamentChannel(channelID string) bool {
	return channelID == "1219658161921724426"
}

func isTournamentOpen(_channelID string) bool {
	return true
}

//"1219658161921724426" == basic tournament example channel
