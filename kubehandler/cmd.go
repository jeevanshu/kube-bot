package kubehandler

import (
	"log"
	"os"
	"strconv"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

// Config struct to store env vars
type Config struct {
	InCluster  bool
	Kubeconfig string
}

// New function to read env values
func New() *Config {
	return &Config{
		InCluster:  getBool("IN_CLUSTER", false),
		Kubeconfig: getEnv("KUBE_CONFIG_PATH", ""),
	}
}

func getBool(name string, defaultVal bool) bool {
	s := getEnv(name, "")
	val, err := strconv.ParseBool(s)
	if err != nil {
		return defaultVal
	}
	return val
}

func getEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultVal
}

func init() {
	c := New()
	inCluster := c.InCluster

	if inCluster == true {
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Printf("Error in getting k8s credentials: %v", err)
		}

		clientset, err = kubernetes.NewForConfig(config)

		if err != nil {
			log.Printf("Error in getting k8s credentials: %v", err)
		}
	} else {
		kubeconfig := c.Kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

		if err != nil {
			log.Printf("Error in getting k8s credentials: %v", err)
		}

		clientset, err = kubernetes.NewForConfig(config)

		if err != nil {
			log.Printf("Error in authentication: %v", err)
		}
	}

}
