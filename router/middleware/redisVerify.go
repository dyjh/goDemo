package middleware

import (
	"crypto/md5"
	"demo/datasource"
	"encoding/binary"
	"encoding/hex"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
	"time"
)

func RedisHandler(ctx iris.Context) {
	Authorization := ctx.GetHeader("Authorization")
	AuthArr := strings.Fields(Authorization)
	if len(AuthArr) != 2 || AuthArr[0] != "bearer"{
		_, _ = ctx.WriteString("token的格式为bearer ***************")
		return
	}
	conn, _ := datasource.InitRedis()
	verifyRes, _ := redis.String(conn.Do("GET", AuthArr))
	if verifyRes == "" {
		_, _ = ctx.WriteString("token已过期")
		return
	}
	ctx.Next() // 继续下一个handler 这里是after
}

func ParseClaimsConfigRedis(ctx iris.Context) UserInfo {
	Authorization := ctx.GetHeader("Authorization")
	AuthArr := strings.Fields(Authorization)
	conn, _ := datasource.InitRedis()
	verifyRes, _ := redis.Int64(conn.Do("GET", AuthArr))
	userInfo := UserInfo{Id:verifyRes}
	return userInfo
}


func GetRedisToken(userId int64) string {
	// 获取当前时间的时间戳
	t := time.Now().Unix()
	str := strconv.FormatInt(t, 10) + "|" + strconv.FormatInt(userId, 10)
	intStr, _ := strconv.ParseInt(str, 10, 64)
	// 生成一个MD5的哈希
	h := md5.New()

	// 将时间戳转换为byte，并写入哈希
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(intStr))
	h.Write([]byte(b))

	// 将字节流转化为16进制的字符串
	return hex.EncodeToString(h.Sum(nil))
}