package server

import (
	"REST/Report/delivery"
	repository "REST/Report/repository/postgress"
	redis2 "REST/Report/repository/redis"
	config2 "REST/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
)

func Run() {

	config, err := config2.LoadConfig()

	if err != nil {
		fmt.Println("config error")
		return
	}

	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.DBHost, config.Postgres.DBPort, config.Postgres.DBUser, config.Postgres.DBPassword, config.Postgres.DBName)
	db, err := sqlx.Open("postgres", dbConnectionString)
	if err != nil {
		log.Println("ошибка подключение к базе ")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.ADDR,
		Password: config.Redis.PASSWORD, // no password set
		DB:       config.Redis.DB,       // use default DB
	})

	redisRepository := redis2.NewReportRedis(rdb)
	reportRepository := repository.NewReportRepository(db)
	reportHandler := delivery.NewHandler(reportRepository, redisRepository)
	//res, err := reportRepository.InitDB(sql)

	//if err != nil {
	//	log.Fatal(res, err)=
	//}

	r := gin.Default()

	r.GET("/reports", reportHandler.GetAllReport)
	r.POST("/report/create", reportHandler.CreateReport)
	r.DELETE("/report/delete", reportHandler.DeleteReport)
	r.GET("/report/id", reportHandler.GetByID)

	r.Run(":8000")
}
