package models

import "time"

type RedisConfig struct {
	Host         string
	Port         int
	User         string
	Pass         string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}
