package main

type (
	frame [2]int
	entry []byte

	payload struct {
		Time       string `json:"time"`
		Host       string `json:"host"`
		Source     string `json:"source"`
		SourceType string `json:"sourcetype"`
		Index      string `json:"index"`
		Event      struct {
			Line   string `json:"line"`
			Source string `json:"source"`
			Tag    string `json:"tag"`
		} `json:"event"`
	}
)

func parseFrames(input []byte) []frame {
	startpos := -1
	matches := -1
	frames := []frame{}

	for i, c := range input {
		if c == '{' {
			if startpos < 0 {
				startpos = i
				matches = 1
			} else {
				matches++
			}
		}

		if c == '}' {
			matches--
			if startpos >= 0 && matches == 0 {
				frames = append(frames, frame{startpos, i + 1})
				startpos = -1
				matches = -1
			}
		}
	}

	return frames
}

func parseEntries(input []byte) []entry {
	frames := parseFrames(input)
	result := make([]entry, len(frames))

	for idx, fr := range frames {
		result[idx] = input[fr[0]:fr[1]]
	}

	return result
}
