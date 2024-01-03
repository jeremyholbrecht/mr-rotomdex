package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// insert your token
func main() {
	rotomdex, err := discordgo.New("Bot ")

	if err != nil {
		log.Fatal(err)
	}

	rotomdex.AddHandler(MessageCreate)

	rotomdex.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = rotomdex.Open()

	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("RotomDex is online. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer rotomdex.Close()

}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "bulbasaur" {
		bulbasaur := discordgo.MessageEmbed{
			Title:     "Bulbasaur",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0x00ff00, //green
			Image: &discordgo.MessageEmbedImage{
				URL: "https://archives.bulbagarden.net/media/upload/9/97/001Bulbasaur_RB.png",
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "DexNr",
					Value:  "0001",
					Inline: true,
				},

				{
					Name:   "Type",
					Value:  "Grass/Poison",
					Inline: true,
				},

				{
					Name:   "Species",
					Value:  "Seed Pokémon",
					Inline: true,
				},

				{
					Name:   "Abilities",
					Value:  "Overgrow/Chlorophyll(Hidden Ability)",
					Inline: false,
				},

				{
					Name:   "Height",
					Value:  "0.7 m (2′04″)",
					Inline: true,
				},

				{
					Name:   "Weight",
					Value:  "6.9 kg (15.2 lbs)",
					Inline: true,
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &bulbasaur)
	}
}
