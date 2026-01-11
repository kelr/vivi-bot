package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mrz1836/go-sanitize"
)

var (
	Token = flag.String("t", "", "Bot authentication token")
	BotId = flag.String("u", "", "Bot user ID")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
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

	// Sanitize whitespace in message content for only alphanumeric, <, >, @ and standard spaces
	m.Content = sanitize.Custom(m.Content, `[^a-zA-Z0-9<@>\s/*-+]`)

	// Message handlers
	reactToMessageWithSticker(s, m)
	reactToMessageWithEmoji(s, m)

	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}
}

// Listens to messages and sends a sticker when a match is detected
func reactToMessageWithSticker(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Sends Vivi sticker when the bot is mentioned
	if strings.Contains(m.Content, *BotId) {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{selectRandom(ViviSusStickers)}})
	}
	// Send Bancho Lock In sticker when someone mentions "lock in"
	if LockInRegexCompiled.MatchString(m.Content) {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{BanchoLockInSticker}})
	}
}

// Reacts to messages with emojis when a match is detected
func reactToMessageWithEmoji(s *discordgo.Session, m *discordgo.MessageCreate) {

	// TODO: Use a parallelization package if this ever grows large enough to matter
	// Dynamically loop thru OmgMemNameMappings and look for matches
	for _, holoMemKVP := range OmgMemNameMappings {
		if holoMemKVP.RegexExpr.MatchString(m.Content) {
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(holoMemKVP.EmojiList))
		}
	}

	// Custom case handling
	if OmgAutoFisterRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FluffyCC)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FuzzyGG)
	}
	if OmgFuwaMocoRegexCompiled.MatchString(m.Content) {
		fwmcPair := OmgFuwaMocoEmojis[rand.Intn(len(OmgFuwaMocoEmojis))]
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[0])
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[1])
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau1)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau2)
	}
	if OmgMocoFuwaRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+MococoDoro)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FuwawaDoro)
	}
	if OmgFuwawaRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgFuwawaEmojis))
		if !OmgMococoRegexCompiled.MatchString(m.Content) && !OmgFuwaMocoRegexCompiled.MatchString(m.Content) && !OmgMocoFuwaRegexCompiled.MatchString(m.Content) {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> **WHAT ABOUT MOCOCOEH!?**", m.Author.ID))
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{MococoHOEHSticker}})
		}
	}
	if uuuuuCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FaunaUUUUU)
	}
}

// Randomly selects an element from a slice of strings
func selectRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
