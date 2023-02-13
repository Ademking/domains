package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extract_domain(input string) []string {
	r, err := regexp.Compile(`(?m)([a-z0-9][a-z0-9\-]{0,61}[a-z0-9]\.)+[a-z0-9][a-z0-9\-]*[a-z0-9]`)
	if err != nil {
		return []string{}
	}
	matchs := r.FindAllString(input, -1)
	return matchs

}

func help() {
	fmt.Println(`       __                      _           `)
	fmt.Println(`  ____/ /___  ____ ___  ____ _(_)___  _____`)
	fmt.Println(` / __  / __ \/ __ ` + "`" + `__ \/ __ ` + "`" + `/ / __ \/ ___/`)
	fmt.Println(`/ /_/ / /_/ / / / / / / /_/ / / / / (__  ) `)
	fmt.Println(`\__,_/\____/_/ /_/ /_/\__,_/_/_/ /_/____/  `)
	fmt.Println(`                                           `)
	fmt.Println("Usage: cat domains.txt | domains [-prefix string] [-suffix string]")
	flag.PrintDefaults()
	fmt.Println(`Examples:`)
	fmt.Println(`  cat logs.txt | domains`)
	fmt.Println(`  cat logs.txt | domains --prefix "https://" --suffix "/api"`)
	fmt.Println(`  domains --input logs.txt`)
	fmt.Println()
	os.Exit(0)
}

func format_domain(prefix string, domain string, suffix string) string {
	if prefix != "" {
		domain = prefix + domain
	}
	if suffix != "" {
		domain = domain + suffix
	}
	return domain
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	domains := make(map[string]bool)
	prefixStr := flag.String("prefix", "", "Add a string at the beginning of the domain")
	suffixStr := flag.String("suffix", "", "Add a string at the end of the domain")
	filePath := flag.String("input", "", "Read from a file")
	helpStr := flag.Bool("h", false, "Show this help")
	flag.Parse()

	if *helpStr {
		help()
	}

	if *filePath == "" {
		for scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				continue
			}
			domains_list := extract_domain(input)
			if len(domains_list) == 0 {
				continue
			}
			for _, domain := range domains_list {
				domains[domain] = true
			}
		}

		for domain := range domains {
			domain = format_domain(*prefixStr, domain, *suffixStr)
			os.Stdout.WriteString(domain + "\n")
		}

	} else {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		// Read domains from file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				continue
			}
			domains_list := extract_domain(input)
			if len(domains_list) == 0 {
				continue
			}
			for _, domain := range domains_list {
				domains[domain] = true
			}
		}

		for domain := range domains {
			domain = format_domain(*prefixStr, domain, *suffixStr)
			os.Stdout.WriteString(domain + "\n")
		}
	}
}
