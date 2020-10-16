package kubehandler

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func sendConfirmMsg(s *discordgo.Session, m *discordgo.MessageCreate, object string) *discordgo.Message {
	confirmMsg, confErr := s.ChannelMessageSend(m.ChannelID, "Please react with ✅ emoji within 1 min to confirm deletion of : "+object)
	if confErr != nil {
		log.Printf("Error in sending message: %v", confErr)
	}
	return confirmMsg
}

func checkConfirmation(s *discordgo.Session, m *discordgo.MessageCreate, confirmMsg *discordgo.Message) bool {
	var confirmVar bool
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

	return confirmVar
}

// DeleteDeployment function to delete deployments
func DeleteDeployment(s *discordgo.Session, m *discordgo.MessageCreate, deploy string, namespace string) {

	confirmMsg := sendConfirmMsg(s, m, deploy)

	time.Sleep(1 * time.Minute)
	confirmVar := checkConfirmation(s, m, confirmMsg)

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

// DeleteSvc function to delete services
func DeleteSvc(s *discordgo.Session, m *discordgo.MessageCreate, svc string, namespace string) {

	confirmMsg := sendConfirmMsg(s, m, svc)
	time.Sleep(1 * time.Minute)

	confirmVar := checkConfirmation(s, m, confirmMsg)

	if confirmVar == true {
		deletePolicy := metav1.DeletePropagationForeground
		err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), svc, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Printf("Error in deleting svc : %v", err)
		}

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: svc + " Deleted",
			},
			Description: "Successfully deleted service " + svc,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: svc + " Deletion Aborted",
			},
			Description: "Deletetion aborted for service " + svc,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}

// DeleteIngress function to delete ingress
func DeleteIngress(s *discordgo.Session, m *discordgo.MessageCreate, ingress string, namespace string) {

	confirmMsg := sendConfirmMsg(s, m, ingress)
	time.Sleep(1 * time.Minute)

	confirmVar := checkConfirmation(s, m, confirmMsg)

	if confirmVar == true {
		deletePolicy := metav1.DeletePropagationForeground
		err := clientset.NetworkingV1beta1().Ingresses(namespace).Delete(context.TODO(), ingress, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Printf("Error in deleting ingress : %v", err)
		}

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: ingress + " Deleted",
			},
			Description: "Successfully deleted ingress " + ingress,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: ingress + " Deletion Aborted",
			},
			Description: "Deletetion aborted for ingress " + ingress,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}

// DeleteNamespace function to delete namespace
func DeleteNamespace(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {

	confirmMsg := sendConfirmMsg(s, m, namespace)
	time.Sleep(1 * time.Minute)

	confirmVar := checkConfirmation(s, m, confirmMsg)

	if confirmVar == true {
		deletePolicy := metav1.DeletePropagationForeground
		err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Printf("Error in deleting namespace : %v", err)
		}

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: namespace + " Deleted",
			},
			Description: "Successfully deleted namespace " + namespace,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: namespace + " Deletion Aborted",
			},
			Description: "Deletetion aborted for namespace " + namespace,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}

// DeleteCm function to delete services
func DeleteCm(s *discordgo.Session, m *discordgo.MessageCreate, cm string, namespace string) {

	confirmMsg := sendConfirmMsg(s, m, cm)
	time.Sleep(1 * time.Minute)

	confirmVar := checkConfirmation(s, m, confirmMsg)

	if confirmVar == true {
		deletePolicy := metav1.DeletePropagationForeground
		err := clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), cm, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Printf("Error in deleting cm : %v", err)
		}

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: cm + " Deleted",
			},
			Description: "Successfully deleted ConfigMap " + cm,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {

		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name: cm + " Deletion Aborted",
			},
			Description: "Deletetion aborted for ConfigMap " + cm,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

}
