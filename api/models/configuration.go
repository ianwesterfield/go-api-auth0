package models

// Configuration - Holds the configuration variables for this package.
type Configuration struct {
	Auth0  Auth0Settings  `json:"auth0,omitempty" bson:",omitempty"`
	Cors   CorsSettings   `json:"cors,omitempty" bson:",omitempty"`
	Db     DbSettings     `json:"db,omitempty" bson:",omitempty"`
	Server ServerSettings `json:"server,omitempty" bson:",omitempty"`
}

// ServerSettings -
type ServerSettings struct {
	Port int
}

// DbSettings -
type DbSettings struct {
	Dsn      string
	Server   string
	Database string
	Password string
	Username string
}

// Auth0Settings -
type Auth0Settings struct {
	Audience    string
	Certificate string
	Issuer      string
}

// CorsSettings -
type CorsSettings struct {
	AllowCredentials bool
	AllowedHeaders   []string
	AllowedOrigins   []string
	Debug            bool
}
