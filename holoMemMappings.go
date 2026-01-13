package main

import (
	"fmt"
	"regexp"
)

type OmgMemKVP struct {
	RegexExpr *regexp.Regexp
	EmojiList []string
}

func newOmgMemKVP(regexExpr string, emojiList []string) OmgMemKVP {
	return OmgMemKVP{
		RegexExpr: regexp.MustCompile(fmt.Sprintf(OmgMemGenericRegex, regexExpr)),
		EmojiList: emojiList,
	}
}

var (
	OmgMemGenericRegex = `(?i)\b(omf?g+[-\s]?(%s))\b`
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
		"suisei":   newOmgMemKVP("sui|suisei", OmgSuiEmojis), // Sui at the top to keep Kybo happy
		"fauna":    newOmgMemKVP("fauna", OmgFaunaEmojis),
		"nimi":     newOmgMemKVP("nimi|tapir", OmgNimiEmojis),
		"pekora":   newOmgMemKVP("pe[kg]o(ra)?", OmgPekoEmojis),
		"achan":    newOmgMemKVP("achan", OmgAchanEmojis),
		"akirose":  newOmgMemKVP("akirose(nthal)?", OmgAkiroseEmojis),
		"ame":      newOmgMemKVP("ame+|amelia|amechan|watson", OmgAmeEmojis),
		"anya":     newOmgMemKVP("anya", OmgAnyaEmojis),
		"aokun":    newOmgMemKVP("ao(-?kun)?", OmgAokunEmojis),
		"aqua":     newOmgMemKVP("aqua|onion|ateshi|aqutan", OmgAquaEmojis),
		"ayame":    newOmgMemKVP("ayame|ojou|nakiri", OmgAyameEmojis),
		"azki":     newOmgMemKVP("azki(chi)?|azu?chan|azaz", OmgAzkiEmojis),
		"bae":      newOmgMemKVP("hakos|bae(lz)?|rat|chaos|hakotaro|peasant", OmgBaeEmojis),
		"bancho":   newOmgMemKVP("bancho|hajim(e|ya)", OmgBanchoEmojis),
		"biboo":    newOmgMemKVP("biboo|bijou|beejoe|beebs", OmgBibooEmojis),
		"botan":    newOmgMemKVP("botan|shishiron|lalion", OmgBotanEmojis),
		"calliope": newOmgMemKVP("calli(ope)?|mori|dad", OmgCalliopeEmojis),
		"cecilia":  newOmgMemKVP("cecilia|cc|ceci|cece|clanker|cdollar", OmgCCEmojis),
		"chihaya":  newOmgMemKVP("chihaya|rindo", OmgChihayaEmojis),
		"chloe":    newOmgMemKVP("chloe|sakamata", OmgChloeEmojis),
		"choco":    newOmgMemKVP("choco(sen(sei)?)?", OmgChocoEmojis),
		"coco":     newOmgMemKVP("coco|kaichou", OmgCocoEmojis),
		"dooby":    newOmgMemKVP("dooby?|doobert", OmgDoobyEmojis),
		"flare":    newOmgMemKVP("flare|fuutan|(shira)?nui", OmgFlareEmojis),
		"fubuki":   newOmgMemKVP("fbk|fubuki|foobs?|friend|fubuchan|shirakami|fubuking", OmgFubukiEmojis),
		"gigi":     newOmgMemKVP("gigi|gg|geegee|fister|dafister|gpain", OmgGGEmojis),
		"gura":     newOmgMemKVP("gura|goom?ba|goob(idiba)?|same|gawr", OmgGuraEmojis),
		"haachama": newOmgMemKVP("haachama|haato", OmgHaachamaEmojis),
		"ina":      newOmgMemKVP("ina('nis)?|wah|tako", OmgInaEmojis),
		"iofi":     newOmgMemKVP("iofi(fteen)?|airani", OmgIofiEmojis),
		"iroha":    newOmgMemKVP("iroha|gozaru(san|chan)?", OmgIrohaEmojis),
		"irys":     newOmgMemKVP("h?irys|hope", OmgIrysEmojis),
		"kaela":    newOmgMemKVP("kaela|blacksmith|ela", OmgKaelaEmojis),
		"kanade":   newOmgMemKVP(`kanade|yellow\s?creature|otonose`, OmgKanadeEmojis),
		"kanata":   newOmgMemKVP(`kanata(n)?|pp\s?tenshi|gorilla`, OmgKanataEmojis),
		"kiara":    newOmgMemKVP("kiara|wawa|tenchou|takanashi", OmgKiaraEmojis),
		"kobo":     newOmgMemKVP("kobo|kanaeru|bokoboko", OmgKoboEmojis),
		"korone":   newOmgMemKVP("koro(ne|san)|doog|inugami", OmgKoroneEmojis),
		"koyori":   newOmgMemKVP("koyo(ri|te)?|hakui", OmgKoyoriEmojis),
		"kronii":   newOmgMemKVP(`kronii|ouro|tam\s?gandr`, OmgKroniiEmojis),
		"lamy":     newOmgMemKVP("[lw]amy|yukihana", OmgLamyEmojis),
		"laplus":   newOmgMemKVP("laplus|la+|lap(u-?)?chan|yamada", OmgLaplusEmojis),
		"liz":      newOmgMemKVP("liz|elizabeth|erb", OmgLizEmojis),
		"lui":      newOmgMemKVP("lui|looi", OmgLuiEmojis),
		"luna":     newOmgMemKVP("luna", OmgLunaEmojis),
		"marine":   newOmgMemKVP("marine|marin|maririn|baba|ahoy|senchou", OmgMarineEmojis),
		"matsuri":  newOmgMemKVP("matsuri|natsuiro|god", OmgMatsuriEmojis),
		"mel":      newOmgMemKVP("mel|yozora|banpire", OmgMelEmojis),
		"michiru":  newOmgMemKVP("michiru|ihihi(-?san|-?chan)?", OmgMichiruEmojis),
		"miko":     newOmgMemKVP("mi[kg]o", OmgMikoEmojis),
		"mio":      newOmgMemKVP("mio(sha)?", OmgMioEmojis),
		"mint":     newOmgMemKVP("minto?", OmgMintEmojis),
		"moona":    newOmgMemKVP("moona|hoshinova", OmgMoonaEmojis),
		"mococo":   newOmgMemKVP("mococo|mogogo|mocochan|mo[cg]ojy?an", OmgMococoEmojis),
		"mumei":    newOmgMemKVP("mumei|moom(ers)?", OmgMumeiEmojis),
		"nene":     newOmgMemKVP("nene(chi)?", OmgNeneEmojis),
		"nerissa":  newOmgMemKVP("(ne)?rissa", OmgNerissaEmojis),
		"niko":     newOmgMemKVP("niko(-?tan)|koganei", OmgNikoEmojis),
		"nodoka":   newOmgMemKVP("nodoka", OmgNodokaEmojis),
		"noel":     newOmgMemKVP("noel", OmgNoelEmojis),
		"okayu":    newOmgMemKVP("o[kg]ayu", OmgOkayuEmojis),
		"ollie":    newOmgMemKVP("ollie", OmgOllieEmojis),
		"polka":    newOmgMemKVP("polka|(oma|pol)pol|omarun|zachou|polulu", OmgPolkaEmojis),
		"raden":    newOmgMemKVP("raden|juufuutei", OmgReineEmojis),
		"raora":    newOmgMemKVP("raora|rao", OmgRaoraEmojis),
		"reine":    newOmgMemKVP("reine|pavolia", OmgReineEmojis),
		"riona":    newOmgMemKVP("riona|isaki", OmgRionaEmojis),
		"ririka":   newOmgMemKVP("ririkan?|ichijou", OmgRirikaEmojis),
		"risu":     newOmgMemKVP("risu|ayunda", OmgRisuEmojis),
		"roboco":   newOmgMemKVP(`roboco([-\s]?san)?`, OmgRobocoEmojis),
		"saba":     newOmgMemKVP("saba|fish|feesh", OmgSabaEmojis),
		"sana":     newOmgMemKVP("sana(na)?", OmgSanaEmojis),
		"sayaka":   newOmgMemKVP("sayaka|hanazono", OmgSayakaEmojis),
		"shion":    newOmgMemKVP("shion|garlic|murasaki", OmgShionEmojis),
		"shiori":   newOmgMemKVP("shiorin?", OmgShioriEmojis),
		"sora":     newOmgMemKVP(`so[rd]a([-\s]?(chan|san))?`, OmgSoraEmojis),
		"su":       newOmgMemKVP(`su{1,2}([-\s]?chan)?|mizumiya`, OmgSuEmojis),
		"subaru":   newOmgMemKVP("subaru|duck|oozora|shuba|subacchi", OmgSubaruEmojis),
		"towa":     newOmgMemKVP(`towa([-\s]?sama)?`, OmgTowaEmojis),
		"vivi":     newOmgMemKVP("vivi|kikirara", OmgViviEmojis),
		"watame":   newOmgMemKVP("watame(lon)?|tsunomaki|sheep", OmgWatameEmojis),
		"yuki":     newOmgMemKVP("yuki|kazeshiro", OmgYukiEmojis),
		"zeta":     newOmgMemKVP("zeta|vestia", OmgZetaEmojis)}
)
