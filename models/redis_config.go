package models

import "time"

type RedisConfig struct {
	Host         string        `json:"host,omitempty" xml:"host"`
	Port         int           `json:"port,omitempty" xml:"port"`
	User         string        `json:"user,omitempty" xml:"user"`
	Pass         string        `json:"pass,omitempty" xml:"pass"`
	DialTimeout  time.Duration `json:"dialTimeout,omitempty" xml:"dialTimeout"`
	ReadTimeout  time.Duration `json:"readTimeout,omitempty" xml:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout,omitempty" xml:"writeTimeout"`
	PoolSize     int           `json:"poolSize,omitempty" xml:"poolSize"`
	PoolTimeout  time.Duration `json:"poolTimeout,omitempty" xml:"poolTimeout"`
}
