/*
	Channel + 无锁版本
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

const StudentsNum = 10
const BufferNum = 5

var wg sync.WaitGroup
var Scores chan Score = make(chan Score, StudentsNum)
var trafficlimit chan int = make(chan int, BufferNum)

type Score struct {
	name  string
	score int
}

func main() {

	for i := 0; i < StudentsNum; i++ {
		wg.Add(1)
		trafficlimit <- i
		go submitScores(i)
	}
	wg.Wait()
	close(Scores)
	fmt.Printf("学生成绩统计情况：\n")
	for score := range Scores {
		fmt.Printf("%s:%d\n", score.name, score.score)
	}
}

func submitScores(i int) {
	defer wg.Done()
	Scores <- randInfo(i)
	<-trafficlimit
}

func randInfo(i int) Score {
	return Score{
		name:  strconv.Itoa(i) + "_stus",
		score: rand.Intn(100),
	}
}
