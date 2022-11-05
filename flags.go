package main

import (
	"flag"
	"strings"
)

var (
	urls            Urls
	wordlist        string
	cores           int
	enumerationtype int
	success         Success
)

type Urls []string

type Success []string

func (u *Urls) Set(value string) error {
	*u = append(*u, value)
	return nil
}

func (u *Urls) String() string {
	return strings.Join(*u, " ")
}

func (s *Success) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *Success) String() string {
	return strings.Join(*s, " ")
}

func init() {
	flag.Var(&urls, "u", "Target URL to be bruteforced [required]")
	flag.Var(&success, "s", "Status code for success, default 200")
	flag.StringVar(&wordlist, "w", "", "Path to wordlist [required]")
	flag.IntVar(&cores, "c", 1, "Number of cores, default 1")
	flag.IntVar(&enumerationtype, "t", 0, "Type of enumeration, default 0") // 0 -> sub-domain, 1 -> directory
}
