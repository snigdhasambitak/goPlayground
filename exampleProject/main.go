package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	kubeConfiPath := filepath.Join(home, ".kube", "config")

	fmt.Println(kubeConfiPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfiPath)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %v", err)
	}

	namespace := "apps"

	podsClient := client.CoreV1().Pods(namespace)

	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

}
