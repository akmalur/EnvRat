package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/AkmalUr/test1/ratings"
	"github.com/shopspring/decimal"
)

func main() {
	rawInput, rawQuery := readInput()

	input := processRawInput(rawInput)
	query := processRawQuery(rawQuery)

	repo := ratings.NewInMemoryRepository()
	svc := ratings.NewService(repo)

	for _, i := range input {
		svc.SaveData(i)
	}

	for _, q := range query {
		r := svc.GetRating(q)
		fmt.Println(r.String())
	}
}

func readInput() (rawData []string, rawQuery []string) {
	rawData = []string{}
	rawQuery = []string{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Provide input: ")

	buffer := []string{}

	emptyLineRead := false
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" && !emptyLineRead {
			rawData = buffer
			buffer = []string{}
			emptyLineRead = true
		} else if text == "" && emptyLineRead {
			break
		} else {
			buffer = append(buffer, text)
		}
	}
	rawQuery = buffer
	return rawData, rawQuery
}

func processRawInput(rawInput []string) []*ratings.InputData {
	re, err := regexp.Compile(`"(.*)" "(.*)" (.*)`)
	if err != nil {
		log.Fatal(err)
	}

	result := []*ratings.InputData{}
	for _, rid := range rawInput {
		rawInputData := re.FindStringSubmatch(rid)

		rv, err := decimal.NewFromString(rawInputData[3])
		if err != nil {
			log.Fatal(err)
		}
		rv = rv.Round(3)

		data := ratings.NewInputData(rawInputData[1], processLocationData(rawInputData[2]), &rv)
		result = append(result, data)
	}
	return result
}

func processRawQuery(rawQuery []string) []*ratings.InputQuery {
	re, err := regexp.Compile(`"(.*)" "(.*)"`)
	if err != nil {
		log.Fatal(err)
	}

	result := []*ratings.InputQuery{}
	for _, riq := range rawQuery {
		rawInputQuery := re.FindStringSubmatch(riq)

		data := ratings.NewInputQuery(rawInputQuery[1], processLocationData(rawInputQuery[2]))
		result = append(result, data)
	}
	return result
}

func processLocationData(rawLocation string) *ratings.Location {
	raw := strings.Split(rawLocation, "/")
	result := &ratings.Location{}
	switch len(raw) {
	case 1:
		return ratings.NewLocation(raw[0], "", "")
	case 2:
		return ratings.NewLocation(raw[0], raw[1], "")
	case 3:
		return ratings.NewLocation(raw[0], raw[1], raw[2])
	default:
		log.Fatal(fmt.Sprintf("invalid region: %+v", rawLocation))
	}
	return result
}
