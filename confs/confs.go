package confs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn      = `host=172.29.246.124 user=root password="123" dbname=postgres port=5432 sslmode=disable`
	net_addr = `localhost:1234`
)

func ConnectToDb() *gorm.DB {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_Name")
	// if host == "" || port == "" || user == "" || pass == "" || dbname == "" {
	// 	panic("database host, port, user, pass or dbname is not defined")
	// }
	// fmt.Sprintf(`host=%s port=%s user=%s pass=%s dbname=%s sslmode=disable`, host, port, user, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to db" + err.Error())
	}
	return db
}

func RunGin(ge *gin.Engine) {
	// host := os.Getenv("HOST_ADDR")
	// port := os.Getenv("HOST_PORT")
	// if host == "" || port == "" {
	// 	panic("host or port is not defined")
	// }
	// ge.Run(fmt.Sprintf("%s:%s", host, port))
	ge.Run(net_addr)
}
