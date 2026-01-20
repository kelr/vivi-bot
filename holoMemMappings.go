package main

import (
	"fmt"
	"regexp"
)

type RegExpEmojiKVP struct {
	RegexExpr *regexp.Regexp
	EmojiList [][]string
}

func newOmgMemKVP(regexExpr string, emojiList [][]string) RegExpEmojiKVP {
	return RegExpEmojiKVP{
		RegexExpr: regexp.MustCompile(fmt.Sprintf(OmgMemGenericRegex, regexExpr)),
		EmojiList: emojiList,
	}
}

func newPhraseKVP(regexExpr string, emojiList [][]string) RegExpEmojiKVP {
	return RegExpEmojiKVP{
		RegexExpr: regexp.MustCompile(fmt.Sprintf(RegularPhraseRegex, regexExpr)),
		EmojiList: emojiList,
	}
}

type RegexMatch struct {
	name string
	idx  int
	KVP  RegExpEmojiKVP
}

var (
	OmgMemGenericRegex = `(?i)\b(omf?g+[-\s]?(%s))\b`
	RegularPhraseRegex = `(?i)\b(%s)\b`
	// Hard coded regex for special handling
	OmgSuiRegexCompiled = regexp.MustCompile(`(?i)\b(omf?g+[-\s]?(sui|suisei|comet))\b`)

	// Mappings - Mem Name: Regex Pattern -> Emoji IDs List
	OmgMemNameMappings = map[string]RegExpEmojiKVP{
		"achan":      newOmgMemKVP("achan", [][]string{OmgAchanEmojis}),
		"akirose":    newOmgMemKVP("akirose(nthal)?", [][]string{OmgAkiroseEmojis}),
		"ame":        newOmgMemKVP("ame+|amelia|amechan|watson", [][]string{OmgAmeEmojis}),
		"anya":       newOmgMemKVP("anya", [][]string{OmgAnyaEmojis}),
		"aokun":      newOmgMemKVP("ao(-?kun)?", [][]string{OmgAokunEmojis}),
		"aqua":       newOmgMemKVP("aqua|onion|ateshi|a[qk]utan", [][]string{OmgAquaEmojis}),
		"autofister": newOmgMemKVP("autofister|ccgg|cecigigi", [][]string{[]string{FluffyCC}, []string{FuzzyGG}}),
		"ayame":      newOmgMemKVP("ayame|ojou|nakiri", [][]string{OmgAyameEmojis}),
		"azki":       newOmgMemKVP("azki(chi)?|azu?chan|azaz", [][]string{OmgAzkiEmojis}),
		"bae":        newOmgMemKVP("hakos|bae(lz)?|rat|chaos|hakotaro|peasant", [][]string{OmgBaeEmojis}),
		"bancho":     newOmgMemKVP("banchou?|hajim(e|ya)", [][]string{OmgBanchoEmojis}),
		"biboo":      newOmgMemKVP("biboo|bijou|beejoe|beebs|koseki", [][]string{OmgBibooEmojis}),
		"botan":      newOmgMemKVP("botan|shishiron|lalion|lion", [][]string{OmgBotanEmojis}),
		"calliope":   newOmgMemKVP("calli(ope)?|mori|dad", [][]string{OmgCalliopeEmojis}),
		"cecilia":    newOmgMemKVP("cecilia|cc|ceci|clanker|cdollar|immerhater|immergreen|automaton", [][]string{OmgCCEmojis}),
		"chihaya":    newOmgMemKVP("chihaya|rindo", [][]string{OmgChihayaEmojis}),
		"chloe":      newOmgMemKVP("chloe|sakamata", [][]string{OmgChloeEmojis}),
		"choco":      newOmgMemKVP("choco(sen(sei)?)?", [][]string{OmgChocoEmojis}),
		"coco":       newOmgMemKVP("coco|kaichou", [][]string{OmgCocoEmojis}),
		"dooby":      newOmgMemKVP("dooby?|doobert", [][]string{OmgDoobyEmojis}),
		"fauna":      newOmgMemKVP("fauna|faufau|kirin|konfauna", [][]string{OmgFaunaEmojis}),
		"flare":      newOmgMemKVP("flare|fuutan|(shira)?nui", [][]string{OmgFlareEmojis}),
		"fubuki":     newOmgMemKVP("fbk|fubuki|foobs?|friend|fubuchan|shirakami|fubuking", [][]string{OmgFubukiEmojis}),
		"fuwamoco":   newOmgMemKVP(`fuwamoco|fwmc|bau[\s]?bau`, [][]string{}),
		"mocofuwa":   newOmgMemKVP("mocofuwa", [][]string{[]string{MococoDoro}, []string{FuwawaDoro}}),
		"fuwawa":     newOmgMemKVP("fuwawa", [][]string{OmgFuwawaEmojis}),
		"gigi":       newOmgMemKVP("gigi|gg|geegee|fister|dafister|gpain", [][]string{OmgGGEmojis}),
		"gura":       newOmgMemKVP("gura|goom?ba|goob(idiba)?|same|gawr|shark", [][]string{OmgGuraEmojis}),
		"haachama":   newOmgMemKVP("haachama|haato|chama", [][]string{OmgHaachamaEmojis}),
		"ina":        newOmgMemKVP("ina('nis)?|wah|tako|inya", [][]string{OmgInaEmojis}),
		"iofi":       newOmgMemKVP("iofi(fteen)?|airani", [][]string{OmgIofiEmojis}),
		"iroha":      newOmgMemKVP("iroha|gozaru(san|chan)?", [][]string{OmgIrohaEmojis}),
		"irys":       newOmgMemKVP("h?irys|hope", [][]string{OmgIrysEmojis}),
		"kaela":      newOmgMemKVP("kaela|blacksmith|ela", [][]string{OmgKaelaEmojis}),
		"kanade":     newOmgMemKVP(`kanade|yellow\s?creature|otonose|getsuyoubi|omae+|omaera`, [][]string{OmgKanadeEmojis}),
		"kanata":     newOmgMemKVP(`kanata(n)?|pp\s?tenshi|gorilla|ppt`, [][]string{OmgKanataEmojis}),
		"kiara":      newOmgMemKVP("kiara|wawa|tenchou|takanashi", [][]string{OmgKiaraEmojis}),
		"kobo":       newOmgMemKVP("kobo|kanaeru|bokoboko", [][]string{OmgKoboEmojis}),
		"korone":     newOmgMemKVP(`korone|koro[-\s]?san|doog|inugami|koone`, [][]string{OmgKoroneEmojis}),
		"koyori":     newOmgMemKVP("koyo(ri|te)?|hakui", [][]string{OmgKoyoriEmojis}),
		"kronii":     newOmgMemKVP(`kronii|ouro|tam\s?gandr|clock|tam\s?tender|ribbon|kronini`, [][]string{OmgKroniiEmojis}),
		"lamy":       newOmgMemKVP("[lw]amy|yukihana", [][]string{OmgLamyEmojis}),
		"laplus":     newOmgMemKVP("laplus|la+|lap(u-?)?chan|yamada", [][]string{OmgLaplusEmojis}),
		"liz":        newOmgMemKVP("liz|elizabeth|erb", [][]string{OmgLizEmojis}),
		"lui":        newOmgMemKVP("lui|looi", [][]string{OmgLuiEmojis}),
		"luna":       newOmgMemKVP("luna", [][]string{OmgLunaEmojis}),
		"marine":     newOmgMemKVP("marine|marin|maririn|baba|ahoy|senchou", [][]string{OmgMarineEmojis}),
		"matsuri":    newOmgMemKVP("matsuri|natsuiro|god", [][]string{OmgMatsuriEmojis}),
		"mel":        newOmgMemKVP("mel|yozora|banpire", [][]string{OmgMelEmojis}),
		"michiru":    newOmgMemKVP("michiru|ihihi(-?san|-?chan)?", [][]string{OmgMichiruEmojis}),
		"miko":       newOmgMemKVP("mi[kg]o(chi)?", [][]string{OmgMikoEmojis}),
		"mint":       newOmgMemKVP("minto?|minki", [][]string{OmgMintEmojis}),
		"mio":        newOmgMemKVP("mio(sha)?", [][]string{OmgMioEmojis}),
		"mococo":     newOmgMemKVP("mococo|mogogo|mocochan|mo[cg]ojy?an", [][]string{OmgMococoEmojis}),
		"moona":      newOmgMemKVP("moona|hoshinova", [][]string{OmgMoonaEmojis}),
		"mumei":      newOmgMemKVP("mumei|moom(ers)?", [][]string{OmgMumeiEmojis}),
		"nene":       newOmgMemKVP("nene(chi)?", [][]string{OmgNeneEmojis}),
		"nerissa":    newOmgMemKVP(`(ne)?rissa|wet\s?cat`, [][]string{OmgNerissaEmojis}),
		"niko":       newOmgMemKVP("niko(-?tan)|koganei", [][]string{OmgNikoEmojis}),
		"nimi":       newOmgMemKVP("nimi|tapir", [][]string{OmgNimiEmojis}),
		"nodoka":     newOmgMemKVP("nodoka", [][]string{OmgNodokaEmojis}),
		"noel":       newOmgMemKVP("noel", [][]string{OmgNoelEmojis}),
		"okayu":      newOmgMemKVP("o[kg]ayu", [][]string{OmgOkayuEmojis}),
		"ollie":      newOmgMemKVP("ollie", [][]string{OmgOllieEmojis}),
		"pekora":     newOmgMemKVP("pe[kg]o(ra)?", [][]string{OmgPekoEmojis}),
		"polka":      newOmgMemKVP("polka|(oma|pol)pol|omarun|zachou|polulu", [][]string{OmgPolkaEmojis}),
		"raden":      newOmgMemKVP("raden|juufuutei", [][]string{OmgReineEmojis}),
		"raora":      newOmgMemKVP("raora|rao", [][]string{OmgRaoraEmojis}),
		"reine":      newOmgMemKVP("reine|pavolia", [][]string{OmgReineEmojis}),
		"riona":      newOmgMemKVP("riona|isaki", [][]string{OmgRionaEmojis}),
		"ririka":     newOmgMemKVP("ririkan?|ichijou", [][]string{OmgRirikaEmojis}),
		"risu":       newOmgMemKVP("risu|ayunda", [][]string{OmgRisuEmojis}),
		"rushia":     newOmgMemKVP("rushia", [][]string{OmgRushiaEmojis}),
		"roboco":     newOmgMemKVP(`roboco([-\s]?san)?|rbc`, [][]string{OmgRobocoEmojis}),
		"saba":       newOmgMemKVP("saba|fish|feesh", [][]string{OmgSabaEmojis}),
		"sana":       newOmgMemKVP("sana(na)?", [][]string{OmgSanaEmojis}),
		"sayaka":     newOmgMemKVP("sayaka|hanazono", [][]string{OmgSayakaEmojis}),
		"shion":      newOmgMemKVP("shion|garlic|murasaki", [][]string{OmgShionEmojis}),
		"shiori":     newOmgMemKVP("shiorin?", [][]string{OmgShioriEmojis}),
		"sora":       newOmgMemKVP(`so[rd]a([-\s]?(chan|san))?`, [][]string{OmgSoraEmojis}),
		"su":         newOmgMemKVP(`su{1,2}([-\s]?chan)?|mizumiya`, [][]string{OmgSuEmojis}),
		"subaru":     newOmgMemKVP("subaru|duck|oozora|shuba|subacchi", [][]string{OmgSubaruEmojis}),
		"towa":       newOmgMemKVP(`towa([-\s]?sama)?|tmd|devil|debiru`, [][]string{OmgTowaEmojis}),
		"vivi":       newOmgMemKVP("vivi|kikirara", [][]string{OmgViviEmojis}),
		"watame":     newOmgMemKVP("watame(lon)?|tsunomaki|sheep", [][]string{OmgWatameEmojis}),
		"yuki":       newOmgMemKVP("yuki|kazeshiro", [][]string{OmgYukiEmojis}),
		"zeta":       newOmgMemKVP("zeta|vestia", [][]string{OmgZetaEmojis}),

		"bloodraven": newOmgMemKVP("bloodraven", [][]string{OmgLizEmojis, OmgNerissaEmojis}),
		"novelflame": newOmgMemKVP("novelflame", [][]string{OmgLizEmojis, OmgShioriEmojis}),
		"amesame":    newOmgMemKVP("amesame", [][]string{OmgGuraEmojis, OmgAmeEmojis}),
		"baerys":     newOmgMemKVP("baerys", [][]string{OmgBaeEmojis, OmgIrysEmojis}),
		"kronmei":    newOmgMemKVP("kronmei", [][]string{OmgKroniiEmojis, OmgMumeiEmojis}),
		"faunmei":    newOmgMemKVP("faunmei", [][]string{OmgFaunaEmojis, OmgMumeiEmojis}),
		"kronfau":    newOmgMemKVP("kronfau", [][]string{OmgKroniiEmojis, OmgFaunaEmojis}),
		"takamori":   newOmgMemKVP("takamori", [][]string{OmgKiaraEmojis, OmgCalliopeEmojis}),
		"micomet":    newOmgMemKVP("micomet|business", [][]string{OmgSuiEmojis, OmgMikoEmojis}),
		"okakoro":    newOmgMemKVP("okakoro", [][]string{OmgOkayuEmojis, OmgKoroneEmojis}),
		"noefure":    newOmgMemKVP("noefure", [][]string{OmgNoelEmojis, OmgFlareEmojis}),
		"aqupeko":    newOmgMemKVP("aqupeko|akupeko", [][]string{OmgAquaEmojis, OmgPekoEmojis}),

		"myth":       newOmgMemKVP("myth", [][]string{OmgInaEmojis, OmgKiaraEmojis, OmgAmeEmojis, OmgCalliopeEmojis, OmgGuraEmojis}),
		"promise":    newOmgMemKVP("promise", [][]string{OmgFaunaEmojis, OmgIrysEmojis, OmgBaeEmojis, OmgKroniiEmojis, OmgMumeiEmojis}),
		"council":    newOmgMemKVP("council", [][]string{OmgFaunaEmojis, OmgSanaEmojis, OmgBaeEmojis, OmgKroniiEmojis, OmgMumeiEmojis}),
		"councilrys": newOmgMemKVP("councilrys", [][]string{OmgFaunaEmojis, OmgSanaEmojis, OmgBaeEmojis, OmgKroniiEmojis, OmgMumeiEmojis, OmgIrysEmojis}),
		"advent":     newOmgMemKVP("advent", [][]string{OmgShioriEmojis, OmgBibooEmojis, OmgNerissaEmojis, OmgFuwawaEmojis, OmgMococoEmojis}),
		"justice":    newOmgMemKVP("justice", [][]string{OmgLizEmojis, OmgGGEmojis, OmgCCEmojis, OmgRaoraEmojis}),
		"gamers":     newOmgMemKVP("gamers", [][]string{OmgFubukiEmojis, OmgMioEmojis, OmgOkayuEmojis, OmgKoroneEmojis}),
		"flowglow":   newOmgMemKVP("flowglow", [][]string{OmgNikoEmojis, OmgSuEmojis, OmgRionaEmojis, OmgViviEmojis, OmgChihayaEmojis}),
		"regloss":    newOmgMemKVP("regloss", [][]string{OmgRirikaEmojis, OmgBanchoEmojis, OmgKanadeEmojis, OmgAokunEmojis, OmgRadenEmojis}),
		"gen0":       newOmgMemKVP(`gen[-\s]?0|(zero|0)kisei`, [][]string{OmgSoraEmojis, OmgRobocoEmojis, OmgAzkiEmojis, OmgMikoEmojis, OmgSuiEmojis}),
		"gen1":       newOmgMemKVP(`gen[-\s]?1|(ik|1)kisei`, [][]string{OmgAkiroseEmojis, OmgHaachamaEmojis, OmgFubukiEmojis, OmgMatsuriEmojis}),
		"gen2":       newOmgMemKVP(`gen[-\s]?2|(ni|2)kisei`, [][]string{OmgAyameEmojis, OmgChocoEmojis, OmgSubaruEmojis, OmgAquaEmojis, OmgShionEmojis}),
		"gen3":       newOmgMemKVP(`gen[-\s]?3|(san|3)kisei`, [][]string{OmgRushiaEmojis, OmgPekoEmojis, OmgFlareEmojis, OmgNoelEmojis, OmgMarineEmojis}),
		"gen4":       newOmgMemKVP(`gen[-\s]?4|(yon|4)kisen`, [][]string{OmgWatameEmojis, OmgTowaEmojis, OmgLunaEmojis, OmgKanataEmojis, OmgCocoEmojis}),
		"gen5":       newOmgMemKVP(`gen[-\s]?5|(go|5)kisei`, [][]string{OmgLamyEmojis, OmgNeneEmojis, OmgBotanEmojis, OmgPolkaEmojis}),
		"holox":      newOmgMemKVP(`holox|gen[-\s]?6|(rok|6)kisei`, [][]string{OmgLaplusEmojis, OmgLuiEmojis, OmgKoyoriEmojis, OmgIrohaEmojis, OmgChloeEmojis}),
		"holoid":     newOmgMemKVP(`holo[-\s]?id`, [][]string{OmgRisuEmojis, OmgMoonaEmojis, OmgIofiEmojis, OmgOllieEmojis, OmgAnyaEmojis, OmgReineEmojis, OmgZetaEmojis, OmgKaelaEmojis, OmgKoboEmojis}),

		"stupidnonsense": newOmgMemKVP("not marine|(pizza|green|blue|pink|yellow) [wl]amy", [][]string{Placeholders}),
		"glue":           newOmgMemKVP("glue", [][]string{{GlueEmoji}}),
		"faunauuuuu":     newPhraseKVP(`[u]{3,}`, [][]string{{FaunaUUUUU}}),
		"lockin":         newPhraseKVP(`lock(?:ed|ing|s)?[-\s]?in`, [][]string{{BaeLock, IRySLock}}),
		"haeh":           newPhraseKVP(`haeh`, [][]string{{Haeh}}),
		"hoeh":           newPhraseKVP(`hoeh`, [][]string{{MococoHOEH, MococoHOEH2, MococoHOEH3}}),
		"jdon":           newPhraseKVP(`jdon|jdon\s?my\s?soul`, [][]string{{JDONMYSOUL}, {JDON}}),
		"zaiko":          newPhraseKVP(`zaiko`, [][]string{{Zaiko}}),
		"eplus":          newPhraseKVP(`eplus|e+`, [][]string{{EPlus}}),
		"ticketmaster":   newPhraseKVP(`ticketmaster`, [][]string{{TicketMaster}}),
		"ltike":          newPhraseKVP(`ltike`, [][]string{{ltike}}),
		"zenloss":        newPhraseKVP(`zenloss`, [][]string{{Zaiko}, {EPlus}, {TicketMaster}, {ltike}}),
		"rokunana":       newPhraseKVP(`roku\s?nana|6\s?7`, [][]string{}),
	}

	// Mappings for phrases to check for and local files to react with
	FileEmbedMappings = map[string]RegExpEmojiKVP{
		"wizardojisan":  newOmgMemKVP(`wizard\s?ojisan|ojisan`, [][]string{{`images\BirdsTookOldMan.mp4`, `images\OldManSpinning.mp4`, `images\OldManSun.mp4`}}),
		"mococar":       newPhraseKVP(`mococar`, [][]string{{`images\mococar.mp4`}}),
		"pregnant":      newPhraseKVP(`pregnant`, [][]string{{`images\Mint_GET_PREGNANT.mp4`}}),
		"ietaiga":       newPhraseKVP(`ietaiga`, [][]string{{`images\ietaiga.mp4`}}),
		"hoeh":          newPhraseKVP(`hoeh`, [][]string{{`images\HOEH.mp4`}}),
		"freak":         newPhraseKVP(`freak`, [][]string{{`images\nimi_youfreak.mp4`}}),
		"manilovefauna": newPhraseKVP(`man i love fauna|i miss fauna`, [][]string{{`images\manilovefauna.mp4`}}),
	}

	// Mappings for StickerIds and emojis to react to it with
	StickerIdMappings = map[string]RegExpEmojiKVP{
		"getsuyoubi":     RegExpEmojiKVP{RegexExpr: regexp.MustCompile(`^1363559590603657538$`), EmojiList: [][]string{KanadeSmugEmojis, MukaMukaEmojis}},
		"lockin":         RegExpEmojiKVP{RegexExpr: regexp.MustCompile(`^(1447975441183936736|1447412433273491518)$`), EmojiList: [][]string{{BaeLock, IRySLock}}},
		"willnotbethere": RegExpEmojiKVP{RegexExpr: regexp.MustCompile(`^(1447412609165951066|1447412724484145204)$`), EmojiList: [][]string{}},
	}
)
