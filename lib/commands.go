package lib

import "github.com/bwmarrin/discordgo"

func generate() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "tourney",
			Description: "Manage Discord server tournaments",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "new",
					Description: "Create a new tournament",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "name",
							Description: "Name of the tournament",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "type",
							Description: "Type of the tournament",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{
									Name:  "Single Elimination",
									Value: 1,
								},
								{
									Name:  "Double Elimination",
									Value: 2,
								},
								{
									Name:  "Round Robin",
									Value: 3,
								},
							},
							Required: true,
						},
						{
							Name:        "date",
							Description: "Date of the tournament",
							Options: []*discordgo.ApplicationCommandOption{
								{
									Name:        "month",
									Description: "Month",
									Type:        discordgo.ApplicationCommandOptionInteger,
									Choices:     []*discordgo.ApplicationCommandOptionChoice{},
									Required:    true,
								},
								{
									Name:        "year",
									Description: "Year",
									Type:        discordgo.ApplicationCommandOptionInteger,
									Required:    true,
								},
							},
							Required: true,
						}}}}}}
}
