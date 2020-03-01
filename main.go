package main

import (
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/testaction/models"
	"io/ioutil"
	"os"
)

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	dat, err := ioutil.ReadFile(eventPath)
	if err != nil {
    fmt.Printf("unable to read event")
    return
	}

	// have the data, deserialise
  var ev models.GollumEventModel
	err = json.Unmarshal(dat, &ev)
	if err != nil {
		fmt.Printf("cannot unmarshal event data")
		return
	}

  for _,page := range ev.Pages {
  	fmt.Printf("page title: %s , pagename: %s\n", page.Title, page.PageName)
  }

	//fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))
}
