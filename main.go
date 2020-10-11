package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
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
	var namespace string
	args := strings.Fields(command)
	if args[0] != "!k" {
		return
	}
	if args[1] == "get" {
		switch {
		case args[2] == "pods":
			if len(args) == 4 {
				namespace = args[3]
			}
			kubehandler.GetPods(s, m, namespace)
		case args[2] == "ns":
			kubehandler.GetNamespace(s, m)

		case args[2] == "deploy":
			if len(args) == 4 {
				namespace = args[3]
			}
			kubehandler.GetDeploy(s, m, namespace)
		case args[2] == "svc":
			if len(args) == 4 {
				namespace = args[3]
			}
			kubehandler.GetSvc(s, m, namespace)
		case args[2] == "ingress":
			if len(args) == 4 {
				namespace = args[3]
			}
			kubehandler.GetIngress(s, m, namespace)
		case args[2] == "cm":
			if len(args) == 4 {
				namespace = args[3]
			}
			kubehandler.GetConfigMap(s, m, namespace)
		case args[2] == "nodes":
			kubehandler.GetNodes(s, m)
		}
	}

}
