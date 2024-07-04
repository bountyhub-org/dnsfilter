package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

var ErrOutOfScope = errors.New("domain is out of scope")

type stringSlice []string

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func main() {
	var (
		scope      stringSlice
		outOfScope stringSlice
		silent     bool
	)

	flag.Var(&scope, "scope", "Suffix for domain in scope")
	flag.Var(&outOfScope, "out-of-scope", "Suffix for domain out of scope")
	flag.BoolVar(&silent, "silent", false, "Silent errors (do not print to stderr)")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		ascii, err := ToASCII(domain, scope, outOfScope)
		if err == nil {
			fmt.Println(ascii)
			continue
		}

		if !silent {
			fmt.Fprintf(os.Stderr, "Error converting %q: %v\n", domain, err)
		}
	}
}

func ToASCII(domain string, scope, outOfScope []string) (string, error) {
	if domain == "" {
		return "", errors.New("domain is empty")
	}
	if domain == "." {
		return "", errors.New("domain is root")
	}
	if domain[len(domain)-1] == '.' {
		domain = domain[:len(domain)-1]
	}

	result, err := idna.Lookup.ToASCII(domain)
	if err != nil {
		return "", err
	}

	for _, suffix := range outOfScope {
		if strings.HasSuffix(domain, suffix) {
			return "", ErrOutOfScope
		}
	}

	if len(scope) == 0 {
		return result, nil
	}

	for _, suffix := range scope {
		if strings.HasSuffix(domain, suffix) {
			return result, nil
		}
	}

	return "", ErrOutOfScope
}
