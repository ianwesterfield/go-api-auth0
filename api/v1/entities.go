package main

// Configuration - Holds the configuration variables for this package.
type Configuration struct {
	Server ServerSettings
	Db     DbSettings
	Auth0  Auth0Settings
	Cors   CorsSettings
}

// ServerSettings -
type ServerSettings struct {
	Port string
}

// DbSettings -
type DbSettings struct {
	Username string
	Password string
	Name     string
	DbInfo   string
}

// Auth0Settings -
type Auth0Settings struct {
	Audience    string
	Issuer      string
	Certificate string
}

// CorsSettings -
type CorsSettings struct {
	AllowCredentials bool
	Debug            bool
	AllowedOrigins   []string
	AllowedHeaders   []string
}

// Status -
type Status struct {
	Authorized bool `json:"authorized"`
	Version    int  `json:"version"`
}
