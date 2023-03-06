package pool

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
	"wasmnet/utils"
)

var wg sync.WaitGroup // Used to wait for goroutines to finish

type Job struct {
	ID        int
	Focus     string
	ExecFunc  ExecFunc
	ExecMacro interface{}
}

func CreateJob(id int) Job {
	time.Sleep(time.Millisecond * time.Duration(500))

	// Randomized macro selection
	var i int
	archive := make(map[int]string)
	for macro := range MacroLibrary { // for macro, executeFunc := range MacroLibrary
		i += 1
		//fmt.Println("Key:", macro, "=>", "Element:", executeFunc)
		archive[i] = macro
	}

	min := 1
	max := len(MacroLibrary)
	macroSelectionID := rand.Intn(max-min) + min
	var macroSelection string

	switch macroSelectionID {
	case 0:
		macroSelection = archive[macroSelectionID]
		break
	case 1:
		macroSelection = archive[macroSelectionID]
		break
	case 2:
		macroSelection = archive[macroSelectionID]
		break
	case 3:
		macroSelection = archive[macroSelectionID]
		break
	default:
		macroSelection = archive[0]
	}

	executeMacro, err := CallEmbedded(macroSelection, "") // HelloWorld
	if err != nil {
		log.Fatalln("macroIdentity error:", err)
	}
	var execMacro interface{} = executeMacro

	var newJob Job
	if id%2 == 0 {
		newJob = Job{ID: id, Focus: macroSelection, ExecFunc: nil, ExecMacro: execMacro} // tConv = macro 1 (Hello World)

		utils.UpdatePoolOutput(fmt.Sprintf("execMacro: #%d '%s' = %v", newJob.ID, newJob.Focus, newJob.ExecMacro))
	} else {
		newJob = Job{ID: id, Focus: "Winning Numbers", ExecFunc: TicketPool, ExecMacro: nil} // TicketPool = macro 2 (MegaMillions ticket numbers)
		utils.UpdatePoolOutput(fmt.Sprintf("execFunc: #%d '%s' = %v", newJob.ID, newJob.Focus, newJob.ExecFunc()))
	}

	return newJob
}

func InitWorkload(targetPayload int) {
	jobChannel := make(chan Job)

	for i := 1; i <= targetPayload; i++ {
		wg.Add(1)
		go func(id int) {
			for {
				defer wg.Done()

				//var execJob Job
				//var productSignal ProductSignal
				//var targetOutput string
				job := CreateJob(id)

				if job.ExecFunc == nil {
					_, _ = ExecuteMacro(job.ID, job.Focus, nil, job.ExecMacro)
				} else {
					_, _ = ExecuteMacro(job.ID, job.Focus, job.ExecFunc, nil)
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(jobChannel)
	}()
}
