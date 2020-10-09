package kubehandler

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	kubeconfig, envErr := os.LookupEnv("KUBE_CONFIG_PATH")
	log.Printf("Error in getting env var: %v", envErr)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		log.Printf("Error in getting k8s credentials: %v", err)
	}

	clientset, err = kubernetes.NewForConfig(config)

	if err != nil {
		log.Printf("Error in authentication: %v", err)
	}

}
