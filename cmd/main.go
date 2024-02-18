package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	layoutISO      = "2006-01-02"
	shroveTuesday  = "Fettisdagen"
	cinnaonRollDay = "Kanelbullens dag"
)

type Year struct {
	Chetime   string `json:"cachetid,omitempty"`
	Startdate string `json:"startdatum,omitempty"`
	Enddate   string `json:"slutdatum,omitempty"`
	Days      []Date `json:"dagar"`
}

type Date struct {
	Date        string   `json:"datum"`
	Weekday     string   `json:"veckodag"`
	WorkfreeDay string   `json:"arbetsfri dag"`
	Holiday     string   `json:"helgdag"`
	RedDay      string   `json:"röd dag"`
	Week        string   `json:"vecka"`
	WeekdayNo   string   `json:"dag i vecka"`
	Nameday     []string `json:"namnsdag"`
	Flagday     string   `json:"flaggdag"`
	Specialday  []string `json:"specialdag,omitempty"`
}

// calcShroveTuesday takes a date in form of YYYY-MM-DD and returns the date subracted with 47 days
func calcShroveTuesday(s string) string {
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		log.Fatalln(err)
	}
	a := t.AddDate(0, 0, -47)
	return a.Format(layoutISO)
}

func main() {
	var easterDays []string

	bunRegex := regexp.MustCompile(`20\d\d-10-04`)
	easterDay := regexp.MustCompile("Påskdagen")

	dates := make(map[string]Date)

	files, err := os.ReadDir("../input/")
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		b, err := os.ReadFile("../input/" + f.Name())
		if err != nil {
			log.Fatalln(err)
		}
		var year Year

		err = json.Unmarshal(b, &year)
		if err != nil {
			log.Fatalln(err)
		}
		for _, date := range year.Days {

			if bunRegex.Match([]byte(date.Date)) {
				date.Specialday = append(date.Specialday, cinnaonRollDay)
			}

			if easterDay.Match([]byte(date.Holiday)) {
				easterDays = append(easterDays, date.Date)
			}
			dates[date.Date] = date

		}

	}

	for _, k := range easterDays {
		newDate := calcShroveTuesday(k)
		d := dates[newDate]
		d.Specialday = append(d.Specialday, shroveTuesday)
		dates[newDate] = d
	}

	b, err := json.Marshal(dates)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("../dates.json", b, 0600)
	if err != nil {
		log.Fatalln(err)
	}

}
