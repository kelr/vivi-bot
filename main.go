package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sort"
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
	reactToMessageWithFile(s, m)

	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}
}

// Listens to messages and sends a sticker when a match is detected
func reactToMessageWithSticker(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Sends Vivi sticker when the bot is mentioned
	if strings.Contains(m.Content, *BotId) {
		if rand.Intn(10) == 0 {
			sendEmbedFile(s, m, fmt.Sprintf("<@%s>", m.Author.ID), `images\pekora_gonnahityou.mp4`)
		} else {
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{selectRandom(ViviSusStickers)}})
		}
	}
}

// Listens to messages and sends a file when a match is detected
func reactToMessageWithFile(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, FileEmbedKVP := range FileEmbedMappings {
		regexMatch := FileEmbedKVP.RegexExpr.MatchString(m.Content)
		if regexMatch {
			sendEmbedFile(s, m, fmt.Sprintf("<@%s>", m.Author.ID), selectRandom(FileEmbedKVP.EmojiList[0]))
		}
	}
}

// Reacts to messages with emojis when a match is detected
func reactToMessageWithEmoji(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Sui first no matter what because we are not a cult
	if OmgSuiRegexCompiled.MatchString(m.Content) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(OmgSuiEmojis))
		if rand.Intn(5) == 0 {
			s.ChannelMessageSend(m.ChannelID, selectRandom(SuiGifs))
		}
	}

	// Store matched "omg mem" from the message with their index
	matches := []RegexMatch{}
	matchedNames := map[string]bool{}
	matchedStickers := []RegexMatch{}

	// Dynamically loop thru OmgMemNameMappings and look for matches within the text and store an indexed list
	if len(m.StickerItems) > 0 {
		for name, StickerIdKVP := range StickerIdMappings {
			regexMatch := StickerIdKVP.RegexExpr.MatchString(m.StickerItems[0].ID)
			if regexMatch {
				matchedStickers = append(matchedStickers, RegexMatch{name: name, KVP: StickerIdKVP})
			}
		}
	} else {
		for name, holoMemKVP := range OmgMemNameMappings {
			regexMatch := holoMemKVP.RegexExpr.FindStringIndex(m.Content)
			if regexMatch != nil {
				matches = append(matches, RegexMatch{name: name, idx: regexMatch[0], KVP: holoMemKVP})
				matchedNames[name] = true
			}
		}
	}

	// Sort the matches by index to order the emoji output
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].idx < matches[j].idx
	})

	// React to messages according to order of matches
	for _, match := range matches {
		switch match.name {
		// Special case handling
		case "fuwamoco":
			sendEmbedFile(s, m, fmt.Sprintf("<@%s>", m.Author.ID), `images\high_res_baubau.mp4`)
			fwmcPair := OmgFuwaMocoEmojis[rand.Intn(len(OmgFuwaMocoEmojis))]
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[0])
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+fwmcPair[1])
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau1)
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+Bau2)
		case "fuwawa":
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(match.KVP.EmojiList[0]))
			if !matchedNames["mococo"] && !matchedNames["fuwamoco"] && !matchedNames["mocofuwa"] && !matchedNames["advent"] {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> **WHAT ABOUT MOCOCOEH!?**", m.Author.ID))
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{MococoHOEHSticker}})
			}
		case "lockin":
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(match.KVP.EmojiList[0]))
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{StickerIDs: []string{BanchouLockInSticker}})
		case "rokunana":
			if rand.Intn(4) == 0 {
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{Content: selectRandom(IHateMyselfForThisForgiveMe)})
			} else {
				s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+LihengzSus)
			}
		default:
			for _, emojis := range match.KVP.EmojiList {
				s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(emojis))
			}
		}
	}

	// Handle stickers
	for _, match := range matchedStickers {
		switch match.name {
		case "willnotbethere":
			s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(AngryEmojis))
			if rand.Intn(3) == 0 {
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{Content: fmt.Sprintf(`<@%s> **LOCK IN**`, m.Author.ID)})
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{Content: IrohaShotgun})
			} else {
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{Content: fmt.Sprintf(`<@%s> **LOCK IN**`, m.Author.ID), StickerIDs: []string{"1462523001957384426"}})
			}
		default:
			for _, emojis := range match.KVP.EmojiList {
				s.MessageReactionAdd(m.ChannelID, m.ID, "customemoji:"+selectRandom(emojis))
			}
		}
	}
}

// Randomly selects an element from a slice of strings
func selectRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func sendEmbedURL(s *discordgo.Session, m *discordgo.MessageCreate, message string, embedURL string) {
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: message,
		Embeds: []*discordgo.MessageEmbed{
			{
				Image: &discordgo.MessageEmbedImage{
					URL: embedURL,
				},
			},
		},
	})
}

func sendEmbedFile(s *discordgo.Session, m *discordgo.MessageCreate, message string, embedFilePath string) {
	file, err := os.Open(embedFilePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: message,
		Files: []*discordgo.File{
			{
				Name:   embedFilePath,
				Reader: file,
			},
		},
	})
}
