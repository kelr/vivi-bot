package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	BanchoLockIn  = "1447975441183936736"
	ViviSusLeft   = "1456174507059314861"
	ViviSusCenter = "1457213252839932067"
	ViviSusRight  = "1457228017590861834"
)

var (
	Token           = flag.String("t", "", "Bot authentication token")
	ViviSusStickers = []string{ViviSusLeft, ViviSusCenter, ViviSusRight}
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

	SendViviSticker(s, m)
	SendLockInSticker(s, m)

	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}
}

func SendViviSticker(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "<@1457571257766772957>") {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			StickerIDs: SelectViviSticker(),
		})
	}
}

func SelectViviSticker() []string {
	stickerIndex := rand.Intn(len(ViviSusStickers))
	return []string{ViviSusStickers[stickerIndex]}
}

func SendLockInSticker(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(strings.ToLower(m.Content), "lock in") {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			StickerIDs: []string{BanchoLockIn},
		})
	}
}
