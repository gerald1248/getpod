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

	// determine config path first
	var kubeconfig *string
	// if presented as flag, use that
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// failing that, use ${KUBECONFIG}
	if len(*kubeconfig) == 0 {
		*kubeconfig = os.Getenv("KUBECONFIG")
	}
	// as a last resort, use standard location
	if len(*kubeconfig) == 0 {
		*kubeconfig = filepath.Join(homeDir(), ".kube", "config")
	}

	// other flags
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
