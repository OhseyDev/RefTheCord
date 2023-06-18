package main

import (
	"github.com/bwmarrin/discordgo"
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
