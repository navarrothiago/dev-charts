// sync users and groups to the gitlab api.

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

type Config struct {
	// Gitlab API token.
	Token string `mapstructure:"token"`
	// Gitlab API base url.
	BaseURL string `mapstructure:"base_url"`

	// Gitlab users.
	Users []User `mapstructure:"users"`
	// Gitlab groups.
	Groups []Group `mapstructure:"groups"`
}

type User struct {
	// Gitlab username.
	Username string `mapstructure:"username"`
	// Gilab name.
	Name string `mapstructure:"name"`
	// Gitlab email.
	Email string `mapstructure:"email"`
	// Gitlab password.
	Password string `mapstructure:"password"`
	// Gitlab id.
	ID int `mapstructure:"id"`
}

type Group struct {
	// Gitlab group name.
	Name string `mapstructure:"name"`
	// Gitlab group path.
	Path string `mapstructure:"path"`
}

var config Config

func ListUsers(git *gitlab.Client) {
	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})

	if err != nil {
		panic(err)
	}

	println("============ Users ============")
	for _, user := range users {
		println(user.Username, user.Name, user.Email, user.ID)
	}
}

func ListGroups(git *gitlab.Client) {
	groups, _, err := git.Groups.ListGroups(&gitlab.ListGroupsOptions{})

	if err != nil {
		panic(err)
	}

	println("============ Groups ============")
	for _, group := range groups {
		println(group.Name, group.Path, group.ID)
	}
}

func CreateUsers(git *gitlab.Client) {
	ListUsers(git)

	for _, user := range config.Users {
		_, _, err := git.Users.CreateUser(&gitlab.CreateUserOptions{
			Username: &user.Username,
			Name:     &user.Name,
			Email:    &user.Email,
			Password: &user.Password,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	ListUsers(git)
}

func CreateGroups(git *gitlab.Client) {
	ListGroups(git)

	for _, groups := range config.Groups {
		_, _, err := git.Groups.CreateGroup(&gitlab.CreateGroupOptions{
			Name: &groups.Name,
			Path: &groups.Path,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	ListGroups(git)
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// create new gitlab client.
	git, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.BaseURL), gitlab.WithoutRetries(), gitlab.WithHTTPClient(&http.Client{Transport: tr}))

	if err != nil {
		panic(err)
	}

	CreateUsers(git)
	CreateGroups(git)
}
