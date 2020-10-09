package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jeevanshu/kube-bot/kubehandler"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error in creating Discord session", err)
		return
	}

	dg.AddHandler(commandHandler)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection ", err)
		return
	}

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	user := m.Author
	channel := m.ChannelID
	command := m.Content
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Printf("Command %v from user %v in channel %v \n", command, user, channel)
	switch {
	case command == "!k get pods":
		kubehandler.GetPods(s, m)
	case command == "!k get ns":
		kubehandler.GetNamespace(s, m)
	case command == "!k get deploy":
		kubehandler.GetDeploy(s, m)
	case command == "!k get svc":
		kubehandler.GetSvc(s, m)
	case command == "!k get ingress":
		kubehandler.GetIngress(s, m)
	case command == "!k get cm":
		kubehandler.GetConfigMap(s, m)
	case command == "!k get nodes":
		kubehandler.GetNodes(s, m)

	}

}
