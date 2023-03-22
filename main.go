package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	apiKey = ""
)

func main() {

	err := godotenv.Load()
	if err == nil {
		apiKey = os.Getenv("API_KEY")
	}
	c := Client{id: "", apikey: apiKey}

	users, err := c.getUsers()
	printJson(users, err)

	machines, err := c.getMachines("")
	printJson(machines, err)

	user, err := c.postUser(&PostUserName{Name: "testuser"})
	printJson(user, err)
	userId := user.User.ID
	fmt.Printf("user id: %s", userId)

	newusers, err := c.getUsers()
	printJson(newusers, err)
	for _, u := range newusers.Users {
		if u.ID == userId {
			fmt.Printf("found : %s", user.User.Name)
		}
	}

	rerr, err := c.deleteUser("testuser")
	printJson(rerr, err)

	newusers, err = c.getUsers()
	printJson(newusers, err)
	found := false
	for _, u := range newusers.Users {
		if u.ID == userId {
			fmt.Printf("!!found : %s", user.User.Name)
			found = true
		}
	}
	if !found {
		fmt.Printf("user id : %s cant be found", userId)
	}

	router := gin.Default()
	router.Run()
}

func printJson[T Data](data T, err error) {
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		us, _ := json.MarshalIndent(data, "", "   ")
		fmt.Printf(string(us))
	}
}
