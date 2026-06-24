package adventure_time

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func StoryHandler(writer http.ResponseWriter, req *http.Request) {
	arc := strings.TrimPrefix(req.URL.Path, "/")
	fullData := ReadStory()

	if arc == "" {
		arc = "intro"
	}

	chapter := fullData[arc]

	pageData := PageData{
		PageTitle:   chapter.Title,
		PageBody:    chapter.Story,
		PageOptions: chapter.Options,
	}
	template, err := template.ParseFiles("./adventure_time/storytime.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	template.Execute(writer, pageData)
}

func ReadStory() FullStory {
	data, err := os.ReadFile("./adventure_time/story.json")
	if err != nil {
		panic(err)
	}
	var story StoryStruct
	storymap := make(map[string]Chapter)

	err = json.Unmarshal(data, &story)
	if err != nil {
		panic(err)
	}

	storymap["intro"] = story.Intro
	storymap["new-york"] = story.NewYork
	storymap["debate"] = story.Debate
	storymap["denver"] = story.Denver
	storymap["sean-kelly"] = story.SeanKelly
	storymap["mark-bates"] = story.MarkBates
	storymap["home"] = story.Home

	return storymap
}

type FullStory map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type PageData struct {
	PageTitle   string
	PageBody    []string
	PageOptions []Option
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type StoryStruct struct {
	Intro     Chapter `json:"intro"`
	NewYork   Chapter `json:"new-york"`
	Debate    Chapter `json:"debate"`
	SeanKelly Chapter `json:"sean-kelly"`
	MarkBates Chapter `json:"mark-bates"`
	Denver    Chapter `json:"denver"`
	Home      Chapter `json:"home"`
}
