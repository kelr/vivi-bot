package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	// Guild bound
	BanchoLockInSticker  = "1447975441183936736"
	ViviSusLeftSticker   = "1456174507059314861"
	ViviSusCenterSticker = "1457213252839932067"
	ViviSusRightSticker  = "1457228017590861834"

	// App bound
	PekoPogEmoji       = "1458287997815361650"
	PekoHappyEmoji     = "1458289554715840583"
	PekoFeelGoodEmoji  = "1458289480741163091"
	PekoHeheEmoji      = "1458289430048673967"
	PekoFeelGood2Emoji = "1458289895544983644"

	SuiWobbleEmoji = "1458287981277216838"
	SuiSwayEmoji   = "1458288617985146951"
	SuiDanceEmoji  = "1458288483813822686"
	SuiWavyEmoji   = "1458288381271216138"
	SuiBounceEmoji = "1458288349927309467"
)

var (
	Token = flag.String("t", "", "Bot authentication token")

	// Precompile regex
	lockInRegexCompiled  = regexp.MustCompile(`(?i)\b(lock(?:ed|ing|s)?[-\s]?in)\b`)
	omgSuiRegexCompiled  = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?sui+)\b`)
	omgPekoRegexCompiled = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?peko+)\b`)

	// List of stickers and emojis to randomly select from
	ViviSusStickers = []string{ViviSusLeftSticker, ViviSusCenterSticker, ViviSusRightSticker}
	OmgSuiEmojis    = []string{SuiWobbleEmoji, SuiSwayEmoji, SuiDanceEmoji, SuiWavyEmoji, SuiBounceEmoji}
	OmgPekoEmojis   = []string{PekoPogEmoji, PekoHappyEmoji, PekoFeelGoodEmoji, PekoHeheEmoji, PekoFeelGood2Emoji}
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

	// Message handlers
	ReactToMessageWithSticker(s, m)
	ReactToMessageWithEmoji(s, m)

	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}
}

// Listens to messages and sends a sticker when a match is detected
func ReactToMessageWithSticker(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Sends Vivi sticker when the bot is mentioned
	if strings.Contains(m.Content, "<@1457443748601659554>") {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{selectRandom(ViviSusStickers)}})
	}
	// Send Bancho Lock In sticker when someone mentions "lock in"
	if lockInRegexCompiled.MatchString(m.Content) {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{BanchoLockInSticker}})
	}
}

// Reacts to messages with emojis when a match is detected
func ReactToMessageWithEmoji(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Reacts with a random Sui emoji when "omg sui" is mentioned
	if omgSuiRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgSuiEmojis))
	}
	// Reacts with a random Peko emoji when "omg peko" is mentioned
	if omgPekoRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgPekoEmojis))
	}
}

// Randomly selects an element from a slice of strings
func selectRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
