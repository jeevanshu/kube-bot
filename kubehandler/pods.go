package kubehandler

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func init() {
	kubeconfig := "/Users/a1570543/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		log.Printf("Error in getting k8s credentials: %v", err)
	}

	clientset, err = kubernetes.NewForConfig(config)

	if err != nil {
		log.Printf("Error in authentication: %v", err)
	}

}
func GetPods(s *discordgo.Session, m *discordgo.MessageCreate) {

	// nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error in reading pods")
	}
	var podlist string
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
		podlist += fmt.Sprintf("%v\n", pod.Name)
	}

	// message := "hi"
	s.ChannelMessageSend(m.ChannelID, podlist)
}
