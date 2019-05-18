package config

import (
	"bytes"
	"fmt"
)

const (
	RedisKeyValueAdapter KeyValueAdapter = "redisgo"
)

type KeyValueAdapter string

type KeyValue struct {
	Adapter  KeyValueAdapter `default:"redisgo"`
	Hostname string          `default:"redisgo"`
	Username string          `default:""`
	Password string          `default:""`
	Port     int             `default:"6379"`
	Params   map[string]interface{}
}

func (s *KeyValue) URL(withAdapter bool) string {
	var buffer bytes.Buffer

	if withAdapter {
		buffer.WriteString(fmt.Sprintf("%s://", s.Adapter))
	}
	buffer.WriteString(fmt.Sprintf("%s:%d", s.Hostname, s.Port))
	return buffer.String()
}
