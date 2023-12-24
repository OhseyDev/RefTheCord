package main
import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"flag"
	"os"
	"log"
)

var (
	Token string
	DBConnect string
)

func main() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&DBConnect, "db", "", "Database Connect String")
	flag.Parse()
	if Token == nil || len(Token) <= 0 {
		Token = os.Getenv("BOT_TOKEN")
		if Token == nil {
			log.Fatalf("No Discord Bot Token provided")
			return
		}
	}
	discord, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatalf("Error on Discord login: " + err)
		return
	}
	discord.AddHandler(messageCreate)
	log.Print("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.UserID {
		return
	}
	chanid := m.ChannelID
}

