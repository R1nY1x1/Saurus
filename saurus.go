package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Translate(text, source, target string) string {
	endpoint := fmt.Sprintf("https://script.google.com/macros/s/%s/exec", os.Getenv("GAS_KEY"))
	values := url.Values{}
	values.Set("text", text)
	values.Set("source", source)
	values.Set("target", target)

	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	response := string(body)

	return response
}

func Thesaurus(text string) string {
	url := fmt.Sprintf("https://wordsapiv1.p.rapidapi.com/words/%s", text)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("WORDSAPI_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	response := string(body)

	return response
}

func main() {
	var (
		doTransInput  bool
		doTransOutput bool
		source        string
		target        string
	)
	flag.BoolVar(&doTransInput, "xin", false, "do Translate Input")
	flag.BoolVar(&doTransOutput, "xout", false, "do Translate Output")
	flag.StringVar(&source, "s", "ja", "Translate souece")
	flag.StringVar(&target, "t", "en", "Translate target")
	flag.Parse()

	var text string
	if doTransInput {
		text = Translate(flag.Arg(0), source, target)
		fmt.Println("Input:", text, "<", flag.Arg(0), ">")
	} else {
		text = flag.Arg(0)
		fmt.Println("Input:", text)
	}

	response := Thesaurus(text)
	var tmp map[string]interface{}
	json.Unmarshal([]byte(response), &tmp)
	results := (tmp["results"].([]interface{}))
	for i := len(results) - 1; i >= 0; i-- {
		result := (results[i].(map[string]interface{}))
		if synonyms, ok := result["synonyms"]; ok {
			fmt.Println(synonyms, result["definition"])
			if doTransOutput {
				text = Translate(result["definition"].(string), "en", "ja")
				fmt.Println("<", text, ">")
			}
		}
	}
}
