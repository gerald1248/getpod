package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage: %s [-kubeconfig=PATH] [-a] [-n NAMESPACE] REGEX\n",
			filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

	// flags
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	all := flag.Bool("a", false, "return all matching pods")
	namespace := flag.String("n", "", "namespace")
	flag.Parse()

	// params
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No search expression given\n")
		flag.Usage()
		os.Exit(1)
	}

	search := args[0]
	re, err := regexp.Compile(search)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regular expression '%s'\n", search)
		os.Exit(2)
	}

	names := getPods(*kubeconfig, *namespace)

	for i, name := range names {
		if re.MatchString(name) {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf(name)
			if *all == false {
				break
			}
		}
	}
}
