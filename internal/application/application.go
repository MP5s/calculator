package application

import (
	"log"
	"net/http"

	"github.com/MP5s/calculator/internal/web"
	"github.com/MP5s/calculator/pkg/dir"
	"github.com/MP5s/calculator/pkg/rpn"
	"github.com/gorilla/mux"
)

// Хранилище для выражений
var ExpressionStore = make(map[IDExpression]*Expression)

// Хранилище для задач
var TaskStore = rpn.NewConcurrentTaskMap()

// Структура приложения
type Application struct {
	// Конфигурация
	Config           *config
	Agent            http.Client
	ActiveGoroutines int
	Router           *mux.Router
}

// Создание нового экземпляра приложения
func New() *Application {
	return &Application{
		Router: mux.NewRouter(),
		Config: newConfig(),
	}
}

// Запуск сервера
func (app *Application) StartServer() {
	rpn.InitEnv(dir.Env_file())            // Инициализация переменных окружения
	startChannel := make(chan struct{}, 1) // Канал для запуска сервера

	go func() {
		startChannel <- struct{}{}
		if app.Config.Debug {
			log.Println("Orchestrator Started")
		}
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	// Настройка маршрутов
	app.Router.HandleFunc("/api/v1/calculate", app.AddExpressionHandler)
	app.Router.HandleFunc("/api/v1/expressions/{id}", app.GetExpressionHandler)
	app.Router.HandleFunc("/api/v1/expressions", app.GetExpressionsHandler)
	app.Router.HandleFunc("/api/v1/internal/task", app.TaskHandler)

	if app.Config.Web {
		web.HandleToRouter(app.Router)
	}
	http.Handle("/", app.Router)

	<-startChannel // Ожидание запуска сервера
	if err := app.runAgent(); err != nil {
		panic(err)
	}
}
