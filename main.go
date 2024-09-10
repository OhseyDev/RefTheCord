package main
import (
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v10"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "time"
	//"flag"
)

type TournamentType int

const (
	INVALID_TOURNAMENT TournamentType = 0
	BASIC_TOURNAMENT TournamentType = 1
)

type Tournament struct {
	channel string
	variant TournamentType
}

type Config struct {
	Token		string	`env:"TOKEN"`
}

var commands = []*discordgo.ApplicationCommand {
	{
		Name: "stop",
		Description: "Shutdown Discord bot; development only",
	},
}

var cfg = loadConfig()
var db = prepareDB(&cfg)

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
	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"stop": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User.ID == "890285201287151656" {
				running = false
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData { Content: "Shutting down bot", },
				})
			}
		},
	}
	discord.AddHandler(func (s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	if err = discord.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}
	defer discord.Close()
	//defer db.Close()

	addPositionRoles(discord)

	log.Print("Bot is now running. Press CTRL-C to exit.")
	for running {
	}
}

func addPositionRoles(s *discordgo.Session) {
	vals := []string {"Goalkeeper", "Left-back", "Left centre-back", "Right centre-back", "Right-back", "Central Defensive Midfielder", "Centre Midfielder", "Central Attacking Midfielder", "Left-wing", "Right-wing", "Striker", }
	for _, item := range vals {
		role := discordgo.RoleParams{
			Name: "Position -> " + item,
			Color: 0x02193d,
			Hoist: false,
			Permissions: 0,
			Mentionable: false,
		}
		s.GuildRoleCreate("1116028177324711956", &role)
	}
}

func prepareDB(cfg *Config) *sql.DB {
	var db, err = sql.Open("sqlite3", "refthecord.db")
	if err != nil { log.Fatalf("Failed to connect to database: %v", err) }
	return db
}

func loadConfig() Config {
	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil { log.Fatalf("%v\n", err) }
	return cfg
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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isTournamentChannel(m.ChannelID) { return }
	if m.Author.ID == s.State.User.ID { return }
	if !isValidEntry(m) || isTournamentOpen(m.ChannelID) {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		return
	}
	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	s.MessageReactionAdd(m.ChannelID, m.ID, "❎")
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if !isTournamentChannel(r.ChannelID) {	return }
	id := r.Emoji.APIName()
	if id != "✅" && id != "❎" {
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.Member.User.ID)
	}
}

//"1219658161921724426" == basic tournament example channel

