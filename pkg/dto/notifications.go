package dto

type (
	NewUserNotification struct {
		User   CreateUser `json:"user,omitempty"`
		Server struct {
			Name      string `json:"name,omitempty"`
			IpAddress string `json:"ip,omitempty"`
		} `json:"server,omitempty"`
	}

	UserDeletedNotification struct {
		Username string
		Server struct {
			Name      string `json:"name,omitempty"`
			IpAddress string `json:"ip,omitempty"`
		} `json:"server,omitempty"`
	}
)
