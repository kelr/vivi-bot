package main

import (
	"fmt"
	"regexp"
)

type OmgMemKVP struct {
	RegexExpr *regexp.Regexp
	EmojiList []string
}

func newHoloMemKVP(regexExpr string, emojiList []string) OmgMemKVP {
	return OmgMemKVP{
		RegexExpr: regexp.MustCompile(fmt.Sprintf(OmgHolomemGenericRegex, regexExpr)),
		EmojiList: emojiList,
	}
}

var (
	OmgHolomemGenericRegex = `(?i)\b(omf?g+[-\s]?(%s))\b`
	// Hard coded regex for special handling
	LockInRegexCompiled        = regexp.MustCompile(`(?i)\b(lock(?:ed|ing|s)?[-\s]?in)\b`)
	OmgAutoFisterRegexCompiled = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(autofister|ccgg))\b`)
	OmgFuwaMocoRegexCompiled   = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(fuwamoco|fwmc))\b`)
	OmgMocoFuwaRegexCompiled   = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?mocofuwa)\b`)
	OmgFuwawaRegexCompiled     = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?fuwawa)\b`)
	OmgMococoRegexCompiled     = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(mococo|mogogo|mocochan|mo(c|g)ojy?an))\b`)
	uuuuuCompiled              = regexp.MustCompile(`(?i)\b[u]{5,}\b`)

	// Mappings - Mem Name: Regex Pattern -> Emoji IDs List
	OmgMemNameMappings = map[string]OmgMemKVP{
		"suisei":   newHoloMemKVP("sui|suisei", OmgSuiEmojis), // Sui at the top to keep Kybo happy
		"fauna":    newHoloMemKVP("fauna", OmgFaunaEmojis),
		"pekora":   newHoloMemKVP("pe[kg]o(ra)?", OmgPekoEmojis),
		"achan":    newHoloMemKVP("achan", Placeholders),
		"akirose":  newHoloMemKVP("akirose(nthal)?", Placeholders),
		"ame":      newHoloMemKVP("ame+|amelia|amechan|watson", Placeholders),
		"anya":     newHoloMemKVP("anya", Placeholders),
		"aokun":    newHoloMemKVP("ao(-?kun)?", Placeholders),
		"aqua":     newHoloMemKVP("aqua|onion|ateshi|aqutan", Placeholders),
		"ayame":    newHoloMemKVP("ayame|ojou|nakiri", Placeholders),
		"azki":     newHoloMemKVP("azki(chi)?|azu?chan|azaz", Placeholders),
		"bae":      newHoloMemKVP("hakos|bae(lz)?|rat|chaos|hakotaro|peasant", OmgBaeEmojis),
		"bancho":   newHoloMemKVP("bancho|hajim(e|ya)", Placeholders),
		"biboo":    newHoloMemKVP("biboo|bijou|beejoe|beebs", Placeholders),
		"botan":    newHoloMemKVP("botan|shishiron|lalion", Placeholders),
		"calliope": newHoloMemKVP("calli(ope)?|mori|dad", Placeholders),
		"cecilia":  newHoloMemKVP("cecilia|cc|ceci|cece|clanker|cdollar", OmgCCEmojis),
		"chihaya":  newHoloMemKVP("chihaya|rindo", Placeholders),
		"chloe":    newHoloMemKVP("chloe|sakamata", Placeholders),
		"choco":    newHoloMemKVP("choco(sen(sei)?)?", Placeholders),
		"coco":     newHoloMemKVP("coco|kaichou", Placeholders),
		"flare":    newHoloMemKVP("flare|fuutan|(shira)?nui", Placeholders),
		"fubuki":   newHoloMemKVP("fbk|fubuki|foobs?|friend|fubuchan|shirakami|fubuking", Placeholders),
		"fuwawa":   newHoloMemKVP("fuwawa|fuwachan", OmgFuwawaEmojis),
		"gigi":     newHoloMemKVP("gigi|gg|geegee|fister|dafister|gpain", OmgGGEmojis),
		"gura":     newHoloMemKVP("gura|goom?ba|goob(idiba)?|same|gawr", Placeholders),
		"haachama": newHoloMemKVP("haachama|haato", Placeholders),
		"ina":      newHoloMemKVP("ina('nis)?", Placeholders),
		"iofi":     newHoloMemKVP("iofi(fteen)?|airani", Placeholders),
		"iroha":    newHoloMemKVP("iroha|gozaru(san|chan)?", Placeholders),
		"irys":     newHoloMemKVP("h?irys|hope", Placeholders),
		"kaela":    newHoloMemKVP("kaela|blacksmith|ela", Placeholders),
		"kanade":   newHoloMemKVP(`kanade|yellow\s?creature|otonose`, Placeholders),
		"kanata":   newHoloMemKVP(`kanata(n)?|pp\s?tenshi|gorilla`, Placeholders),
		"kiara":    newHoloMemKVP("kiara|wawa|tenchou|takanashi", Placeholders),
		"kobo":     newHoloMemKVP("kobo|kanaeru", Placeholders),
		"korone":   newHoloMemKVP("koro(ne|san)|doog|inugami", Placeholders),
		"koyori":   newHoloMemKVP("koyo(ri|te)?|hakui", Placeholders),
		"kronii":   newHoloMemKVP(`kronii|ouro|tam\s?gandr`, Placeholders),
		"lamy":     newHoloMemKVP("[lw]amy|yukihana", Placeholders),
		"laplus":   newHoloMemKVP("laplus|la+|lap(u-?)?chan|yamada", Placeholders),
		"liz":      newHoloMemKVP("liz|elizabeth|erb", Placeholders),
		"lui":      newHoloMemKVP("lui|looi", Placeholders),
		"luna":     newHoloMemKVP("luna", Placeholders),
		"marine":   newHoloMemKVP("marine|marin|maririn|baba|ahoy|senchou", Placeholders),
		"matsuri":  newHoloMemKVP("matsuri|natsuiro|god", Placeholders),
		"michiru":  newHoloMemKVP("michiru|(ihihi(-?san|-?chan))?", Placeholders),
		"miko":     newHoloMemKVP("mi[kg]o", Placeholders),
		"mio":      newHoloMemKVP("mio(sha)?", Placeholders),
		"moona":    newHoloMemKVP("moona|hoshinova", Placeholders),
		"mococo":   newHoloMemKVP("mococo|mogogo|mocochan|mo[cg]ojy?an", OmgMococoEmojis),
		"mumei":    newHoloMemKVP("mumei|moom(ers)?", Placeholders),
		"nene":     newHoloMemKVP("nene(chi)?", Placeholders),
		"nerissa":  newHoloMemKVP("(ne)?rissa", Placeholders),
		"niko":     newHoloMemKVP("niko(-?tan)|koganei", Placeholders),
		"nodoka":   newHoloMemKVP("nodoka", Placeholders),
		"noel":     newHoloMemKVP("noel", Placeholders),
		"okayu":    newHoloMemKVP("o[kg]ayu", Placeholders),
		"ollie":    newHoloMemKVP("ollie", Placeholders),
		"polka":    newHoloMemKVP("polka|(oma|pol)pol|omarun|zachou|polulu", Placeholders),
		"raden":    newHoloMemKVP("raden|juufuutei", Placeholders),
		"raora":    newHoloMemKVP("raora|rao", Placeholders),
		"reine":    newHoloMemKVP("reine|pavolia", Placeholders),
		"riona":    newHoloMemKVP("riona|isaki", Placeholders),
		"ririka":   newHoloMemKVP("ririkan?|ichijou", Placeholders),
		"risu":     newHoloMemKVP("risu|ayunda", Placeholders),
		"roboco":   newHoloMemKVP(`roboco([-\s]?san)?`, Placeholders),
		"sana":     newHoloMemKVP("sana(na)?", Placeholders),
		"sayaka":   newHoloMemKVP("sayaka|hanazono", Placeholders),
		"shion":    newHoloMemKVP("shion|garlic|murasaki", Placeholders),
		"shiori":   newHoloMemKVP("shiorin?", Placeholders),
		"sora":     newHoloMemKVP(`so[rd]a([-\s]?(chan|san))?`, Placeholders),
		"su":       newHoloMemKVP(`su([-\s]?chan)|mizumiya`, Placeholders),
		"subaru":   newHoloMemKVP("subaru|duck|oozora|shuba|subacchi", Placeholders),
		"towa":     newHoloMemKVP(`towa([-\s]sama)?`, OmgTowaEmojis),
		"vivi":     newHoloMemKVP("vivi|kikirara", Placeholders),
		"watame":   newHoloMemKVP("watame(lon)?|tsunomaki|sheep", Placeholders),
		"yuki":     newHoloMemKVP("yuki|kazeshiro", Placeholders),
		"zeta":     newHoloMemKVP("zeta|vestia", Placeholders)}
)
