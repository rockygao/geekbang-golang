package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var (
	rdb *redis.Client
)

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

// 初始化 reds 连接
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", ip, port),
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	_ = context.Background()
}

func main() {
	write(10000, "len10_10k", generateValue(10))
	write(50000, "len10_50k", generateValue(10))
	write(500000, "len10_500k", generateValue(10))

	write(10000, "len1000_10k", generateValue(1000))
	write(50000, "len1000_50k", generateValue(1000))
	write(500000, "len1000_500k", generateValue(1000))

	write(10000, "len5000_10k", generateValue(5000))
	write(50000, "len5000_50k", generateValue(5000))
	write(500000, "len5000_500k", generateValue(5000))

	analysis()
}

func write(n int, key, value string) {
	for i := 0; i < n; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := rdb.Set(k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
