package config

// === App ===

// Http http config
type Http struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// Rpc http config
type Rpc struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// Web config
type Web struct {
	Url string `json:"url" yaml:"url"`
}

// Gateway config
type Gateway struct {
	Url string `json:"url" yaml:"url"`
}

// Plugin spider config
type Plugin struct {
	Path string `json:"path" yaml:"path"`
}

// Storage config
type Storage struct {
	Path string `json:"path" yaml:"path"`
}

// === Middleware ===

// Mysql config
type Mysql struct {
	Url string `json:"url" yaml:"url"`
}

// Redis config
type Redis struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
}

// Influx config
type Influx struct {
	Token  string `json:"token" yaml:"token"`
	Org    string `json:"org" yaml:"org"`
	Bucket string `json:"bucket" yaml:"bucket"`
	Url    string `json:"url" yaml:"url"`
}

// Rabbitmq config
type Rabbitmq struct {
	Url string `json:"url" yaml:"url"`
}

// Jaeger config
type Jaeger struct {
	Reporter    struct {
		LocalAgentHostPort string `json:"localAgentHostPort" yaml:"localAgentHostPort"`
	} `json:"reporter" yaml:"reporter"`
	Sampler struct {
		Type  string  `json:"type" yaml:"type"`
		Param float64 `json:"param" yaml:"param"`
	} `json:"sampler" yaml:"sampler"`
}

// Nats config
type Nats struct {
	Url string `json:"url" yaml:"url"`
}

// === Vendor ===

// Slack config
type Slack struct {
	Verification string `json:"verification" yaml:"verification"`
	Signing      string `json:"signing" yaml:"signing"`
	Token        string `json:"token" yaml:"token"`
	Webhook      string `json:"webhook" yaml:"webhook"`
}

// Rollbar config
type Rollbar struct {
	Token       string `json:"token" yaml:"token"`
	Environment string `json:"environment" yaml:"environment"`
}

// Telegram config
type Telegram struct {
	Token string `json:"token" yaml:"token"`
}
