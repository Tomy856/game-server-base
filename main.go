package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache" // Memcached用
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"          // PostgreSQL用
	"github.com/redis/go-redis/v9" // Redis用
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// 1. PostgreSQLの接続確認
		dsn := "host=db user=user password=password dbname=game_db sslmode=disable"
		db, _ := sql.Open("postgres", dsn)
		dbErr := db.Ping()

		// 2. Redisの接続確認
		rdb := redis.NewClient(&redis.Options{
			Addr: "redis:6379",
		})
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		redisErr := rdb.Ping(ctx).Err()

		// 3. Memcachedの接続確認 (ホスト名は docker-compose で指定した "memcached")
		mc := memcache.New("memcached:11211")
		mcErr := mc.Ping()

		// 判定と結果送信
		status := "success!!"
		dbStatus := "OK"
		redisStatus := "OK"
		mcStatus := "OK"

		if dbErr != nil {
			dbStatus = fmt.Sprintf("Error: %v", dbErr)
			status = "error"
		}
		if redisErr != nil {
			redisStatus = fmt.Sprintf("Error: %v", redisErr)
			status = "error"
		}
		if mcErr != nil {
			mcStatus = fmt.Sprintf("Error: %v", mcErr)
			status = "error"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      status,
			"postgres":    dbStatus,
			"redis":       redisStatus,
			"memcached":   mcStatus,
			"server_time": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.Run(":8080")
}
