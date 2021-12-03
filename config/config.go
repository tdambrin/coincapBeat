package config

import (
	"github.com/elastic/beats/libbeat/outputs"
	"time"
)

// Defaults for config variables which are not set
const (
	DefaultSchedule             string        = "@every 10s"
	DefaultTimeout              time.Duration = 10 * time.Second
	DefaultDocumentType         string        = "httpbeat"
	DefaultOutputFormat         string        = "string"
	DefaultJsonDotModeCharacter string        = "_"
)

type HttpbeatConfig struct {
	Hosts []HostConfig
}

type HostConfig struct {
	Schedule             string
	Url                  string
	BasicAuth            BasicAuthenticationConfig `config:"basic_auth"`
	Method               string
	Body                 string
	Headers              map[string]string
	ProxyUrl             string `config:"proxy_url"`
	Timeout              *int64
	DocumentType         string            `config:"document_type"`
	Fields               map[string]string `config:"fields"`
	SSL                  *outputs.TLSConfig
	OutputFormat         string `config:"output_format"`
	JsonDotMode          string `config:"json_dot_mode"`
	JsonDotModeCharacter string `config:"json_dot_mode_character"`
}

type BasicAuthenticationConfig struct {
	Username string
	Password string
}

type ConfigSettings struct {
	Httpbeat HttpbeatConfig
}
