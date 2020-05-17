package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type storyArc struct {
	Title   string
	Story   []string
	Options []option
}

type option struct {
	Text string
	Arc  string
}

func main() {
	storyArcs := getStoryData("gopher.json")
	fmt.Println(storyArcs["intro"].Story[2])
}

func getStoryData(filename string) map[string]storyArc {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Store raw data from file
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	storyArcs := make(map[string]storyArc)
	for key, value := range result {
		storyArcObj := value.(map[string]interface{})
		title := storyArcObj["title"].(string)
		story := storyArcObj["story"].([]interface{})

		var lines []string
		for _, line := range story {
			lines = append(lines, line.(string))
		}

		opts := storyArcObj["options"].([]interface{})
		var options []option
		for _, opt := range opts {
			o := opt.(map[string]interface{})
			arc := o["arc"].(string)
			text := o["text"].(string)
			optionStruct := option{
				Text: text,
				Arc:  arc,
			}

			options = append(options, optionStruct)
		}

		storyArcStruct := storyArc{
			Title:   title,
			Story:   lines,
			Options: options,
		}
		storyArcs[key] = storyArcStruct
	}

	return storyArcs
}
