package lynx

import (
	"errors"
	"time"
	"bufio"
	"bytes"
	"strings"
)

const (
	date_layout = `Jan _2 2006` // ref time: Mon Jan 2 15:04:05 -0700 MST 2006
)

var (
	errNoFrontMatter = errors.New("no front matter")
	errInvalidYaml = errors.New("invalid yaml front matter")
	errUnrecognizedKey = errors.New("unrecognized yaml key")
)

func parseDateFrom(raw string) (time.Time, error) {
	return time.Parse(date_layout, raw)
}

func parseFrontMatterIn(b []byte) (map[string]interface{}, error) {
	buf := bytes.NewBuffer(b)
	r := bufio.NewReader(buf)

	m := make(map[string]interface{})
	i, hasFrontMatter := 0, false
	for l, err := r.ReadBytes('\n'); err == nil; l, err = r.ReadBytes('\n') {
		
		line := string(l)
		Switch:
		switch {

		// Return if first line didn't have frontmatter
		case !hasFrontMatter && i > 0:
			return m, errNoFrontMatter
		
		// Begin frontmatter parse if first line is `---\n`
		case line == "---\n" && i == 0: // !hasFrontMatter:
			hasFrontMatter = true

		// If we reach another `---\n` consider as closing delimiter
		case line == "---\n":
			return m, nil

		// Parse front matter in current line
		case hasFrontMatter:
			k, v, err := parseFrontMatterLine(line[:len(line)-1])
			if err != nil {
				break Switch
			}
			m[k] = v
		}

		i++
	}

	return m, nil
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