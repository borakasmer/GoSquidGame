package main

import (
	"fmt"
	"log"
	"squidGameGo/coreLib"
	"squidGameGo/model"
	"squidGameGo/mongoDB"
	"sync"
	"time"
)

func main() {
	defer calculationTime(time.Now(), "Get All Squid Players")
	squidArray := make([]model.SquidPlayer, 0)
	isContinue := true
	personChannel := make(chan []model.SquidPlayer, 0)

	csvReader, f, noList := coreLib.CreatReader()
	defer f.Close()
	isFirstRow := true
	var wg sync.WaitGroup
	for isContinue {
		wg.Add(1)
		dataPerson, lastRow := coreLib.ReadSquidList(csvReader, isFirstRow)
		isFirstRow = false
		if len(dataPerson) > 0 {
			go PrintPerPlayer(personChannel, dataPerson, noList, lastRow, &wg)
		} else {
			isContinue = false
		}
		if lastRow {
			isContinue = false
		}
		wg.Wait()
	}

	mongoClient, ctx, err := mongoDB.MongoOpen()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(*ctx)

	for persons := range personChannel {
		//squidArray = append(squidArray, person)
		//fmt.Printf("Nick: %s No:%d\n", person.Nick, person.No)
		//time.Sleep(300 * time.Millisecond)
		for _, person := range persons {
			squidArray = append(squidArray, person)
			//fmt.Printf("Nick: %s No:%d\n", person.Nick, person.No)
			//time.Sleep(300 * time.Millisecond)
		}
	}
	coreLib.SortArray(squidArray)

	for _, person := range squidArray {
		fmt.Printf("Nick: %s No:%d\n", person.Nick, person.No)
	}

	for _, partArray := range coreLib.ParseArray(squidArray) {
		fmt.Printf("%v\n", partArray)
		err = mongoDB.BulkInsert(mongoClient, ctx, partArray)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func PrintPerPlayer(personChannel chan []model.SquidPlayer, personList [][]string, noList *[]int, lastRow bool, wg *sync.WaitGroup) {
	if lastRow {
		defer close(personChannel)
	}
	modelList := make([]model.SquidPlayer, 0)
	for _, person := range personList {
		modelUser := model.SquidPlayer{
			Nick: coreLib.CreateNick(person),
			No:   coreLib.GenerateNo(1, 456, noList),
		}
		modelList = append(modelList, modelUser)
	}
	wg.Done()
	personChannel <- modelList
}

func calculationTime(start time.Time, description string) {
	elapsed := time.Since(start)
	fmt.Printf("%s Totla Time:%s", description, elapsed)
}
