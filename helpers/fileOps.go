package helpers

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func MakeOutputFile(symbol string, outputDir string) os.File {
	sanitizedFilename := regexp.MustCompile(`[!"#$%&':;]`).ReplaceAllString(symbol, "_")

	now := time.Now()
	file, err := os.OpenFile(
		fmt.Sprintf(
			"%s%s_%d_%d_%d.txt",
			outputDir,
			strings.ToLower(sanitizedFilename),
			now.Year(),
			now.Month(),
			now.Day()),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0755)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return *file
}
