package env

import (
	"os"
	"strconv"
	"strings"
)

// 不存在返回默认值
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// 整数类型环境变量，不存在或转换失败返回默认值
func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// 布尔类型环境变量，不存在返回默认值
func GetEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// 环境变量 MY_ARRAY="a,b,c,d"
func GetEnvStringArr(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		if value == "" {
			return defaultValue
		}
		return strings.Split(value, ",")
	}
	return defaultValue
}
