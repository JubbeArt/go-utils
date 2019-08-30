package fchan

import (
	"strings"
)

type Board struct {
	Name        string
	Category    Category
	Title       string
	Description string
}

type Category int

const (
	CategoryJapaneseCulture Category = iota
	CategoryVideoGames
	CategoryInterests
	CategoryCreative
	CategoryOther
	CategoryMisc
	CategoryAdult
	CategoryHidden
)

func GetBoard(board string) (Board, bool) {
	board = strings.ToLower(board)

	for _, b := range Boards {
		if b.Name == board {
			return b, true
		}
	}
	return Board{}, false
}

func (c Category) String() string {
	switch c {
	case CategoryJapaneseCulture:
		return "Japanese Culture"
	case CategoryVideoGames:
		return "Video Games"
	case CategoryInterests:
		return "Interests"
	case CategoryCreative:
		return "Creative"
	case CategoryOther:
		return "Other"
	case CategoryMisc:
		return "Misc."
	case CategoryAdult:
		return "Adult"
	case CategoryHidden:
		return "Hidden"
	default:
		return ""
	}
}

var Boards = []Board{
	{Name: "a", Category: CategoryJapaneseCulture, Title: "Anime & Manga", Description: "Board dedicated to the discussion of Japanese animation and manga."},
	{Name: "c", Category: CategoryJapaneseCulture, Title: "Anime/Cute", Description: "Board for cute and moe anime images."},
	{Name: "w", Category: CategoryJapaneseCulture, Title: "Anime/Wallpapers", Description: "Board for posting Japanese anime wallpapers."},
	{Name: "m", Category: CategoryJapaneseCulture, Title: "Mecha", Description: "Board for discussing Japanese mecha robots and anime, like Gundam and Macross."},
	{Name: "cgl", Category: CategoryJapaneseCulture, Title: "Cosplay & EGL", Description: "Board for the discussion of cosplay, elegant gothic lolita (EGL), and anime conventions."},
	{Name: "cm", Category: CategoryJapaneseCulture, Title: "Cute/Male", Description: "Board for posting pictures of cute anime males."},
	{Name: "f", Category: CategoryJapaneseCulture, Title: "Flash", Description: "Board for sharing Adobe Flash files (SWFs)."},
	{Name: "n", Category: CategoryJapaneseCulture, Title: "Transportation", Description: "Board for discussing modes of transportation like trains and bicycles."},
	{Name: "jp", Category: CategoryJapaneseCulture, Title: "Otaku Culture", Description: "Board for discussing Japanese otaku culture."},

	{Name: "v", Category: CategoryVideoGames, Title: "Video Games", Description: "Board dedicated to the discussion of PC and console video games."},
	{Name: "vg", Category: CategoryVideoGames, Title: "Video Game Generals", Description: "Board dedicated to the discussion of PC and console video games."},
	{Name: "vp", Category: CategoryVideoGames, Title: "Pokémon", Description: "Board dedicated to discussing the Pokémon series of video games and shows."},
	{Name: "vr", Category: CategoryVideoGames, Title: "Retro Games", Description: "Board for discussing retro console video games and classic arcade games."},

	{Name: "co", Category: CategoryInterests, Title: "Comics & Cartoons", Description: "Board dedicated to the discussion of Western cartoons and comics."},
	{Name: "g", Category: CategoryInterests, Title: "Technology", Description: "Board for discussing computer hardware and software, programming, and general technology."},
	{Name: "tv", Category: CategoryInterests, Title: "Television & Film", Description: "Board dedicated to the discussion of television and film."},
	{Name: "k", Category: CategoryInterests, Title: "Weapons", Description: "Board for discussing all types of weaponry, from military tanks to guns and knives."},
	{Name: "o", Category: CategoryInterests, Title: "Auto", Description: "Board for discussing cars and motorcycles."},
	{Name: "an", Category: CategoryInterests, Title: "Animals & Nature", Description: "Board for posting pictures of animals, pets, and nature."},
	{Name: "tg", Category: CategoryInterests, Title: "Traditional Games", Description: "Board for discussing traditional gaming, such as board games and tabletop RPGs."},
	{Name: "sp", Category: CategoryInterests, Title: "Sports", Description: "Board for sports discussion."},
	{Name: "asp", Category: CategoryInterests, Title: "Alternative Sports & Wrestling", Description: "Board for the discussion of alternative and extreme sports such as wrestling and paintball."},
	{Name: "sci", Category: CategoryInterests, Title: "Science & Math", Description: "Board for the discussion of science and math."},
	{Name: "his", Category: CategoryInterests, Title: "History & Humanities", Description: "Board for discussing and debating history."},
	{Name: "int", Category: CategoryInterests, Title: "International", Description: "International board, for the exchange of foreign language and culture."},
	{Name: "out", Category: CategoryInterests, Title: "Outdoors", Description: "Board for discussing survivalist skills and outdoor activities such as hiking."},
	{Name: "toy", Category: CategoryInterests, Title: "Toys", Description: "Board for talking about all kinds of toys!"},

	{Name: "i", Category: CategoryCreative, Title: "Oekaki", Description: "Board for drawing and sharing art."},
	{Name: "po", Category: CategoryCreative, Title: "Papercraft & Origami", Description: "Board for posting papercraft and origami templates and instructions."},
	{Name: "p", Category: CategoryCreative, Title: "Photography", Description: "Board for sharing and critiquing photos."},
	{Name: "ck", Category: CategoryCreative, Title: "Food & Cooking", Description: "Board for food pictures and cooking recipes."},
	{Name: "ic", Category: CategoryCreative, Title: "Artwork/Critique", Description: "Board for the discussion and critique of art."},
	{Name: "wg", Category: CategoryCreative, Title: "Wallpapers/General", Description: "Board for posting general wallpapers."},
	{Name: "lit", Category: CategoryCreative, Title: "Literature", Description: "Board for the discussion of books, authors, and literature."},
	{Name: "mu", Category: CategoryCreative, Title: "Music", Description: "Board for discussing all types of music."},
	{Name: "fa", Category: CategoryCreative, Title: "Fashion", Description: "Board for images and discussion relating to fashion and apparel."},
	{Name: "3", Category: CategoryCreative, Title: "3DCG", Description: "Board for 3D modeling and imagery."},
	{Name: "gd", Category: CategoryCreative, Title: "Graphic Design", Description: "Board for graphic design."},
	{Name: "diy", Category: CategoryCreative, Title: "Do It Yourself", Description: "Board for DIY/do it yourself projects, home improvement, and makers."},
	{Name: "wsg", Category: CategoryCreative, Title: "Worksafe GIF", Description: "Board dedicated to sharing worksafe animated GIFs and WEBMs."},
	{Name: "qst", Category: CategoryCreative, Title: "Quests", Description: "Board for grinding XP."},

	{Name: "biz", Category: CategoryOther, Title: "Business & Finance", Description: "Board for the discussion of business and finance, and cryptocurrencies such as Bitcoin and Dogecoin."},
	{Name: "trv", Category: CategoryOther, Title: "Travel", Description: "Board dedicated to travel and the countries of the world."},
	{Name: "fit", Category: CategoryOther, Title: "Fitness", Description: "Board for weightlifting, health, and fitness."},
	{Name: "x", Category: CategoryOther, Title: "Paranormal", Description: "Board for the discussion of paranormal, spooky pictures and conspiracy theories."},
	{Name: "adv", Category: CategoryOther, Title: "Advice", Description: "Board for giving and receiving advice. "},
	{Name: "lgbt", Category: CategoryOther, Title: "LGBT", Description: "Board for Lesbian-Gay-Bisexual-Transgender-Queer and sexuality discussion."},
	{Name: "mlp", Category: CategoryOther, Title: "Pony", Description: "Board dedicated to the discussion of My Little Pony: Friendship is Magic."},
	{Name: "news", Category: CategoryOther, Title: "Current News", Description: "Board for current news. "},
	{Name: "wsr", Category: CategoryOther, Title: "Worksafe Requests", Description: "Board dedicated to fulfilling non-NSFW requests."},
	{Name: "vip", Category: CategoryOther, Title: "Very Important Posts", Description: "Board for Pass users."},

	{Name: "b", Category: CategoryMisc, Title: "Random", Description: "The birthplace of Anonymous, and where people go to discuss random topics and create memes on 4chan."},
	{Name: "r9k", Category: CategoryMisc, Title: "ROBOT9001", Description: "Board for hanging out and posting greentext stories."},
	{Name: "pol", Category: CategoryMisc, Title: "Politically Incorrect", Description: "Board for discussing and debating politics and current events."},
	{Name: "bant", Category: CategoryMisc, Title: "International/Random", Description: "International hanging out board, where you can have fun with Anonymous all over the world."},
	{Name: "soc", Category: CategoryMisc, Title: "Cams & Meetups", Description: "Board for camwhores and meetups."},
	{Name: "s4s", Category: CategoryMisc, Title: "Shit 4chan Says", Description: "Board for posting dank memes :^)"},

	{Name: "s", Category: CategoryAdult, Title: "Sexy Beautiful Women", Description: "Board dedicated to sharing images of softcore pornography."},
	{Name: "hc", Category: CategoryAdult, Title: "Hardcore", Description: "Board for the posting of adult hardcore pornography."},
	{Name: "hm", Category: CategoryAdult, Title: "Handsome Men", Description: "Board dedicated to sharing adult images of handsome men."},
	{Name: "h", Category: CategoryAdult, Title: "Hentai", Description: "Board for adult Japanese anime hentai images."},
	{Name: "e", Category: CategoryAdult, Title: "Ecchi", Description: "Board for suggestive (ecchi) hentai images."},
	{Name: "u", Category: CategoryAdult, Title: "Yuri", Description: "Board for yuri hentai images."},
	{Name: "d", Category: CategoryAdult, Title: "Hentai/Alternative", Description: "Board for alternative hentai images."},
	{Name: "y", Category: CategoryAdult, Title: "Yaoi", Description: "Board for posting adult yaoi hentai images."},
	{Name: "t", Category: CategoryAdult, Title: "Torrents", Description: "Board for posting links and descriptions to torrents."},
	{Name: "hr", Category: CategoryAdult, Title: "High Resolution", Description: "Board for the sharing of high resolution images."},
	{Name: "gif", Category: CategoryAdult, Title: "Adult GIF", Description: "Board dedicated to animated adult GIFs and WEBMs."},
	{Name: "aco", Category: CategoryAdult, Title: "Adult Cartoons", Description: "Board for posting western-styled adult cartoon artwork."},
	{Name: "r", Category: CategoryAdult, Title: "Adult Requests", Description: "Board dedicated to fulfilling all types of user requests."},

	{Name: "qa", Category: CategoryHidden, Title: "Question & Answer", Description: "Board for question and answer threads."},
	{Name: "trash", Category: CategoryHidden, Title: "Off-Topic", Description: "Board jail for off-topic threads."},
}
