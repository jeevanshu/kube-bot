package kubehandler

import (
	"context"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UpdateDeployment function to scale replicas
func UpdateDeployment(s *discordgo.Session, m *discordgo.MessageCreate, namespace string, deploy string, replicas string) {
	deployCurr, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploy, metav1.GetOptions{})
	if err != nil {
		log.Printf("Error in reading deployment: %v", err)
	}
	deployCurr.Spec.Replicas = int32Ptr(replicas)

	_, scaleErr := clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployCurr, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("Error in scaling deployment: %v", scaleErr)
	}
	msg := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: deploy + " Updated",
		},
		Description: "Scaled deployment " + deploy + " to: " + replicas,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, msg)
}

func int32Ptr(i string) *int32 {
	s, _ := strconv.Atoi(i)
	k := int32(s)
	return &k
}
