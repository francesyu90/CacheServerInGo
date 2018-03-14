package utils

import (
	"fmt"
	"strconv"
	"github.com/spf13/viper"
)

type (
	Utilities struct {
		v *viper.Viper
	}
)

func NewUtilites(v *viper.Viper) *Utilities {
	return &Utilities{v}
}

func (u Utilities) GetIntConfigValue(key string) int {
	return u.v.GetInt(key)
}

func (u Utilities) GetStringConfigValue(key string) string {
	return u.v.GetString(key)
}

func (u Utilities) GetMapArrConfigValue(key string) []map[string]interface{} {

	return toArrayMap(u.v.Get(key))
}

func (u Utilities) GetActiveEnv() int {
	return u.GetIntConfigValue("general.active")
}

func (u Utilities) GetActiveEnvHost() string {	
	prefix := u.getActiveEnvPrefix()
	hostKey := fmt.Sprintf("%s.host", prefix)
	return  u.GetStringConfigValue(hostKey)
} 

func (u Utilities) GetRedisHost() string {	
	prefix := u.getActiveEnvPrefix()
	hostKey := fmt.Sprintf("%s.redis.host", prefix)
	return  u.GetStringConfigValue(hostKey)
} 

/*
	Private methods
*/

func (u Utilities) getActiveEnvPrefix() string {
	env := u.GetActiveEnv()
	envMap := u.GetMapArrConfigValue("env-map")

	switch env {
		case toIntFromInt64Inteface(envMap[0]["index"]):
			return envMap[0]["type"].(string)
		case toIntFromInt64Inteface(envMap[1]["index"]):
			return envMap[1]["type"].(string)
		case toIntFromInt64Inteface(envMap[2]["index"]):
			return envMap[2]["type"].(string)
		case toIntFromInt64Inteface(envMap[3]["index"]):
			return envMap[3]["type"].(string)
		default:
			return ""
	}
}

func toArray(i interface{}) []interface{} {
	arr, ok := i.([] interface{})
	if !ok {
		// TODO: add error handling later
	}
	return arr
} 

func toMap(i interface{}) map[string]interface{} {
	m, ok := i.(map[string]interface{})
	if !ok {
		// TODO: add error handling later
	}
	return m
}

func toArrayMap(i interface{}) []map[string]interface{} {
	arr := toArray(i)

	mapArr := make([]map[string]interface{}, len(arr))

	for i := 0; i < len(arr); i++ {
		mapArr[i] = toMap(arr[i])
	}

	return mapArr
}

func toIntFromInt64Inteface(i interface{}) int {

	intString := fmt.Sprintf("%d", i)
	integer, _:= strconv.ParseInt(intString, 10, 0)

	return int(integer)
}