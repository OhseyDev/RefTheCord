package main
import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"github.com/caarlos0/env/v10"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"log"
	"syscall"
	"time"
)

type config struct {
	Token		string	`env:"TOKEN"`
	ConnectStr	string	`env:"DBCONN"`
	MaxConnDB	int	`env:"DBCONNLIMIT"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil { fmt.Printf("%+v\n", err) }
	db, err := sql.Open("mysql", cfg.ConnectStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(cfg.MaxConnDB)
	db.SetMaxIdleConns(cfg.MaxConnDB)
	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Error on Discord login: %v", err)
		return
	}
	discord.Identify.Intents = 17995913063488
	discord.AddHandler(messageCreate)
	if err = discord.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}
	defer discord.Close()
	log.Print("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

