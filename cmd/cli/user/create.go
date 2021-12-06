package user

import (
	"github.com/leebenson/conform"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/pkg/dto"
)

func createUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create",
		RunE:              createUser,
		PersistentPreRunE: command.LoadConfig,
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()

	flags.StringP("username", "u", "", "Username for the new user")
	flags.StringP("email", "i", "", "Email for the new user")
	flags.StringP("name", "n", "", "User's name")
	flags.StringP("surname", "s", "", "User's surname")
	flags.StringP("public-key", "p", "", "Public key of the user")
	flags.StringP("shell", "b", "/bin/bash", "Default shell for the new user")
	flags.StringP("role", "r", "", "Role for the user")
	flags.StringArrayP("groups", "g", []string{}, "Groups of servers where user will be created")
	flags.StringArrayP("system-groups", "m", []string{}, "System groups for all servers (including master)")

	return cmd
}

func getCreateUserStruct(cmd *cobra.Command) (dto.CreateUser, error) {
	flags := cmd.PersistentFlags()

	name, err := flags.GetString("name")
	if err != nil {
		return dto.CreateUser{}, err
	}

	surname, err := flags.GetString("surname")
	if err != nil {
		return dto.CreateUser{}, err
	}

	email, err := flags.GetString("email")
	if err != nil {
		return dto.CreateUser{}, err
	}

	username, err := flags.GetString("username")
	if err != nil {
		return dto.CreateUser{}, err
	}

	shell, err := flags.GetString("shell")
	if err != nil {
		return dto.CreateUser{}, err
	}

	role, err := flags.GetString("role")
	if err != nil {
		return dto.CreateUser{}, err
	}

	groups, err := flags.GetStringArray("groups")
	if err != nil {
		return dto.CreateUser{}, err
	}

	systemGroups, err := flags.GetStringArray("system-groups")
	if err != nil {
		return dto.CreateUser{}, err
	}

	publicKey, err := flags.GetString("public-key")
	if err != nil {
		return dto.CreateUser{}, err
	}

	pass := password.MustGenerate(20, 2, 2, false, true)

	return dto.CreateUser{
		User: dto.User{
			Name:         name,
			Surname:      surname,
			Username:     username,
			Email:        email,
			Password:     pass,
			Shell:        shell,
			Role:         role,
			Groups:       groups,
			SystemGroups: systemGroups,
		},
		PublicSSHKey: publicKey,
	}, nil
}

func createUser(cmd *cobra.Command, args []string) error {
	c := command.GetContainer()

	ctx := cmd.Context()

	userService := c.GetUserService()
	validator := c.GetValidator()

	userDto, err := getCreateUserStruct(cmd)

	if err != nil {
		return err
	}

	if err := conform.Strings(&userDto); err != nil {
		return err
	}

	if err := validator.Struct(userDto); err != nil {
		return err
	}

	u, _, err := userService.Create(ctx, userDto)
	if err != nil {
		return err
	}

	c.GetDefaultLogger().
		Info().
		Str("username", u.Username).
		Str("password", userDto.User.Password).
		Msg("New User created successfully")

	return nil
}
