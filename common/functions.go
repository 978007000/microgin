// creator: zacyuan
// date: 2019-12-28

package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanzhangcai/microgin/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ReturnJSON 返回json
func ReturnJSON(ctx *gin.Context, ret errors.ErrorCode, msg string, data ...interface{}) {
	params := make(map[string]interface{})
	params["ret"] = ret
	params["msg"] = msg
	if len(data) > 0 {
		params["data"] = data[0]
	}

	ctx.JSON(200, params)
	ctx.Set("response", params)
	return
}

// TimeToStr 时间戳转日期
func TimeToStr(fmt string, value interface{}) string {
	str := ""
	var sec int64
	switch value.(type) {
	case int:
		sec = int64(value.(int))
	case int64:
		sec = value.(int64)
	case string:
		sec, _ = strconv.ParseInt(value.(string), 10, 64)
	}

	str = time.Unix(sec, 0).Format(fmt)
	return str
}

// StrToTime 日期转时间戳
func StrToTime(fmt string, value string) int64 {
	tm, _ := time.ParseInLocation(fmt, value, time.Local)
	return tm.Unix()
}

// ParseInt64 字符串转int64
func ParseInt64(str string) int64 {
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

// Md5Str 计算md5，返回字符串
func Md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Byte 计算md5，返回字节
func Md5Byte(str string) []byte {
	h := md5.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

// Decimal 保留几位小数
func Decimal(value float64, num int) float64 {
	format := "%." + strconv.Itoa(num) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return value
}

// GetRandomString 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GbkToUtf8 GBK?UTF8??
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Utf8ToGbk UTF8?GBK??
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CloneObject 深度clone对像
func CloneObject(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = CloneObject(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = CloneObject(v)
		}

		return newSlice
	}

	return value
}
