package kubehandler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/bwmarrin/discordgo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods function to get pods
func GetPods(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in reading pods: %v", err)
	}
	var podList string
	for _, pod := range pods.Items {
		podList += fmt.Sprintf("%v\n", pod.Name)
	}
	s.ChannelMessageSend(m.ChannelID, podList)
}

// GetNamespace function to get namespaces
func GetNamespace(s *discordgo.Session, m *discordgo.MessageCreate) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting namespaces: %v", err)
	}
	var nsList string
	for _, namespace := range namespaces.Items {
		nsList += fmt.Sprintf("%v\n", namespace.Name)
	}
	s.ChannelMessageSend(m.ChannelID, nsList)

}

// GetDeploy function to get deployments
func GetDeploy(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {

	deploy, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting deployments: %v", err)
	}
	var dsList string
	for _, d := range deploy.Items {
		dsList += fmt.Sprintf("%v\n", d.Name)
	}
	s.ChannelMessageSend(m.ChannelID, dsList)
}

// GetSvc function to get Services
func GetSvc(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {
	service, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting service: %v", err)
	}
	var svcList string
	for _, svc := range service.Items {
		svcList += fmt.Sprintf("%v\n", svc.Name)
	}
	s.ChannelMessageSend(m.ChannelID, svcList)
}

// GetIngress function to get Ingress
func GetIngress(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {
	ingress, err := clientset.NetworkingV1beta1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting ingress: %v", err)
	}
	var ingressList string
	for _, ing := range ingress.Items {
		ingressList += fmt.Sprintf("%v\n", ing.Name)
	}
	s.ChannelMessageSend(m.ChannelID, ingressList)
}

// GetNodes function to get Nodes
func GetNodes(s *discordgo.Session, m *discordgo.MessageCreate) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting ingress: %v", err)
	}
	var nodesList string
	for _, node := range nodes.Items {
		nodesList += fmt.Sprintf("%v\n", node.Name)
	}
	s.ChannelMessageSend(m.ChannelID, nodesList)

}

// GetConfigMap function to get configmap
func GetConfigMap(s *discordgo.Session, m *discordgo.MessageCreate, namespace string) {
	configMaps, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting ingress: %v", err)
	}
	var cmList string
	for _, configMap := range configMaps.Items {
		cmList += fmt.Sprintf("%v\n", configMap.Name)
	}
	s.ChannelMessageSend(m.ChannelID, cmList)
}

// GetPodLogs function to get logs from pods
func GetPodLogs(s *discordgo.Session, m *discordgo.MessageCreate, podName string, namespace string) {
	pods := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{})

	podLogs, err := pods.Stream(context.TODO())
	if err != nil {
		log.Printf("Error in reading pods: %v", err)
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		log.Printf("Error in copy information from podLogs to buf")
	}
	str := buf.String()
	msg := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: podName + " logs",
		},
		Description: str[len(str)-2000:],
	}
	s.ChannelMessageSendEmbed(m.ChannelID, msg)
}
