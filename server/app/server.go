package app

import (
	"flag"
	"fmt"
	"log"
	"mini-marketplace/database/seeders"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (s *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to ", appConfig.AppName)

	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port : ", addr)
	s.Router.Run(addr)
}

func (s *Server) InitializeDB(dbConfig DBConfig) {
	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println("Using mysql driver")
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.DBHost,
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBName,
			dbConfig.DBPort,
		)

		s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		fmt.Println("Using postgres driver")
	}

	if err != nil {
		panic("Failed connecting to database server")
	}

	fmt.Println("Connected to database server")
}

func (s *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := s.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			fmt.Println("Failed to migrate database : ", err)
		}
	}

	fmt.Println("Database migrated sucessfully")
}

func (s *Server) initCommands(appConfig AppConfig, dbConfig DBConfig) {
	s.InitializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				s.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(s.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	appConfig.AppName = getEnv("APP_NAME", "mini-marketplace")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "localhost:8080")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "marketplace")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(appConfig.AppPort)
	}
}
