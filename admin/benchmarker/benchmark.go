package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var host = "http://127.0.0.1"
var totalScore = 0
var totalResp = map[bool]int{}

type myEvent struct {
	Workload int    `json:"workload"`
	IP       string `json:"ip"`
	Username string `json:"username"`
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest handler
func HandleRequest(ctx context.Context, event myEvent) (string, error) {
	workload := event.Workload
	ip := event.IP
	host = "https://" + ip
	username := event.Username

	createClients(workload * 5)
	startBenchmark(workload)

	return fmt.Sprintf("Done %s", username), nil
}

func startBenchmark(workload int) {
	getInitialize()
	log.Print("期日前投票を開始します")
	validateInitialize()
	log.Print("期日前投票が終了しました")
	log.Print("投票を開始します  Workload: " + strconv.Itoa(workload))
	voteTime := time.Now().Add(45 * time.Second)
	wg1 := new(sync.WaitGroup)
	m1 := new(sync.Mutex)
	for i := 0; i < workload+1; i++ {
		wg1.Add(1)
		if i%5 == 0 {
			go loopInvalidVoteScenario(wg1, m1, voteTime)
		} else {
			go loopVoteScenario(wg1, m1, voteTime)
		}
	}
	wg1.Wait()
	log.Print("投票が終了しました")
	finishTime := time.Now().Add(15 * time.Second)
	wg2 := new(sync.WaitGroup)
	m2 := new(sync.Mutex)
	log.Print("投票者が結果を確認しています")
	for i := 0; i < workload+2; i++ {
		wg2.Add(1)
		if i%4 == 0 || i%4 == 3 {
			go loopIndexScenario(wg2, m2, finishTime)
		} else if i%4 == 1 {
			go loopCandidateScenario(wg2, m2, finishTime)
		} else {
			go loopPoliticalPartyScenario(wg2, m2, finishTime)
		}
	}
	wg2.Wait()
	printScore()
}

func loopInvalidVoteScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		if invalidVoteScenario(m, finishTime) == false {
			break
		}
	}
	defer wg.Done()
}

func loopVoteScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		if voteScenario(m, finishTime) == false {
			break
		}
	}
	defer wg.Done()
}

func loopIndexScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	log.Print("start")
	for {
		log.Print("in for")
		if indexScenario(m, finishTime) == false {
			break
		}
	}
	defer log.Print("done")
	defer wg.Done()
}

func loopCandidateScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		if candidateScenario(m, finishTime) == false {
			break
		}
	}
	defer wg.Done()
}

func loopPoliticalPartyScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		if politicalPartyScenario(m, finishTime) == false {
			break
		}
	}
	defer wg.Done()
}

func printScore() {
	log.Print("投票者の感心がなくなりました")
	log.Print("{\"score\": " + strconv.Itoa(totalScore) + ", \"success\": " + strconv.Itoa(totalResp[true]) + ", \"failure\": " + strconv.Itoa(totalResp[false]) + "}")
}
