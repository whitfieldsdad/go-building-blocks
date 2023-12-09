package bb

import (
	"github.com/yusufpapurcu/wmi"
)

type Win32_UserAccount struct {
	Name string
}

func listUsers() ([]User, error) {
	var win32Users []Win32_UserAccount
	err := wmi.Query("SELECT * FROM Win32_UserAccount", &win32Users)
	if err != nil {
		return nil, err
	}
	var users []User
	for _, u := range win32Users {
		user, err := GetUser(u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}
