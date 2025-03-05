package application

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/MP5s/calculator/pkg/rpn"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Обработчик для добавления выражения через POST http://localhost:8080/api/v1/calculate
// Тело запроса: {"expression": "<выражение>"}
func (app *Application) AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	expressionStr, exists := request["expression"]
	if !exists {
		http.Error(w, "Missing expression", http.StatusBadRequest)
		return
	}

	id := uuid.New().ID()
	expression := Expression{expressionStr, WaitStatus, 0}
	Expressions[id] = &expression

	go func() {
		result, err := rpn.Calc(expressionStr, Tasks, app.Config.Debug)
		if err != nil {
			expression.Status = err.Error()
		} else {
			expression.Status = "OK"
			expression.Result = result
		}
	}()

	responseData, err := json.Marshal(AddHandlerResult{id})
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

// Обработчик для получения выражения по ID через GET http://localhost:8080/api/v1/expressions/{id}
func (app *Application) GetExpressionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	expressionID := IDExpression(id)
	expression, exists := Expressions[expressionID]
	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	responseData, err := json.Marshal(GetExpressionHandlerResult{ExpressionWithID{expressionID, *expression}})
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}

// Обработчик для получения всех выражений через GET http://localhost:8080/api/v1/expressions
func (app *Application) GetExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var expressionsWithID []ExpressionWithID
	for id, expression := range Expressions {
		expressionsWithID = append(expressionsWithID, ExpressionWithID{id, *expression})
	}

	responseData, err := json.Marshal(GetExpressionsHandlerResult{expressionsWithID})
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

// Обработчик для задач через GET и POST http://localhost:8080/api/v1/internal/task
func (app *Application) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer r.Body.Close()
		var taskID rpn.TaskID

		for id, task := range *Tasks.Map() {
			if task.Status == WaitStatus {
				taskID = rpn.TaskID{Task: *task, ID: id}
				break
			}
		}

		if taskID.ID == 0 {
			http.Error(w, "No tasks available", http.StatusNotFound)
			return
		}

		responseData, err := json.Marshal(GetTaskHandlerResult{taskID})
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		(*Tasks.Get(taskID.ID)).Status = CalculationStatus
		w.Write(responseData)

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var result rpn.AgentResult
		if err := json.Unmarshal(body, &result); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		task := Tasks.Get(result.ID)
		if task == nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		task.Result = result.Result
		task.Done <- struct{}{}
		task.Status = "OK"

		if app.Config.Debug {
			log.Printf("Result Task %d (%.2f) is handled in TasksMap", result.ID, task.Result)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task result processed successfully"))
	}
}
