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
	rotomdex, err := discordgo.New("Bot MTE5MjA0MzMyNTU0OTcxNTU1Ng.GFprpk.ulX6DcKQUFtunlArGVzcDQEMR-LFweqxt_5Nuw")

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
				URL: "https://images.wikidexcdn.net/mwuploads/wikidex/4/43/latest/20190406170624/Bulbasaur.png",
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

	if m.Content == "ivysaur" {
		ivysaur := discordgo.MessageEmbed{
			Title:     "Ivysaur",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0x00ff00, //green
			Image: &discordgo.MessageEmbedImage{
				URL: "https://images.wikidexcdn.net/mwuploads/wikidex/8/86/latest/20190406151903/Ivysaur.png",
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "DexNr",
					Value:  "0002",
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
					Value:  "1.0 m (3′03″)",
					Inline: true,
				},

				{
					Name:   "Weight",
					Value:  "13.0 kg (28.7 lbs)",
					Inline: true,
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &ivysaur)
	}

	if m.Content == "nidorino" {
		nidorino := discordgo.MessageEmbed{
			Title:     "Nidorino",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0xba55d3, //purple
			Image: &discordgo.MessageEmbedImage{
				URL: "https://images.wikidexcdn.net/mwuploads/wikidex/7/7a/latest/20080909113745/Nidorino.png",
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "DexNr",
					Value:  "0033",
					Inline: true,
				},

				{
					Name:   "Type",
					Value:  "Poison",
					Inline: true,
				},

				{
					Name:   "Species",
					Value:  "Poison Pin Pokémon",
					Inline: true,
				},

				{
					Name:   "Abilities",
					Value:  "Overgrow/Rivalry/Hustle(Hidden Ablity)",
					Inline: false,
				},

				{
					Name:   "Height",
					Value:  "0.9 m (2′11″)",
					Inline: true,
				},

				{
					Name:   "Weight",
					Value:  "19.5 kg (43.0 lbs)",
					Inline: true,
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &nidorino)
	}

}
