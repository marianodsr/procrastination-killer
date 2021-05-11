package util

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/mitchellh/go-ps"
)

//GetProcessByName func
func GetProcessByName(name string) (int, bool) {

	processList, _ := ps.Processes()

	parentIds := []int{}

	check := make(map[int]int)

	found := false

	for _, process := range processList {
		if process.Executable() == name {
			ppid := process.PPid()
			parentIds = append(parentIds, ppid)
			check[ppid] = 0
			found = true

		}

	}

	if !found {
		return 0, found
	}

	for _, ppid := range parentIds {
		if _, exists := check[ppid]; exists {
			check[ppid]++
		}
	}

	max := 0
	finalPPid := 0
	for key, value := range check {
		if value > max {
			max = value
			finalPPid = key
		}

	}

	fmt.Println(finalPPid)
	return finalPPid, found
}

type sender struct {
	username string
	password string
}

//SendMail func
func SendMail(message []byte) error {
	host := "smtp.gmail.com"
	port := "587"
	adress := host + ":" + port
	sender := &sender{
		username: "nitocoding@gmail.com",
		password: "m159123456",
	}
	receiver := []string{
		"em.combiancamignani@gmail.com",
	}
	auth := smtp.PlainAuth("", sender.username, sender.password, host)

	err := smtp.SendMail(adress, auth, sender.username, receiver, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

//GetGolosina func
func GetGolosina() string {

	candyBag := []string{
		"Chocolate",
		"Billiken(50pe)",
		"Gomitas",
		"Jugo Ades",
		"Oreos",
		"Item de librería",
		"Fibron copic",
	}

	rand.Seed(time.Now().UnixNano())

	index := len(candyBag)

	return candyBag[rand.Intn(index)]

}
