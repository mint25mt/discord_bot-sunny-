package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Token      = os.Getenv("DISCORD_BOT_TOKEN")
	BotName    = "874558469708075019"
	stopBot    = make(chan bool)
	vcsession  *discordgo.VoiceConnection
	HelloWorld = "!helloworld"
)

func main() {
	//Discordのセッションを作成
	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	fmt.Println("Listening...")
	<-stopBot //プログラムが終了しないようロック
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	switch m.Content {
	case "!HelloWorld":
		s.ChannelMessageSend(m.ChannelID, "nyanyanyanya!!!")
	}
}

//メッセージを送信する関数
func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
