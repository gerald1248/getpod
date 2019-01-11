package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage: %s [-kubeconfig=PATH] PODNAME`,
			filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No search expression given\n")
		os.Exit(0)
	}

	search := args[0]

	names := getPods(kubeconfig)

	for _, name := range names {
		if strings.Contains(name, search) {
			fmt.Printf(name)
		}
	}
}
