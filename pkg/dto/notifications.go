package dto

import (
	"github.com/SSH-Management/protobuf/common"
)

type (
	NewUserNotification struct {
		User         *common.LinuxUser `json:"user,omitempty"`
		PublicSSHKey string           `json:"public_key,omitempty"`
		Groups       []string         `json:"groups,omitempty"`
	}

	NewUserForClientsNotification struct {
		User         *common.LinuxUser `json:"user,omitempty"`
		PublicSSHKey string           `json:"public_key,omitempty"`
		Server       struct {
			Name      string `json:"name,omitempty"`
			IpAddress string `json:"ip,omitempty"`
		} `json:"server,omitempty"`
	}

	UserDeletedNotification struct {
		Username string `json:"username,omitempty"`
	}

	UserDeletedForClientsNotification struct {
		Username string `json:"username,omitempty"`
		Server   struct {
			Name      string `json:"name,omitempty"`
			IpAddress string `json:"ip,omitempty"`
		} `json:"server,omitempty"`
	}
)
