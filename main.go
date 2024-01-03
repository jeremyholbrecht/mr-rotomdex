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
	rotomdex, err := discordgo.New("Bot")

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
		var bulbasaurImg = discordgo.MessageEmbedThumbnail{
			URL: "http://pldh.net/media/pokemon/ken_sugimori/original/001.png",
		}

		bulbasaur := discordgo.MessageEmbed{
			Title:     "Bulbasaur",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0x00ff00,
			Thumbnail: &bulbasaurImg,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &bulbasaur)
	}
}
