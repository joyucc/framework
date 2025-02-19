package config

const (
	RedisSessionStore  SessionStore = "redisgo"
	CookieSessionStore SessionStore = "cookie"
)

type SessionStore string

type Session struct {
	Store  SessionStore `yaml:"store" default:"cookie"`
	Secret string       `yaml:"secret"`
	Name   string       `yaml:"name"`
}
