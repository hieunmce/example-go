package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"example.com/m/config/database/pg"
	"example.com/m/endpoints"
	serviceHttp "example.com/m/http"
	"example.com/m/service"
	bookSvc "example.com/m/service/book"
	categorySvc "example.com/m/service/category"
	lendBookSvc "example.com/m/service/lendbook"
	userSvc "example.com/m/service/user"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3030
// @BasePath  /

func main() {
	// setup env on local
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by errors: %v", err))
		}
	}
	// setup addrr
	httpAddr := ":" + os.Getenv("K_PORT")

	// setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// setup service
	var (
		pgDB, closeDB = pg.New(os.Getenv("PG_DATASOURCE"))
		s             = service.Service{
			UserService: service.Compose(
				userSvc.NewPGService(pgDB),
				userSvc.ValidationMiddleware(),
			).(userSvc.Service),
			CategoryService: service.Compose(
				categorySvc.NewPGService(pgDB),
				categorySvc.ValidationMiddleware(),
			).(categorySvc.Service),
			BookService: service.Compose(
				bookSvc.NewPGService(pgDB),
				bookSvc.ValidationMiddleware(),
			).(bookSvc.Service),
			LendBookService: service.Compose(
				lendBookSvc.NewPGService(pgDB),
				lendBookSvc.ValidationMiddleware(),
			).(lendBookSvc.Service),
		}
	)
	defer closeDB()

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			endpoints.MakeServerEndpoints(s),
			logger,
			os.Getenv("ENV") == "local",
		)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		err := logger.Log("transport", "HTTP", "addr", httpAddr)
		if err != nil {
			return
		}
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	err := logger.Log("exit", <-errs)
	if err != nil {
		return
	}
}
