package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/kevinwang/jisho/client"
)

var numResults int

func init() {
	flag.IntVar(&numResults, "n", 5, "number of results to display")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-n results] query\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 || numResults < 1 {
		flag.Usage()
		os.Exit(1)
	}

	query := strings.Join(flag.Args(), " ")
	words, err := client.Search(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if len(words) == 0 {
		fmt.Printf("Sorry, couldn't find anything matching %s.\n", query)
		os.Exit(3)
	}

	if len(words) <= numResults {
		printWords(words)
	} else {
		printWords(words[:numResults])
	}
}

func printWords(words []client.Word) {
	for i, word := range words {
		if i != 0 {
			fmt.Println()
		}
		printWord(word)
	}
}

func printWord(word client.Word) {
	title := color.New(color.FgHiGreen).SprintFunc()(word.JapaneseForms[0].String())

	var commonWord string
	if word.IsCommon {
		commonWord = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()(" (common word)")
	}

	fmt.Printf("%s%s\n", title, commonWord)

	for i, sense := range word.Senses {
		if len(sense.PartsOfSpeech) != 0 {
			partsOfSpeech := strings.Join(sense.PartsOfSpeech, ", ")
			if i != 0 {
				fmt.Println()
			}
			color.New(color.FgHiBlue).Println("  " + partsOfSpeech)
		}

		definitions := strings.Join(sense.EnglishDefinitions, "; ")
		notes := color.New(color.FgHiBlue).SprintFunc()(sense.NotesString())

		fmt.Printf("  %d. %s %s\n", i+1, definitions, notes)
	}

	if len(word.JapaneseForms) > 1 {
		color.New(color.FgHiBlue).Println("\n  Other forms")
		forms := make([]string, len(word.JapaneseForms[1:]))
		for i, form := range word.JapaneseForms[1:] {
			forms[i] = form.String()
		}
		fmt.Println("  " + strings.Join(forms, "、"))
	}
}
