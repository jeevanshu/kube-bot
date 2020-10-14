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
	// Token var for getting discord token
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
	} else if args[1] == "logs" {
		kubehandler.GetPodLogs(s, m, args[2], args[3])
	} else if args[1] == "scale" {
		kubehandler.UpdateDeployment(s, m, args[2], args[3], args[4])
	} else if args[1] == "delete" {
		kubehandler.DeleteDeployment(s, m, args[2], args[3])
	} else if args[1] == "help" {
		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: "kube-bot help menu",
			},
			Description: `To use enter **!k** followed by command, object and namespace
			`,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "kube-bot build using client-go and discordgo, source code at https://github.com/jeevanshu/kube-bot",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "get",
					Value: `
					Use this command to read objects **!k get** followed by object
					Available operations are:
					- pods
					- svc
					- deploy
					- ns
					- cm
					- nodes
					- ingress
					example: **!k get deploy test-ns**
					`,
				},
				{
					Name: "logs",
					Value: `
					Use this command to get logs from pod **!k logs <pod-name> <namespace>** 
					example: **!k logs nginx nginx-namespace**
					`,
				},
				{
					Name: "scale",
					Value: `
					To scale up and down number of replicas of deployment **!k scale <namespace> <deployment> <replicas>**
					examoke **!k scale default nginx 3**
					`,
				},
				{
					Name: "delete",
					Value: `
					To delete objects **!k delete <namespace> <deployment>** and give confirmation to message
					examoke **!k delete default nginx**
					`,
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}
