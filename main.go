package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mrz1836/go-sanitize"
)

const (
	ProdViviBotId       = "<@1457443748601659554>"
	BakeyLocalTestBotId = "<@1457571257766772957>"
)

var (
	Token = flag.String("t", "", "Bot authentication token")

	// Precompile regex
	LockInRegexCompiled        = regexp.MustCompile(`(?i)\b(lock(?:ed|ing|s)?[-\s]?in)\b`)
	OmgSuiRegexCompiled        = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?sui+)\b`)
	OmgPekoRegexCompiled       = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?pe(k|g)o+)\b`)
	OmgTowaRegexCompiled       = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?towa+(sama)?)\b`)
	OmgLuiRegexCompiled        = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(lui|looi)+)\b`)
	OmgCCRegexCompiled         = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?cc)\b`)
	OmgGGRegexCompiled         = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?gg)\b`)
	OmgAutoFisterRegexCompiled = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(autofister|ccgg))\b`)
	OmgLamyRegexCompiled       = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(l|w)amy+)\b`)
	OmgBaeRegexCompiled        = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(hakos|bae|baelz|rat))\b`)
	OmgViviRegexCompiled       = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?vivi+)\b`)
	OmgRaoraRegexCompiled      = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?raora)\b`)
	OmgFuwaMocoRegexCompiled   = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(fuwamoco|fwmc))\b`)
	OmgMocoFuwaRegexCompiled   = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(mocofuwa))\b`)
	OmgFuwawaRegexCompiled     = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(fuwawa))\b`)
	OmgMococoRegexCompiled     = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(mococo))\b`)

	// List of stickers and emojis to randomly select from
	ViviSusStickers   = []string{ViviSusLeftSticker, ViviSusCenterSticker, ViviSusRightSticker}
	OmgSuiEmojis      = []string{SuiWobble, SuiSway, SuiDance, SuiWavy, SuiBounce, SuiAha, SuiFukkireta, SuiPuppet, SuiLaugh, SuiLaugh2, SuiVibe, SuiOrbitalSpin, SuiPuppetKocchi}
	OmgPekoEmojis     = []string{PekoPog, PekoHappy, PekoFeelGood, PekoHehe, PekoFeelGood2}
	OmgTowaEmojis     = []string{TowaNuma, TowaDance, TowaLaugh, TowaCool, TowaPose, TowaBapFast, TowaSpin, TowaHug, TowaHeadbang, TowaBlush}
	OmgCCEmojis       = []string{CCCool, CCSmug, CCParty, CCCheer, CCDoro, CCWave, CCShake, CCWide}
	OmgGGEmojis       = []string{GigiNod, GigiBleh, GigiPeek, GigiHug, GigiHD, GigiFukkireta, GigiJam, GigiDoro, GigiCool, GigiBark}
	OmgFuwaMocoEmojis = [][]string{{FuwawaBite, MococoBite}, {FuwawaBau, MococoBau}, {FuwawaEhehe, MococoHOEH3}, {FuwawaEhehe2, MococoHOEH2}, {FuwawaDoro, MococoDoro}}
	OmgFuwawaEmojis   = []string{FuwawaYay, FuwawaBau, FuwawaEhehe, FuwawaEhehe2, FuwawaPeek, FuwawaMou}
	OmgMococoEmojis   = []string{MococoCool, MococoBlink, MococoNod, MococoNod2, MococoBau, MococoHappy, MococoJamFast}
	OmgBaeEmojis      = []string{BaePog, BaePogey, BaeAhoy, BaeDoro, BaeSewer, BaeCool, BaeBlush, BaeBreakdance, BaeBreakdanceFast, BaeBreakdanceFastest, BaeLove, BaeNod, BaePogU, BaeSmug2, BaeSmug1, BaeWave, BaeWink}

	OmgLuiEmojis   = []string{madaikanai}
	OmgLamyEmojis  = []string{madaikanai}
	OmgViviEmojis  = []string{madaikanai}
	OmgRaoraEmojis = []string{madaikanai}

	LocalTestEmojis = []string{LocalTestEmoji, LocalTestEmoji2, LocalTestEmoji3}
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

	// Sanitize whitespace in message content
	m.Content = sanitize.AlphaNumeric(m.Content, true)

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
	if strings.Contains(m.Content, ProdViviBotId) {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{selectRandom(ViviSusStickers)}})
	}
	// Send Bancho Lock In sticker when someone mentions "lock in"
	if LockInRegexCompiled.MatchString(m.Content) {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{BanchoLockInSticker}})
	}
}

// Reacts to messages with emojis when a match is detected
func reactToMessageWithEmoji(s *discordgo.Session, m *discordgo.MessageCreate) {
	if OmgSuiRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgSuiEmojis))
	}
	if OmgPekoRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgPekoEmojis))
	}
	if OmgTowaRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgTowaEmojis))
	}
	if OmgCCRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgCCEmojis))
	}
	if OmgGGRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgGGEmojis))
	}
	if OmgAutoFisterRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FluffyCC)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FuzzyGG)
	}
	if OmgMocoFuwaRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+MococoDoro)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+FuwawaDoro)
	}
	if OmgFuwaMocoRegexCompiled.MatchString(m.Content) {
		fwmcPair := OmgFuwaMocoEmojis[rand.Intn(len(OmgFuwaMocoEmojis))]
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[0])
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[1])
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau1)
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau2)
	}
	if OmgFuwawaRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgFuwawaEmojis))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> **WHAT ABOUT MOCOCOEH!?**", m.Author.ID))
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{MococoHOEHSticker}})
	}
	if OmgMococoRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgMococoEmojis))
	}
	if OmgBaeRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgBaeEmojis))
	}
	if OmgLuiRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgLuiEmojis))
	}
	if OmgLamyRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgLamyEmojis))
	}
	if OmgViviRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgViviEmojis))
	}
	// TODO: see DM's
	if OmgRaoraRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgRaoraEmojis))
	}
	// TODO: Convert into a table and loop thru for the repeated code
}

// Randomly selects an element from a slice of strings
func selectRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
