package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const baseUrl = "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=activity&site=stackoverflow"

var (
	tagsString string
	question   string
)

type searchResp struct {
	Items []searchItem `json:"items"`
}

type searchItem struct {
	Link string `json:"link"`
}

func init() {
	flag.StringVar(&tagsString, "tags", "", "Specify a comma separated list of tags. For example: \"go,concurrency\"")
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	tags := strings.Split(tagsString, ",")
	question = flag.Args()[0]

	if len(question) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	res := performSearch(question, tags)

	fmt.Println(res)
}

func performSearch(question string, tags []string) string {
	u, _ := url.Parse(baseUrl)
	q := u.Query()
	for _, tag := range tags {
		q.Add("tagged", tag)
	}
	q.Set("q", question)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())

	if err != nil {
		log.Fatal(err)
	}

	respStr, _ := ioutil.ReadAll(resp.Body)

	respStruct := searchResp{}
	json.Unmarshal(respStr, &respStruct)

	if len(respStruct.Items) > 0 {
		return respStruct.Items[0].Link
	} else {
		return "I found nothing :("
	}
}
