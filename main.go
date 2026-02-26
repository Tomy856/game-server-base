package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache" // Memcached用
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"          // PostgreSQL用
	"github.com/redis/go-redis/v9" // Redis用
)

var (
	postgresDSN       string
	redisAddr         string
	redisDialTimeout  time.Duration
	memcachedAddr     string
	connectionTimeout time.Duration
)

func init() {
	// 環境変数から読み出し、未設定時はデフォルト値を使用
	postgresDSN = getEnv("POSTGRES_DSN", "host=db user=user password=password dbname=game_db sslmode=disable")
	redisAddr = getEnv("REDIS_ADDR", "redis:6379")
	memcachedAddr = getEnv("MEMCACHED_ADDR", "memcached:11211")
	connectionTimeout = 2 * time.Second
	rangeStr := getEnv("REDIS_DIAL_TIMEOUT", "2s")
	dur, err := time.ParseDuration(rangeStr)
	if err != nil {
		redisDialTimeout = 2 * time.Second
	} else {
		redisDialTimeout = dur
	}
}

// getEnv は環境変数を取得し、未設定ならデフォルトを返す
func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// checkPostgres はPostgreSQLへの接続を確認する
func checkPostgres(ctx context.Context) error {
	db, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

// checkRedis はRedisへの接続を確認する
func checkRedis(ctx context.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		DialTimeout: redisDialTimeout,
	})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

// checkMemcached はMemcachedへの接続を確認する
func checkMemcached() error {
	mc := memcache.New(memcachedAddr)
	if err := mc.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), connectionTimeout)
		defer cancel()

		status := "success!!"
		dbStatus := "OK"
		redisStatus := "OK"
		mcStatus := "OK"

		if err := checkPostgres(ctx); err != nil {
			status = "error"
			dbStatus = fmt.Sprintf("Error: %v", err)
		}
		if err := checkRedis(ctx); err != nil {
			status = "error"
			redisStatus = fmt.Sprintf("Error: %v", err)
		}
		if err := checkMemcached(); err != nil {
			status = "error"
			mcStatus = fmt.Sprintf("Error: %v", err)
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