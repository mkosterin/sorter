package main

import (
	"log"
	"math/rand"
	"sorter/internal/app"
	"sync"
	"time"
)

const (
	ArrayLength = 100000
	sortMethods = 4
)

func main() {
	log.Printf("The program has been started")
	var wg sync.WaitGroup
	sorters := [sortMethods]*app.Sorter{}
	arrayToSort := generateArray(ArrayLength)
	//fmt.Printf("Initial array %v\n", arrayToSort)
	for i := 0; i < sortMethods; i++ {
		sorters[i] = app.NewSorter()
		sorters[i].FullArray(arrayToSort)
	}
	wg.Add(sortMethods)
	go sorters[0].BubbleSort(&wg)
	go sorters[1].InsertSort(&wg)
	go sorters[2].QuickSort(&wg)
	go sorters[3].BuiltInSort(&wg)
	wg.Wait()
	for i := 0; i < sortMethods; i++ {
		sorters[i].ShowArray()
	}

	log.Printf("The program has been finished")
}

func generateArray(l int) (generatedArray []int) {
	salt := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(salt)
	for i := 0; i < l; i++ {
		generatedArray = append(generatedArray, rand1.Intn(l))
	}
	return
}
