package models

// Configuration - Holds the configuration variables for this package.
type Configuration struct {
	Auth0  Auth0Settings
	Cors   CorsSettings
	Db     DbSettings
	Server ServerSettings
}

// ServerSettings -
type ServerSettings struct {
	Port string
}

// DbSettings -
type DbSettings struct {
	DbInfo   string
	Name     string
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
