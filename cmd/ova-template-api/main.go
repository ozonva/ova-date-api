package ova_template_api

import (
	"fmt"
	"ova-date-api/internal/settings"
)

func Hello() {
	fmt.Println(getGreeting())
}

func getGreeting() string {
	return "Hey! This is \"ova-template-api\" service"
}

func Task3() {
	if user, err := settings.MakeUser(11, 27, "Nike"); err == nil {
		fmt.Println(user.GetStringView())

		config, _ := user.SpawnConfig()

		config.AllowUpdate = !config.AllowUpdate

		if err := user.SaveConfig(config); err != nil {
			fmt.Println(err.Error())

			return
		}
	}
}
