package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	UserId   uint64
	Age      uint8
	NickName string
}

type UserConfig struct {
	UserId      uint64
	AllowUpdate bool
	ProfilePath string
}

func (user User) GetConfigFilename() string {
	return fmt.Sprintf("user_%d.conf", user.UserId)
}

func (user User) GetStringView() string {
	return fmt.Sprintf(
		"Id: %d. Nick name: %s. Age: %d.",
		user.UserId,
		user.NickName,
		user.Age)
}

func MakeUser(userId uint64, age uint8, nickName string) (User, error) {
	if nickName == "" {
		return User{}, errors.New("empty nick name")
	} else if age < 18 {
		return User{}, errors.New("age less than 18")
	}

	return User{userId, age, nickName}, nil
}

func (user User) SpawnConfig() (UserConfig, error) {

	if data, err := ioutil.ReadFile(user.GetConfigFilename()); err == nil {
		var config UserConfig

		s := string(data)
		fmt.Println(s)
		if err := json.Unmarshal(data, &config); err != nil {
			return UserConfig{}, err
		}

		return config, nil

	}

	defaultResult := UserConfig{user.UserId, true, user.GetConfigFilename()}

	return defaultResult, nil
}

func (user User) SaveConfig(config UserConfig) error {
	if jsonData, err := json.Marshal(config); err == nil {
		fmt.Println(string(jsonData))

		file, err := os.Create(config.ProfilePath)

		if err != nil {
			return err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		if _, err = file.WriteString(string(jsonData)); err != nil {
			return err
		}
	} else {
		fmt.Println(err.Error())
	}

	return nil
}
