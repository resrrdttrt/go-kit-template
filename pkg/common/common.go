package common

import "os"

func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type GeneralRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessRes(data interface{}) GeneralRes {
	return GeneralRes{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

