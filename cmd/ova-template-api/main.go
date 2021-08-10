package ova_template_api

import "fmt"

func Hello() {
	fmt.Println(getGreeting())
}

func getGreeting() string {
	return "Hey! This is \"ova-template-api\" service"
}
