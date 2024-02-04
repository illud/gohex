package base

import "os"

func BaseDbClient(clientName string) {
	// Add database client
	clientString := ""
	if clientName == "mysql" {
		clientString =
			`package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// db The database connection
	db *sql.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := sql.Open("mysql", "databaseUsername:databasePassword@tcp(localhost:3306)/yourDatabaseTablename")

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	// defer db.Close()
	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *sql.DB {
	return db
}`
		//Adds db conection to main.go
		AppendDbConnectionToMain()
	}

	if clientName == "gorm" {
		clientString =
			`package database
		
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// db The database connection
	db *gorm.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := gorm.Open("mysql", "databaseUsername:databasePassword@tcp(127.0.0.1:3306)/yourDatabaseTablename?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *gorm.DB {
	return db
}`
		//Adds db conection to main.go
		AppendDbConnectionToMain()
	}

	// 	if clientName == "prisma" {

	// 		dir, err := os.Getwd()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(dir)
	// 		var ss []string
	// 		if runtime.GOOS == "windows" {
	// 			ss = strings.Split(dir, "\\")
	// 		} else {
	// 			ss = strings.Split(dir, "/")
	// 		}

	// 		currentDirName := ss[len(ss)-1]

	// 		clientString =
	// 			`package data

	// import (
	// 	"fmt"

	// 	"` + currentDirName + `/data/prisma/db"
	// 	"golang.org/x/net/context"
	// )

	// var (
	// 	// db The database connection
	// 	prismaDdb *db.PrismaClient
	// )

	// func Connect(){
	// 	client := db.NewClient()
	// 	if err := client.Prisma.Connect(); err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	// defer func() {
	// 	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 	// 		panic(err)
	// 	// 	}
	// 	// }()
	// 	prismaDdb = client
	// 	fmt.Println("CONNECTED")
	// }

	// func Client() *db.PrismaClient {
	// 	return prismaDdb
	// }

	// var Context = ContextService()

	// func ContextService() context.Context {
	// 	ctx := context.Background()
	// 	return ctx
	// }`

	// 		//Adds db conection to main.go
	// 		AppendDbConnectionToMain()

	// 		//Insertdata into prisma.schema
	// 		prismaString :=
	// 			`datasource db {
	// 	// could be postgresql or mysql
	// 	provider = "sqlite"
	// 	url      = "file:dev.db"
	// }

	// generator db {
	// 	provider = "go run github.com/prisma/prisma-client-go"
	// 	// set the output folder and package name
	// 	   output           = "./infraestructure/databases/prisma/db"
	// 	   package          = "db"
	// }

	// //This is and example table add your own schemas
	// model Tasks {
	// 	id        Int      @id @default(autoincrement())
	// 	createdAt DateTime @default(now())
	// 	updatedAt DateTime @updatedAt
	// 	title     String
	// 	description String
	// }`

	// 		prismaSChemaBytes := []byte(prismaString)
	// 		os.WriteFile("schema.prisma", prismaSChemaBytes, 0)
	// 	}

	//Add data to db.go
	clientBytes := []byte(clientString)
	os.WriteFile("adapters/database/db.go", clientBytes, 0)
}
