package main

import (
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getPods(kubeconfig *string) []string {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// fetch current namespace
	currentNamespace, err := extractCurrentNamespaceFromFile(*kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(currentNamespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []string
	for _, pod := range pods.Items {
		names = append(names, pod.ObjectMeta.Name)
	}

	return names
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
