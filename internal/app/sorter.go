package app

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Sorter struct {
	arr       []int
	arrLength int
	method    string
	duration  time.Duration
}

func NewSorter() *Sorter {
	var s Sorter
	s.arr = make([]int, 0)
	return &s
}

func (s *Sorter) evaluateLen() {
	s.arrLength = len(s.arr)
}

func (s *Sorter) ShowLen() int {
	s.evaluateLen()
	return s.arrLength
}

func (s *Sorter) RandomFill(l int) {
	//start := time.Now()
	buffer := s.GenerateArray(l)
	for i := 0; i < l; i++ {
		s.arr = append(s.arr, buffer[i])
	}
	s.evaluateLen()
	//duration := time.Since(start)
	//fmt.Println(duration)
}

func (s *Sorter) FullArray(sourceArr []int) {
	for i := 0; i < len(sourceArr); i++ {
		s.arr = append(s.arr, sourceArr[i])
	}
	s.evaluateLen()
}

func (s *Sorter) GenerateArray(l int) (generatedArray []int) {
	salt := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(salt)
	for i := 0; i < l; i++ {
		generatedArray = append(generatedArray, rand1.Intn(l*2))
	}
	return
}

func (s *Sorter) ShowArray() {
	if s.arrLength > 10 {
		step := s.arrLength / 10
		tempArr := make([]int, 0)
		for i := 0; i < s.arrLength; i = i + step {
			tempArr = append(tempArr, s.arr[i])
		}
		fmt.Printf("The array %v (partial visible), was sorted by %v method and consumed %v\n", tempArr, s.method, s.duration)
		return
	}
	fmt.Printf("The array %v, was sorted by %v method and consumed %v\n", s.arr, s.method, s.duration)
}

func (s *Sorter) IsSorted() bool {
	for i := 0; i < s.arrLength-1; i++ {
		if s.arr[i] > s.arr[i+1] {
			return false
		}
	}
	return true
}

func (s *Sorter) BubbleSort(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 0; i < s.arrLength-1; i++ {
		for j := 0; j < s.arrLength-1-i; j++ {
			if s.arr[j] > s.arr[j+1] {
				s.arr[j], s.arr[j+1] = s.arr[j+1], s.arr[j]
			}
		}
	}
	s.method = "Buble"
	s.duration = time.Since(start)

}

func (s *Sorter) InsertSort(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 0; i < s.arrLength; i++ {
		for j := i; j != 0 && s.arr[j] < s.arr[j-1]; j-- {
			s.arr[j-1], s.arr[j] = s.arr[j], s.arr[j-1]
		}
	}
	s.method = "Insert"
	s.duration = time.Since(start)
}

func (s *Sorter) QuickSort(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	recursiveSort(s.arr, 0, s.arrLength-1)

	s.method = "Quick"
	s.duration = time.Since(start)
}

func recursiveSort(s1 []int, begin, end int) {
	if begin >= end {
		return
	}

	pivot := s1[begin]
	i := begin + 1

	for j := begin; j <= end; j++ {
		if pivot > s1[j] {
			s1[i], s1[j] = s1[j], s1[i]
			i++
		}
		//fmt.Printf("Sorting ...:\t%v\n", s1)
	}

	s1[begin], s1[i-1] = s1[i-1], s1[begin]

	recursiveSort(s1, begin, i-2)
	recursiveSort(s1, i, end)
}

func (s *Sorter) BuiltInSort(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	sort.Ints(s.arr)

	s.method = "BuiltIn"
	s.duration = time.Since(start)
}
