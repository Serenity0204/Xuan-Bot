package bot

import (
	"fmt"

	"github.com/Serenity0204/Xuan-Bot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func InitMsg() {

	goBot, e := discordgo.New("Bot " + config.Token)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	user, e := goBot.User("@me")

	if e != nil {
		fmt.Println(e.Error())
		return
	}
	BotID = user.ID
	goBot.AddHandler(Handlers)

	e = goBot.Open()

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func Handlers(ses *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == BotID {
		return
	}
	if msg.Content == "!s6" {
		_, _ = ses.ChannelMessageSend(msg.ChannelID, "第一個王者!")
	}
	if msg.Content == "!文本" {
		_, _ = ses.ChannelMessageSend(msg.ChannelID, "jkl, ming, ning, dopa, 惹晒, ruo, xiye, mlxg, 白色风车, xiaohu")
	}
}
