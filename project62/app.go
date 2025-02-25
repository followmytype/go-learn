package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// redis
// 預設有16個資料庫，用編號表示0-15，初始為0

// string
//   基本儲存類型，採二進制儲存，也可存放圖片等資料
//   中文字串最大512M
// 設定值 set {key} {value} // 若key存在即修改，若不在則新增
// 查詢值 get {key}
// 刪除值 del {key}
// 設定有期限的值 setex {key} {time:s} {value}
// 批量設定值 mset {key} {value} {key} {value}...
// 批量拿取值 mget {key} {key}...
// 切換庫 select {db}
// 查看key數量 dbsize
// 清空當前庫的資料 flushdb，清空所有資料flushall

// hash
//   類似go的map[string]string，一個key內有其他key-value
// 設定值 hset {key} {field} {value}
// 拿取值 hget {key} {field}
// 拿取key內全部的值 hgetall {key}
// 刪除值 hdel {key}
// 一次設定多個 hmset {key} {field} {value} {field} {value}...
// 一次拿取多個 hmget {key} {field} {field}...
// 判斷是否存在 hexists {key} {field}

// list
//   有序的，可指定添加到頭部或尾部
// 從頭部設值 lpush {key} {value1} {value2} {value3} // 放入value1後再放入value2，接著放value3
// 從尾部設值 rpush {key} {value1} {value2} {value3} // 放入value1後再放入value2，接著放value3
// 拿取值 lrange {key} {start} {end} // end為-1代表最後一個元素，-2代表倒數第二
// 頭部拋出元素 lpop {key}
// 尾部拋出元素 rpop {key}
// 刪除列表 del {key}
// 列表長度 llen {key}
// 當list內的值都被pop完了，則list也會一併消失

// set
//   無序的，元素內無法重複
// 設定成員 sadd {key} {value} {value}...
// 查看成員 smembers {key}
// 判斷是否為成員 sismember {key} {value}
// 移除成員 srem {key} {value}
func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	rdb.Del(ctx, "key")

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	rdb.HSet(ctx, "users", "aaa", `{"userId":"aaa","userPwd":"AAA","userName":"imaaa"}`)
	val3, _ := rdb.HGet(ctx, "users", "aaa").Result()
	fmt.Println(val3)
	rdb.Del(ctx, "users")
}

type User struct {
	UserId   string `json:"userId" redis:"userId"`
	UserPwd  string `json:"userPwd" redis:"userPwd"`
	UserName string `json:"userName" redis:"userName"`
}
