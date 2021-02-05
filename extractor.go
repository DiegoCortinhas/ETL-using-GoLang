package etl

import (
	"Transformer/etl/vo"
	"bufio"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"os"
	"regexp"
	"unicode"
)

var employeeRegex = regexp.MustCompile(`\;Empregado:\;{4}([A-Za-z ]*)\;*`)
var workDayRegex = regexp.MustCompile(`([0-9]{2}/[0-9]{2}) - ([a-z]{3})\;{3}([0-9: ]*).*`)

func extract(file * os.File) [] vo.TimeEntry {

	scanner := bufio.NewScanner(file)

	var entries [] vo.TimeEntry

	empregado := ""
	dia := ""
	horas := ""

	for scanner.Scan() {

		line := scanner.Text()

		stripAccents := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

		line, _, _ = transform.String(stripAccents, line)

		var result [] string = matchRegex(line, employeeRegex)

		if result != nil {
			empregado = result[1]
		}

		result = matchRegex(line, workDayRegex)

		if result != nil {
			dia = result[1]
			horas = result[3]

			if horas != "" {
				entry := vo.TimeEntry{Employee: empregado, Date: dia, Hours: horas}
				entries = append(entries, entry)

				logrus.Debugln(entry)
			}
		}
	}
	log.Printf("Foram extraídos: %v\n", len(entries))
	return entries
}

func matchRegex(line string, regex * regexp.Regexp) [] string {

	if regex.MatchString(line) {
		return regex.FindStringSubmatch(line)
	}
	return nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}