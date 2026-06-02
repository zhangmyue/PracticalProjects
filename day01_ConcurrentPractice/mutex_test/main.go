/*
Mutex + 切片共享内存版本
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

const STUDENTNUM = 10
const BufferNum = 5

var wg sync.WaitGroup
var Scores = make([]Score, 0)
var mutex sync.Mutex
var trafficlimit chan int = make(chan int, BufferNum)

type Score struct {
	name  string
	score int
}

func main() {
	for i := 0; i < STUDENTNUM; i++ {
		wg.Add(1)
		trafficlimit <- i
		go submitScores(i)
	}
	wg.Wait()
	fmt.Printf("学生成绩统计情况：\n")
	for i := 0; i < len(Scores); i++ {
		fmt.Printf("%s:%d\n", Scores[i].name, Scores[i].score)
	}
}

func submitScores(i int) {
	defer wg.Done()
	mutex.Lock()
	Scores = append(Scores, randInfo(i))
	mutex.Unlock()
	<-trafficlimit
}

func randInfo(i int) Score {
	return Score{
		name:  strconv.Itoa(i) + "_stus",
		score: rand.Intn(100),
	}
}
