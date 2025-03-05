package application

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MP5s/calculator/pkg/rpn"
)

var RequestInterval = time.Millisecond * 1

func (app *Application) processTask(body io.ReadCloser) {
	app.NumGoroutine++
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	var taskResult GetTaskHandlerResult
	if err := json.Unmarshal(data, &taskResult); err != nil {
		panic(err)
	}

	task := taskResult.Task
	resultValue := task.Run(app.Config.Debug)
	responseData, err := json.Marshal(AgentResult{ID: task.ID, Result: resultValue})
	if err != nil {
		panic(err)
	}

	response, err := app.Agent.Post("http://localhost:8080/api/v1/internal/task", "application/json", bytes.NewReader(responseData))
	if err != nil {
		panic(err)
	}

	if app.Config.Debug {
		log.Println(response.Status)
		bodyContent, _ := io.ReadAll(response.Body)
		log.Println(string(bodyContent))
	}

	app.NumGoroutine--
}

func (app *Application) startAgent() error {
	var errResult error
	doneChannel := make(chan struct{})

	go func() {
		if app.Config.Debug {
			log.Println("Agent Started")
		}
		for {
			<-time.After(RequestInterval)
			if app.NumGoroutine < rpn.COMPUTING_POWER {
				response, err := app.Agent.Get("http://localhost:8080/api/v1/internal/task")
				if err != nil {
					errResult = err
					return
				}
				if response.StatusCode == http.StatusNotFound {
					continue
				}
				defer response.Body.Close()
				go app.processTask(response.Body)
			}
		}
	}()

	<-doneChannel
	return errResult
}
