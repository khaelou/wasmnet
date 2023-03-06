package pool

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

var ProductChannel = make(chan ProductSignal, math.MaxInt8)

type ExecFunc func() interface{}

type ProductSignal struct {
	Product interface{}
}

func checkProductQuality(job Job, product ProductSignal) {
	fmt.Println(fmt.Sprintf("[✓] Job #%d @ %s = %v", job.ID, job.Focus, product))
}

// Execute target built-in macro specified in MacroLibrary
func CallEmbedded(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(MacroLibrary[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is out of index")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value = f.Call(in)
	result = res[0].Interface()
	return
}

// Exectute externally imported macros
func ExecuteMacro(id int, focus string, execFunc ExecFunc, execMacro interface{}) (Job, ProductSignal) {
	var product interface{}

	if execFunc != nil {
		execFunc()

		product = execFunc()
	} else {
		product = execMacro
	}

	productSignal := ProductSignal{Product: product}
	ProductChannel <- productSignal

	job := Job{ID: id, Focus: focus, ExecFunc: execFunc}
	return job, productSignal
}

func TicketPool() interface{} {
	var genAmount = 6
	numberSlice := []string{}

	fmt.Println("---- TICKET ----")
	for i := 1; i <= genAmount; i++ {
		if i != genAmount {
			normalNumber := fmt.Sprintf("#%d", genNumber1To70())
			numberSlice = append(numberSlice, normalNumber)
		} else {
			goldNumber := fmt.Sprintf("#%d", genNumber1To25())
			numberSlice = append(numberSlice, goldNumber)
		}
	}

	return filterDuplicates(numberSlice)
}

func filterDuplicates(numbers []string) []string {
	inResult := make(map[string]bool)
	var result []string

	for i, str := range numbers {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)

			//fmt.Println("POS_TEST", result)

			realI := i + 1
			if realI < len(numbers) {
				fmt.Println(">", str)
			} else {
				fmt.Println(">", str, "[GOLD]")
			}
		} else {
			fillIn := fmt.Sprintf("#%d", genNumber1To70())
			fillInGold := fmt.Sprintf("#%d", genNumber1To25())

			result = append(result, fillIn)

			//fmt.Println("NEG_TEST", result)

			realI := i + 1
			if realI < len(numbers) {
				fmt.Println(">", fillIn)
			} else {
				fmt.Println(">", fillInGold, "[GOLD]")
			}
		}
	}
	fmt.Println()

	return result
}

func genNumber1To70() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 70 // Any # between 1-70
	number := rand.Intn(max-min+1) + min

	return number
}

func genNumber1To25() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 25 // Any # between 1-25
	number := rand.Intn(max-min+1) + min

	return number
}
