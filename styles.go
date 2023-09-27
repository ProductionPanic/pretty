package pretty

import (
	"fmt"
	"strconv"
)

var StyleMap = map[string]string{
	"reset":      "0",
	"bold":       "1",
	"dim":        "2",
	"italic":     "3",
	"underline":  "4",
	"slowblink":  "5",
	"rapidblink": "6",
	"blink":      "5",
	"invert":     "7",
	"hide":       "8",
	"strike":     "9",
	// foreground (text) colors
	"black":   "30",
	"red":     "31",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"magenta": "35",
	"cyan":    "36",
	"white":   "37",
	"default": "39",
	// background colors
	"bgblack":   "40",
	"bgred":     "41",
	"bggreen":   "42",
	"bgyellow":  "43",
	"bgblue":    "44",
	"bgmagenta": "45",
	"bgcyan":    "46",
	"bgwhite":   "47",
	"bgdefault": "49",
	// bright foreground colors
	"brightblack":   "90",
	"brightred":     "91",
	"brightgreen":   "92",
	"brightyellow":  "93",
	"brightblue":    "94",
	"brightmagenta": "95",
	"brightcyan":    "96",
	"brightwhite":   "97",
	// bright background colors
	"bgbrightblack":   "100",
	"bgbrightred":     "101",
	"bgbrightgreen":   "102",
	"bgbrightyellow":  "103",
	"bgbrightblue":    "104",
	"bgbrightmagenta": "105",
	"bgbrightcyan":    "106",
	"bgbrightwhite":   "107",
}

func RgbCode(r, g, b int) string {
	return "38;2;" + fmt.Sprint(r) + ";" + fmt.Sprint(g) + ";" + fmt.Sprint(b)
}

func HexCode(hex string) (string,error) {
	// convert hex to rgb
	// Remove the '#' prefix if present
	if hex[0] == '#' {
		hex = hex[1:]
	}

  onErr:= fmt.Errorf("invalid hex code")
	// Parse the hexadecimal string to integer values
	r, err := strconv.ParseInt(hex[0:2], 16, 0)

	if err != nil {
		return "", onErr
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return "", onErr
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return "", onErr
	}

	return RgbCode(int(r), int(g), int(b)), nil
}
