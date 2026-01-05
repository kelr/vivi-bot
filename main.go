package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	ViviSusStickerId = "1456174507059314861"
)

var (
	Token = flag.String("t", "", "Bot authentication token")
)

func main() {
	flag.Parse()
	if *Token == "" {
		log.Fatal("Missing discord auth token flag, provide it with -t")
	}

	dg, err := discordgo.New("Bot " + *Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Vivi bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Block until shutdown
	<-sc
	err = dg.Close()
	if err != nil {
		log.Println("Could not close session gracefully:", err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println(m.Content)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "<@1457443748601659554>") {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			StickerIDs: []string{ViviSusStickerId},
		})
	}

	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}

}
