package configs

var PurrfectConfig PawsitiveConfiguration

type PawsitiveConfiguration struct {
	Server   ServerConfig   `json:"server" yaml:"server"`
	Database DatabaseConfig `json:"database" yaml:"database"`
	Session  SessionConfig  `json:"session" yaml:"session"`
}

type ServerConfig struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Domain string `json:"domain" yaml:"domain"`
}

type DatabaseConfig struct {
	Type     DatabaseType     `json:"type" yaml:"type"`
	ArangoDB DBConfigArangoDB `json:"arangodb" yaml:"arangodb"`
}

type DatabaseType string

const DatabaseArangoDB DatabaseType = "arangodb"

type DBConfigArangoDB struct {
	Name     string `json:"name" yaml:"name"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

type SessionConfig struct {
	Name       string `json:"name" yaml:"name"`
	HttpOnly   bool   `json:"httpOnly" yaml:"httpOnly"`
	CSRFSecret string `json:"csrfSecret" yaml:"csrfSecret"`
	CSRFHeader string `json:"csrfHeader" yaml:"csrfHeader"`
	CSRFForm   string `json:"csrfForm" yaml:"csrfForm"`
	SessionKey string `json:"sessionKey" yaml:"sessionKey"`
}
