package bb

import (
	"os/user"

	"github.com/pkg/errors"
)

type User struct {
	Name     string   `json:"name,omitempty"`
	Username string   `json:"username"`
	UID      string   `json:"uid"`
	GID      string   `json:"gid"`
	GroupIds []string `json:"group_ids"`
	HomeDir  string   `json:"home_dir"`
}

func ListUsers() ([]User, error) {
	return listUsers()
}

func GetCurrentUser() (*User, error) {
	u, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup current user")
	}
	return parseUserInfo(u)
}

func GetUser(username string) (*User, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup user")
	}
	return parseUserInfo(u)
}

func parseUserInfo(u *user.User) (*User, error) {
	info, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup user")
	}
	gids, err := info.GroupIds()
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup GIDs")
	}
	return &User{
		Username: u.Username,
		Name:     u.Name,
		UID:      u.Uid,
		GID:      u.Gid,
		HomeDir:  u.HomeDir,
		GroupIds: gids,
	}, nil
}
