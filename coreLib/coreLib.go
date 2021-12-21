package coreLib

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"sort"
	"squidGameGo/model"
	"squidGameGo/shared"
	"time"
)

func ParseArray(array []model.SquidPlayer) [][]model.SquidPlayer {
	var rowCount = shared.Config.ROWCOUNT
	parseArray := make([][]model.SquidPlayer, 0)
	var mod = len(array) % rowCount
	for i := 0; i < len(array)-rowCount; i = i + rowCount {
		var arrayPart = array[i : i+rowCount]
		parseArray = append(parseArray, arrayPart)
	}
	var last = 0
	if last = len(array) - rowCount; mod == 0 {

	} else {
		last = len(array) - mod
	}
	var arrayPart = array[last:len(array)]
	parseArray = append(parseArray, arrayPart)

	return parseArray
}

//3,5,9 => 5,9
//5
//11

func Find(what int, where []int) (idx int) {
	if len(where) == 0 {
		return -1
	}
	if what == where[0] {
		return 0
	}
	if idx = Find(what, where[1:]); idx < 0 {
		return -1
	}
	return idx + 1
}

func SortArray(array []model.SquidPlayer) []model.SquidPlayer {
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].No < array[j].No
	})
	return array
}

//3 - 10
func GenerateNo(min int, max int, noList *[]int) int {
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(max-min) + min
		if Find(num, *noList) == -1 {
			//Mutex Lock
			*noList = append(*noList, num)
			//Mutex Unlocal
			return num
		}
	}
	return -1
}

func CreateNick(person []string) string {
	return person[0][0:3] + person[1][0:3]
}

func CreatReader() (*csv.Reader, *os.File, *[]int) {
	fileCsv := shared.Config.FILEURL
	noList := make([]int, 0)
	f, err := os.Open(fileCsv)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = shared.Config.SPERATOR

	return csvReader, f, &noList
}

func ReadSquidList(csvReader *csv.Reader, isFirsRow bool) ([][]string, bool) {
	/*
		fileCsv := shared.Config.FILEURL
		noList := make([]int, 0)
		personData := make([][]string, 0)
		f, err := os.Open(fileCsv)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		csvReader := csv.NewReader(f)
		csvReader.Comma = shared.Config.SPERATOR
	*/
	personData := make([][]string, 0)
	lastRow := false
	dataPerson := make([][]string, 0)
	for i := 0; i < shared.Config.ROWCOUNT; i++ {
		person, err := csvReader.Read()
		if person != nil {
			if err != nil {
				lastRow = true
				log.Fatal(err)
			}
			dataPerson = append(dataPerson, person)
		} else {
			lastRow = true
			break
		}
	}
	for i, line := range dataPerson {

		if (shared.Config.ISHEADER && i > 0 && isFirsRow) || !shared.Config.ISHEADER || !isFirsRow {
			personData = append(personData, []string{line[1], line[2]})
		}
	}
	return personData, lastRow
}
