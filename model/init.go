package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB
var RS *redis.Pool

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	fmt.Println("mysql connect start")
	db, err := gorm.Open("postgres", connString)
	//db.LogMode(true)
	db.SingularTable(true)
	// Error
	if err != nil {
		panic(err)
	}
	//svc.CreatLoop()
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	migration()
	fmt.Println("数据库初始化成功")
}

func RedisPool(redisUrl string, password string) {
	RS = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			conn, e = redis.DialURL(redisUrl)
			if e != nil {
				return nil, fmt.Errorf("redis connection error: %s", e)
			}
			//验证redis密码
			if _, authErr := conn.Do("AUTH", password); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}
			return conn, e
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
	fmt.Println("Redis连接成功")
}
