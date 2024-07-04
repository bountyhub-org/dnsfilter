package main

import "testing"

func TestToASCII(t *testing.T) {
	tt := map[string]bool{
		"example.com":                true,
		"example.com.":               true,
		".example.com":               true,
		"www.müller.example.com":     true,
		"target.example.com":         false,
		"example.target.example.com": false,
		".":                          false,
		"":                           false,
		".com":                       false,
		"example-.com":               false,
		"-example.com":               false,
		"*.example.com":              false,
		"example.net":                false,
		"invalid_domain.com":         false,
	}

	for domain, ok := range tt {
		_, err := ToASCII(domain, []string{"example.com"}, []string{"target.example.com"})
		if ok && err != nil {
			t.Errorf("ToASCII(%q) = %v; want nil", domain, err)
		}
		if !ok && err == nil {
			t.Errorf("ToASCII(%q) = nil; want error", domain)
		}
	}
}
