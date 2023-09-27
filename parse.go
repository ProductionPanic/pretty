package pretty

import "strings"

var styleShortcuts = map[string]string{
	"r": "reset",
	"b": "bold",
	"d": "dim",
	"i": "italic",
	"u": "underline",
}

func Parse(s string) (string, error) {
	// example of syntax:
	// Hello {red,bold}world{reset}!
	var output string
	var prev string
	var curStyle string
	var inStyle bool
	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		if cur == "{" && prev != "\\" && !inStyle {
			// start of style
			inStyle = true
		} else if cur == "}" && prev != "\\" && inStyle {
			// end of style
			inStyle = false
			curStyle = strings.ToLower(curStyle)
			parts := strings.Split(curStyle, ",")
			parsedParts := make([]string, len(parts))
			for _, part := range parts {
				// check if in stylemap or shortcut
				if val, ok := styleShortcuts[part]; ok {
					parsedParts = append(parsedParts, val)
				} else if _, ok := StyleMap[part]; ok {
					parsedParts = append(parsedParts, part)
				}
			}
      output += "\033[" + strings.Join(parsedParts, ";") + "m"
		} else if inStyle {
			curStyle += cur
		} else {
      output += cur
    }

		prev = string(s[i])
	}

  return output, nil
}
