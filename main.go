package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

var (
	GuildID            = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken           = flag.String("token", "", "Bot access token")
	MySQL              = flag.String("mysql", "", "MySQL Database URL")
	SubmissionsChannel = flag.String("submissions", "", "Submissions channel")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	db, err := sql.Open("mysql", *MySQL)
	if err != nil {
		log.Panicf("Cannot connect to MySQL database: %v", err)
	}

	db.Exec("CREATE TABLE Submissions (User varchar(18) NOT NULL, DiscordMsg varchar(), Votes TINYINT(255), PRIMARY KEY (User))")
	db.Exec("CREATE TABLE UserVotes (User varchar(18) NOT NULL, EventYr YEAR NOT NULL, EventMonth BIT(4), Submission longint, PRIMARY KEY (User))")

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
		if c.ChannelID != *SubmissionsChannel {
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
		if d.ChannelID != *SubmissionsChannel {
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

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
	db.Close()
	log.Println("Gracefully shutting down.")
}

func create(num int64) *int64 {
	return &num
}
