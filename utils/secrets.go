package utils

import "os"

type Secrets struct {
	RedisAddr string
	RedisUser string
	RedisPwd  string
}

func LoadSecrets() Secrets {
	s := Secrets{}

	if adrr := os.Getenv("REDIS_HOST"); adrr == "" {
		s.RedisAddr = "127.0.0.1:6379"
	} else {
		s.RedisAddr = adrr
	}

	if user := os.Getenv("REDIS_USER"); user == "" {
		s.RedisUser = ""
	} else {
		s.RedisUser = user
	}

	if pwd := os.Getenv("REDIS_PASSWORD"); pwd == "" {
		s.RedisPwd = ""
	} else {
		s.RedisPwd = pwd
	}

	return s
}
