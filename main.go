package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const jsStirng = `(?:|\[|,|, | ,)\s*['"]([a-zA-Z0-9\_-]+)['"]\s*(?:\s|\]|\)|,|, | ,)`
const bad = `["'\(\):&\[\] ;#]|(^[0-9]+$)`

var patterns = []string{
	`(?:var|const|let)\s+([a-zA-Z_]\w*)\s*=`,  // js varables
	`["']{0,1}([a-zA-Z0-9$_\.-]*?)["']{0,1}:`, // json keys
	`<input (?:[^>]*name=["']([^'"]*)|)`,      // html input name
	`<input (?:[^>]*id=["']([^'"]+)|)`,        // html input id
	`[\?&](?:([^=]+)=)?`,                      // query string key
	`.*\(\s*["|']?([\w\-]+)["|']?\s*` + strings.Repeat(`(\,\s*["|']?([\w\-]+)["|']?\s*)?`, 10) + `\)`, // function inputs
}

func main() {
	// parse CLI arguments
	if len(os.Args) >= 2 {
		if os.Args[1] == "-strings" {
			patterns = append(patterns, jsStirng)
		} else {
			fmt.Println("Usage: pex [-strings]")
			os.Exit(1)
		}
	}

	// Read input
	content, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read stdin: %v\n", err)
		os.Exit(1)
	}

	printed := make(map[string]bool)
	badRe := regexp.MustCompile(bad)

	for _, p := range patterns {
		re := regexp.MustCompile(p)
		matches := re.FindAllStringSubmatch(string(content), -1)

		for _, match := range matches {
			for _, param := range match[1:] {
				if isValidParam(param, printed, badRe) {
					fmt.Println(param)
					printed[param] = true
				}
			}
		}

	}

}

func isValidParam(param string, printed map[string]bool, badRe *regexp.Regexp) bool {
	return len(param) > 0 &&
		len(param) < 35 &&
		!printed[param] &&
		!badRe.MatchString(param)
}
