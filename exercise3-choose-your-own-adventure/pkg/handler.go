package adventure

type AdventureStruct struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type AdventureArray map[string]AdventureStruct

type AdventureHandler struct {
	Content AdventureArray
}
