package commands;

func generate() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand {
		{
			Name: "tourney",
			Description: "Manage Discord server tournaments",
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "new",
					Description: "Create a new tournament",
					Options: []*discordgo.ApplicationCommandOption {},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name: "modify",
					Description: "Modify an existing tournament",
					Options: []*discordgo.ApplicationCommandOption {},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name: "delete",
					Description: "Delete a tournament",
					Options: []*discordgo.ApplicationCommandOption {},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name: "list",
					Description: "List all tournaments",
					Options: []*discordgo.ApplicationCommandOption {},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		}
	};
}

