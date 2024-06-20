package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cut(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a string")
	}

	str := os.Args[1]
	strs := strings.Split(str, "\n")

	var (
		f int
		d string
		s bool
	)

	fs := flag.NewFlagSet(str, flag.ContinueOnError)
	fs.IntVar(&f, "f", 0, "\"fields\" select column number")
	fs.StringVar(&d, "d", "	", "\"delimiter\" use another delimiter")
	fs.BoolVar(&s, "s", false, "\"separated\" strings only with delimiter")
	fs.Parse(os.Args[2:])

	if f == 0 {
		fmt.Println("want string, flag")
		return nil
	}

	res := Fields(strs, f, d, s)

	for _, s := range res {
		fmt.Println(s)
	}

	return nil
}

func Fields(strs []string, f int, d string, s bool) []string {
	res := make([]string, 0, len(strs))

	for _, str := range strs {
		words := strings.Split(str, d)
		if s {
			if len(words) > 1 {
				for j := range words {
					if j+1 == f {
						res = append(res, words[j])
					}
				}
			}
		} else {
			if len(words) > 1 {
				for j := range words {
					if j+1 == f {
						res = append(res, words[j])
						continue
					}
				}
			} else {
				res = append(res, words...)
			}
		}
	}

	return res
}

func main() {
	if err := cut(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
