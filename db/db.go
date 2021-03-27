package db

import (
	"fmt"

	_redis "github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db global postgres instance
var DB *gorm.DB

func init() {
	var err error

	// ENV config
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// DB connect
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("DB_HOST"), viper.GetString("DB_PORT"), viper.GetString("DB_USER"), viper.GetString("DB_PASS"), viper.GetString("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to DB")
		panic(err)
	}
	return
}

//GetDB ...
func GetDB() *gorm.DB {
	return DB
}

//RedisClient ...
var RedisClient *_redis.Client

//InitRedis ...
// func InitRedis(params ...string) {

// 	var redisHost = viper.GetString("REDIS_HOST")
// 	var redisPassword = viper.GetString("REDIS_PASSWORD")

// 	db, _ := strconv.Atoi(params[0])

// 	RedisClient = _redis.NewClient(&_redis.Options{
// 		Addr:     redisHost,
// 		Password: redisPassword,
// 		DB:       db,
// 	})
// }

// //GetRedis ...
// func GetRedis() *_redis.Client {
// 	return RedisClient
// }
