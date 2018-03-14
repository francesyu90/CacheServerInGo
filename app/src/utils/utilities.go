package utils

import (
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

func (u Utilities) GetMapConfigValue(key string) 
	[]map[string]interface{} 

func (u Utilities) GetActiveEnv() int {
	return u.GetIntConfigValue("general.active")
}

func (u Utilities) GetActiveEnvHost() string {
	env := u.GetActiveEnv()
	
	hostStr := u.GetStringConfigValue("")
} 

/*
	Private methods
*/

func toArray(i interface{}, v *viper.Viper) []interface{} {
	arr, ok := i.([] interface{})
	if !ok {
		// TODO: add error handling later
	}
	return arr
} 

func toMap(i interface{}, v *viper.Viper) map[string]interface{} {
	m, ok := i.(map[string]interface{})
	if !ok {
		// TODO: add error handling later
	}
	return m
}