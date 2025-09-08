// In config/config.go
package config

// Config is the main structure for the .portal.toml file
type Config struct {
	Enter PortalAction `mapstructure:"enter"`
	Leave PortalAction `mapstructure:"leave"`
}

// PortalAction defines the commands and message for an action
type PortalAction struct {
	Commands []string `mapstructure:"commands"`
	Message  string   `mapstructure:"message"`
}
