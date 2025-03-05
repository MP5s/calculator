package rpn

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	Arg1Type   = float64
	Arg2Type   = float64
	ResultType = float64
)

type ExprResultType = float64

type Task struct {
	Arg1          Arg1Type      `json:"arg1"`
	Arg2          Arg2Type      `json:"arg2"`
	Operation     string        `json:"operation"`
	OperationTime int           `json:"operation_time"`
	Status        string        `json:"-"`
	Result        ResultType    `json:"-"`
	Done          chan struct{} `json:"-"`
}

type TaskMap = map[TaskID]*Task

type ConcurrentTaskMap struct {
	m  TaskMap
	mx sync.Mutex
}

func NewConcurrentTaskMap() *ConcurrentTaskMap {
	return &ConcurrentTaskMap{make(map[TaskID]*Task), sync.Mutex{}}
}

func (cm *ConcurrentTaskMap) Get(id TaskID) *Task {
	cm.mx.Lock()
	res, ok := cm.m[id]
	if !ok {
		t := &Task{}
		cm.m[id] = t
		cm.mx.Unlock()
		return t
	}
	cm.mx.Unlock()
	return res
}

func (cm *ConcurrentTaskMap) Add(id TaskID, t *Task) {
	cm.mx.Lock()
	cm.m[id] = t
	cm.mx.Unlock()
}

func (cm *ConcurrentTaskMap) Map() *map[TaskID]*Task {
	return &cm.m
}

type TaskID struct {
	ID TaskID `json:"id"`
	Task
}

func (t *TaskID) Run(debug bool) (res float64) {
	if debug {
		log.Printf("Task %d Running\r\n", t.ID)
	}
	start := time.Now()
	switch t.Operation {
	case "+":
		res = t.Arg1 + t.Arg2
	case "-":
		res = t.Arg1 - t.Arg2
	case "*":
		res = t.Arg1 * t.Arg2
	case "/":
		res = t.Arg1 / t.Arg2
	}
	duration := time.Since(start)
	duration = (time.Millisecond * time.Duration(t.OperationTime)) - duration
	time.Sleep(duration)
	if debug {
		log.Printf("Task %d Completed With Result %.2F\r\n", t.ID, res)
	}
	return
}

func parseString(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func isOperator(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

type TaskID = uint32

var ErrInvalidExpr = errors.New("expression is not valid")
var ErrDivByZero = errors.New("division by zero")

func Calc(expr string, tasks *ConcurrentTaskMap, debug bool) (res ExprResultType, err error) {
	if len(expr) < 3 {
		return 0, ErrInvalidExpr
	}
	var buffer string
	var lastOp rune
	resFlag := false
	openIdx := -1
	parenthesisCount := 0

	if isOperator(rune(expr[0])) || isOperator(rune(expr[len(expr)-1])) {
		return 0, ErrInvalidExpr
	}
	if strings.Contains(expr, "(") || strings.Contains(expr, ")") {
		for i := 0; i < len(expr); i++ {
			value := expr[i]
			if value == '(' {
				if parenthesisCount == 0 {
					openIdx = i
				}
				parenthesisCount++
			}
			if value == ')' {
				parenthesisCount--
				if parenthesisCount == 0 {
					subExpr := expr[openIdx+1 : i]
					calc, err := Calc(subExpr, tasks, debug)
					if err != nil {
						return 0, err
					}
					calcStr := strconv.FormatFloat(calc, 'f', 0, 64)
					expr = strings.Replace(expr, expr[openIdx:i+1], calcStr, 1)
					i -= len(subExpr)
					openIdx = -1
				}
			}
		}
	}
	if openIdx != -1 {
		return 0, ErrInvalidExpr
	}
	priority := strings.ContainsRune(expr, '*') || strings.ContainsRune(expr, '/')
	lowPriority := strings.ContainsRune(expr, '+') || strings.ContainsRune(expr, '-')
	if priority && lowPriority {
		for i := 1; i < len(expr); i++ {
			value := rune(expr[i])
			if value == '*' || value == '/' {
				var startIdx int = i - 1
				if startIdx != 0 {
					for startIdx >= 0 {
						if startIdx >= 0 {
							if isOperator(rune(expr[startIdx])) {
								break
							}
						}
						startIdx--
					}
					startIdx++
				}
				endIdx := i + 1
				if endIdx == len(expr) {
					endIdx--
				} else {
					for !isOperator(rune(expr[endIdx])) && endIdx < len(expr)-1 {
						endIdx++
					}
				}
				if endIdx == len(expr)-1 {
					endIdx++
				}
				subExpr := expr[startIdx:endIdx]
				calc, err := Calc(subExpr, tasks, debug)
				if err != nil {
					return 0, err
				}
				calcStr := strconv.FormatFloat(calc, 'f', 0, 64)
				expr = strings.Replace(expr, expr[startIdx:endIdx], calcStr, 1)
				i -= len(subExpr) - 1
			}
			if isOperator(value) {
				lastOp = value
			}
		}
	}

	for _, value := range expr + "s" {
		switch {
		case value == ' ':
			continue
		case value >= '0' && value <= '9' || value == '.':
			buffer += string(value)
		case isOperator(value) || value == 's':
			if resFlag {
				uuid := uuid.New()
				id := uuid.ID()
				task := Task{
					Arg1:          res,
					Arg2:          parseString(buffer),
					Operation:     string(lastOp),
					Status:        "Wait",
					OperationTime: getOperationTime(lastOp),
					Done:          make(chan struct{}),
				}
				if debug {
					log.Println("Creating New Task With ID", id)
				}
				tasks.Add(id, &task)
				<-task.Done
				res = task.Result
				if debug {
					log.Printf("Result Task %d(%.2F) processed in Calc", id, task.Result)
				}
			} else {
				resFlag = true
				res = parseString(buffer)
			}
			buffer = ""
			lastOp = value
		case value == 's':
		default:
			return 0, ErrInvalidExpr
		}
	}
	return res, nil
}

func getOperationTime(op rune) int {
	switch op {
	case '+':
		return TIME_ADDITION_MS
	case '-':
		return TIME_SUBTRACTION_MS
	case '*':
		return TIME_MULTIPLICATIONS_MS
	case '/':
		return TIME_DIVISIONS_MS
	default:
		return 0
	}
}
