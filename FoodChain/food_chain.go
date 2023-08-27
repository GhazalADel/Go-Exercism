package foodchain

import (
	"fmt"
	"strings"
)

const (
	endVerse = "I don't know why she swallowed the fly. Perhaps she'll die."
	filler   = "She swallowed the %s to catch the %s."
	special  = "She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her."
	start    = "I know an old lady who swallowed a %s."
)

var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow"}

var follower = []string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
}

func Song() string {
	return Verses(1, 8)
}

func Verse(line int) string {
	line -= 1
	if line == 7 {
		return `I know an old lady who swallowed a horse.
She's dead, of course!`
	}
	verse := fmt.Sprintf(start, animals[line])
	if line > 0 {
		verse = fmt.Sprintf("%s\n%s", verse, follower[line])
	}
	for i := line; i > 0; i-- {
		if i == 2 {
			verse = fmt.Sprintf("%s\n%s", verse, fmt.Sprintf(special, animals[i], animals[i-1]))
		} else {
			verse = fmt.Sprintf("%s\n%s", verse, fmt.Sprintf(filler, animals[i], animals[i-1]))
		}

	}
	verse = fmt.Sprintf("%s\n%s", verse, endVerse)
	return verse
}

func Verses(start, end int) string {
	out := []string{}
	for i := start; i <= end; i++ {
		out = append(out, Verse(i))
	}
	return strings.Join(out, "\n\n")
}
