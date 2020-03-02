package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"strings"
	"syscall"
)

func main() {
	for {
		fmt.Println("Enter password to check")
		pw, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("Checking...")
		if err != nil {
			fmt.Println(err)
		}

		h := sha1.New()
		h.Write(pw)

		hstr := fmt.Sprintf("%X", h.Sum(nil))

		res, err := http.Get("https://api.pwnedpasswords.com/range/" + hstr[:5])
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		scanner := bufio.NewScanner(res.Body)
		found := ""
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ":")
			if s[0] == hstr[5:] {
				found = s[1]
				break
			}
		}

		if found != "" {
			fmt.Println("Found", found, "times")
		} else {
			fmt.Println("Couldn't find password")
		}
	}
}
