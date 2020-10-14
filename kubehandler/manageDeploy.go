package kubehandler

import (
	"context"
	"log"
	"strconv"
	"time"

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

// DeleteDeployment function to delete deployments
func DeleteDeployment(s *discordgo.Session, m *discordgo.MessageCreate, namespace string, deploy string) {

	var confirmVar bool

	confirmMsg, confErr := s.ChannelMessageSend(m.ChannelID, "Please react with ✅ emoji within 1 min to confirm deletion of "+deploy)
	if confErr != nil {
		log.Printf("Error in sending message: %v", confErr)
	}

	time.Sleep(1 * time.Minute)

	updatedMsg, err := s.ChannelMessage(m.ChannelID, confirmMsg.ID)
	if err != nil {
		log.Printf("Error in getting message: %v", err)
	}
	if len(updatedMsg.Reactions) > 1 {
		confirmVar = false
	} else {
		for _, i := range updatedMsg.Reactions {
			if i.Emoji.Name == "✅" {
				confirmVar = true
			} else {
				confirmVar = false
			}
		}
	}

	if confirmVar == true {
		deletePolicy := metav1.DeletePropagationForeground
		err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deploy, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Printf("Error in deleting deployment : %v", err)
		}

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: deploy + " Deleted",
			},
			Description: "Successfully deleted deployment " + deploy,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: deploy + " Deletion Aborted",
			},
			Description: "Deletetion aborted for deployment " + deploy,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}
func int32Ptr(i string) *int32 {
	s, _ := strconv.Atoi(i)
	k := int32(s)
	return &k
}
