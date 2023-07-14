package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func loadCFG() {
	fmt.Println("Loading configuration files")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/discord-referee/")
	viper.AddConfigPath("$HOME/.discord-referee")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", viper.GetString("mysql"))
	if err != nil {
		log.Panicf("Cannot connect to MySQL database: %v", err)
	}
	return db
}

func main() {
	loadCFG()
	db := connectDB()
	s := setup(db)
	cmds := registerCMD(s)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	db.Close()
	deleteCMD(cmds, s)
}
