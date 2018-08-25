package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var host = "http://127.0.0.1"
var username = ""
var totalScore = 0
var totalResp = map[bool]int{}

type myEvent struct {
	Workload int    `json:"workload"`
	IP       string `json:"ip"`
	Username string `json:"username"`
}

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: ./benchmark [option]
Options:
  --workload	N	run benchmark with N workloads (default: 3)
  --ip	IP	specify target IP Address (default: 127.0.0.1)
	--debug		debug mode (DO NOT USE)`)
	}

	var (
		workload = flag.Int("workload", 3, "")
		ip       = flag.String("ip", "127.0.0.1", "")
		name     = flag.String("name", "", "")
		debug    = flag.Bool("debug", false, "")
	)
	flag.Parse()
	username = *name
	host = "https://" + *ip
	if *debug {
		host = "http://127.0.0.1:8080"
	}

	createClients(*workload * 5)
	startBenchmark(*workload)
	db.Close()
}

func startBenchmark(workload int) {
	flushMessage()
	getInitialize()
	postMessage("期日前投票を開始します")
	validateInitialize()
	postMessage("期日前投票が終了しました")
	postMessage("投票を開始します  Workload: " + strconv.Itoa(workload))
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
	postMessage("投票が終了しました")
	finishTime := time.Now().Add(15 * time.Second)
	wg2 := new(sync.WaitGroup)
	m2 := new(sync.Mutex)
	postMessage("投票者が結果を確認しています")
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
	defer wg.Done()
	for {
		if invalidVoteScenario(m, finishTime) == false {
			break
		}
	}
}

func loopVoteScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	defer wg.Done()
	for {
		if voteScenario(m, finishTime) == false {
			break
		}
	}
}

func loopIndexScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	defer wg.Done()
	for {
		if indexScenario(m, finishTime) == false {
			break
		}
	}
}

func loopCandidateScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	defer wg.Done()
	for {
		if candidateScenario(m, finishTime) == false {
			break
		}
	}
}

func loopPoliticalPartyScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	defer wg.Done()
	for {
		if politicalPartyScenario(m, finishTime) == false {
			break
		}
	}
}

func printScore() {
	postMessage("投票者の感心がなくなりました")
	postMessage("score: " + strconv.Itoa(totalScore) + ", success: " + strconv.Itoa(totalResp[true]) + ", failure: " + strconv.Itoa(totalResp[false]))
	postResult(totalScore, totalResp[true], totalResp[false])
}
