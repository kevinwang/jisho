package client

import (
	"fmt"
	"strings"
)

func (s *Sense) NotesString() string {
	var notes []string

	if len(s.Tags) != 0 {
		notes = append(notes, strings.Join(s.Tags, ", "))
	}

	if len(s.Restrictions) != 0 {
		notes = append(notes, strings.Join(s.Restrictions, ", "))
	}

	if len(s.SeeAlso) != 0 {
		seeAlso := make([]string, len(s.SeeAlso))
		for i, w := range s.SeeAlso {
			seeAlso[i] = "See also " + w
		}
		notes = append(notes, strings.Join(seeAlso, ", "))
	}

	if len(s.Antonyms) != 0 {
		antonyms := make([]string, len(s.Antonyms))
		for i, w := range s.Antonyms {
			antonyms[i] = "Antonym: " + w
		}
		notes = append(notes, strings.Join(antonyms, ", "))
	}

	if len(s.Info) != 0 {
		notes = append(notes, strings.Join(s.Info, ", "))
	}

	if len(s.Source) != 0 {
		sources := make([]string, len(s.Source))
		for i, s := range s.Source {
			sources[i] = fmt.Sprintf("From %s %s", s.Language, s.Word)
		}
		notes = append(notes, strings.Join(sources, ", "))
	}

	return strings.Join(notes, ", ")
}
