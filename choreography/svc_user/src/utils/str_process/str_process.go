package str_process

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GenerateUniqueCode params
// @text: string
// return string
func GenerateUniqueCode(text string) string {
	rand.Seed(time.Now().UTC().UnixNano())

	text = strings.ToLower(text)
	var numb [4]int
	var codes []string

	// store random int to array int
	for i := 0; i < len(numb); i++ {
		numb[i] = rand.Intn(10)
	}

	// replace all char except alphabet and number -> '-'
	reg, err := regexp.Compile("[^a-zA-Z0-9]")
	if err != nil {
		log.Fatal(err)
	}

	for i := range numb {
		temp := strconv.Itoa(numb[i])
		codes = append(codes, temp)

	}
	code := strings.Join(codes, "")

	processedText := reg.ReplaceAllString(text, "-")
	lengthText := len(processedText)

	if lengthText > 0 {
		if processedText[lengthText-1] == '-' {
			processedText = strings.TrimRight(processedText, "-")
		}
		if processedText[0] == '-' {
			processedText = strings.TrimLeft(processedText, "-")
		}
	}

	// merge "processed-text" + "-" + "7291"
	result := processedText + "-" + code
	return result

}
