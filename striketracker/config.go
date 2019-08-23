package striketracker

// Config provides a service configuration for clients.
type Config struct {
	Debug                    bool
	AuthorizationHeaderToken string
	ApplicationID            string
}
