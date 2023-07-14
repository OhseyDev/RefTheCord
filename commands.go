package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:                     "support",
			Description:              "Access guidance and support",
			DefaultMemberPermissions: create(discordgo.PermissionAll),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "info",
					Description: "Support information",
				},
				{
					Name:        "ticket",
					Description: "Open and manage support tickets",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "open",
							Description: "Open a new support ticket",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
						{
							Name:        "close",
							Description: "Close a support ticket",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommandGroup,
				},
			},
		},
		{
			Name:        "rules",
			Description: "Show rules",
		},
		{
			Name:        "submit",
			Description: "Manually submit a clip",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"support": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			content := ""
			switch options[0].Name {
			case "info":
				content = "Not Implemented"
			case "ticket":
				options = options[0].Options
				switch options[0].Name {
				case "open":
					content = "Not implemented"
				case "close":
					content = "Not implemented"
				case "update":
					content = "Not implemented"
				case "view":
					content = "Not implemented"
				default:
					content = "Oops, something went wrong.\n" +
						"Hol' up, you aren't supposed to see this message."
				}
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
		},
		"rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "TODO: Setup rules",
				},
			})
		},
		"submit": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "TODO: Implement manual submit",
				},
			})
		},
	}
)

func registerCMD(s *discordgo.Session) []*discordgo.ApplicationCommand {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, viper.GetString("guild"), v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	return registeredCommands
}

func deleteCMD(cmds []*discordgo.ApplicationCommand, s *discordgo.Session) {
	for _, v := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, viper.GetString("guild"), v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func create(num int64) *int64 {
	return &num
}
