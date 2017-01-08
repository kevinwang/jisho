package client

type Word struct {
	IsCommon      bool     `json:"is_common"`
	Tags          []string `json:"tags"`
	JapaneseForms []Form   `json:"japanese"`
	Senses        []Sense  `json:"senses"`
}

type Form struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

type Sense struct {
	EnglishDefinitions []string `json:"english_definitions"`
	PartsOfSpeech      []string `json:"parts_of_speech"`
	Links              []Link   `json:"links"`
	Tags               []string `json:"tags"`
	Restrictions       []string `json:"restrictions"`
	SeeAlso            []string `json:"see_also"`
	Antonyms           []string `json:"antonyms"`
	Source             []Source `json:"source"`
	Info               []string `json:"info"`
}

type Link struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type Source struct {
	Language string `json:"language"`
	Word     string `json:"word"`
}
