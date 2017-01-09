package lynx

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"time"
)

const (
	date_layout = `Jan _2 2006` // ref time: Mon Jan 2 15:04:05 -0700 MST 2006
)

var (
	errNoFrontMatter   = errors.New("no front matter")
	errInvalidYaml     = errors.New("invalid yaml front matter")
	errUnrecognizedKey = errors.New("unrecognized yaml key")
)

func parseDateFrom(raw string) (time.Time, error) {
	return time.Parse(date_layout, raw)
}

func parseFrontMatterIn(b []byte) (map[string]interface{}, error) {
	buf := bytes.NewBuffer(b)
	r := bufio.NewReader(buf)

	stringAnyMap := make(map[string]interface{})
	i, hasFrontMatter := 0, false
	for l, err := r.ReadBytes('\n'); err == nil; l, err = r.ReadBytes('\n') {

		line := string(l)
	Switch:
		switch {

		// Return if first line didn't have frontmatter
		case !hasFrontMatter && i > 0:
			return stringAnyMap, errNoFrontMatter

		// Begin frontmatter parse if first line is `---\n`
		case line == "---\n" && i == 0: // !hasFrontMatter:
			hasFrontMatter = true

		// If we reach another `---\n` consider as closing delimiter
		case line == "---\n":
			return stringAnyMap, nil

		// Parse front matter in current line
		case hasFrontMatter:
			key, value, err := parseFrontMatterLine(line[:len(line)-1])
			if err != nil {
				break Switch
			}
			stringAnyMap[key] = value
		}

		i++
	}

	return stringAnyMap, nil
}

func parseFrontMatterLine(line string) (string, interface{}, error) {
	colon := strings.Index(line, ":")
	if colon < 0 {
		return "", nil, errInvalidYaml
	}

	key := strings.ToLower(line[:colon])
	value := line[colon+1:]

	// Strip first-rune space off value
	if value[0] == ' ' {
		value = value[1:]
	}

	switch key {
	case "date":
		t, err := time.Parse(date_layout, value)
		if err != nil {
			return "", nil, err
		}
		return key, t, nil

	}
	return key, value, errUnrecognizedKey
}

func stripFrontMatterFrom(b []byte) []byte {
	buf := bytes.NewBuffer(b)
	r := bufio.NewReader(buf)

	i, hasFrontMatter := 0, false
Loop:
	for l, err := r.ReadBytes('\n'); err == nil; l, err = r.ReadBytes('\n') {

		line := string(l)
		switch {

		// Return input bytes as-is if first line didn't have frontmatter
		case !hasFrontMatter && i > 0:
			return b

		// Begin frontmatter parse if first line is `---\n`
		case line == "---\n" && i == 0: // !hasFrontMatter:
			hasFrontMatter = true

		// If we reach another `---\n` after first line,
		// consider as closing delimiter
		// and break Reading loop
		case line == "---\n":
			break Loop
		}

		i++
	}

	// Read the rest of the bytes after the front matter
	writebuf := new(bytes.Buffer)
	writebuf.ReadFrom(r)
	return writebuf.Bytes()
}
