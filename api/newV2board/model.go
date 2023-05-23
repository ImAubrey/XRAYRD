package newV2board

import (
	"encoding/json"
)

type serverConfig struct {
	shadowsocks
	v2ray
	trojan

	ServerPort int `json:"server_port"`
	BaseConfig struct {
		PushInterval int `json:"push_interval"`
		PullInterval int `json:"pull_interval"`
	} `json:"base_config"`
	Routes []route `json:"routes"`
}

type shadowsocks struct {
	Cipher       string `json:"cipher"`
	Obfs         string `json:"obfs"`
	ObfsSettings struct {
		Path string `json:"path"`
		Host string `json:"host"`
	} `json:"obfs_settings"`
	ServerKey string `json:"server_key"`
}

type v2ray struct {
	Network         string `json:"network"`
	NetworkSettings struct {
		Path             string           `json:"path"`
		Headers          *json.RawMessage `json:"headers"`
		ServiceName      string           `json:"serviceName"`
		Header           *json.RawMessage `json:"header"`
		Show             bool             `mapstructure:"Show"`
		Dest             string           `mapstructure:"Dest"`
		ProxyProtocolVer uint64           `mapstructure:"ProxyProtocolVer"`
		ServerNames      []string         `mapstructure:"ServerNames"`
		PrivateKey       string           `mapstructure:"PrivateKey"`
		MinClientVer     string           `mapstructure:"MinClientVer"`
		MaxClientVer     string           `mapstructure:"MaxClientVer"`
		MaxTimeDiff      uint64           `mapstructure:"MaxTimeDiff"`
		ShortIds         []string         `mapstructure:"ShortIds"`
		EnableREALITY    bool             `mapstructure:"EnableREALITY"`
	} `json:"networkSettings"`
	Tls int `json:"tls"`
}

type trojan struct {
	Host       string `json:"host"`
	ServerName string `json:"server_name"`
}

type route struct {
	Id          int      `json:"id"`
	Match       []string `json:"match"`
	Action      string   `json:"action"`
	ActionValue string   `json:"action_value"`
}

type user struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	SpeedLimit int    `json:"speed_limit"`
}
