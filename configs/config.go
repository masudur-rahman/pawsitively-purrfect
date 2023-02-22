package configs

var PurrfectConfig PawsitiveConfiguration

type PawsitiveConfiguration struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Session  SessionConfig  `json:"session"`
}

type ServerConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Domain string `json:"domain"`
}

type DatabaseConfig struct {
	Type     DatabaseType     `json:"type"`
	ArangoDB DBConfigArangoDB `json:"arangodb"`
}

type DatabaseType string

const DatabaseArangoDB DatabaseType = "arangodb"

type DBConfigArangoDB struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type SessionConfig struct {
	Name       string `json:"name"`
	HttpOnly   bool   `json:"httpOnly"`
	CSRFSecret string `json:"csrfSecret"`
}
