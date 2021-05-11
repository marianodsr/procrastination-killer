package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/marianodsr/ProcrastinationKiller/util"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	found := false
	dayAway := time.Now().Add(time.Hour * 24)

	fecha := time.Now().Format("January 02, 2006")
	golosina := util.GetGolosina()

	message :=
		`=============================================================
																			` + fecha + `	

DEBTH: you failed to code the required time
						 
You owe: ` + golosina + `
		
==============================================================`
	for !found {

		time.Sleep(time.Second * 3)
		_, found = util.GetProcessByName("Code.exe")
		fmt.Println(found)
		if time.Now().After(dayAway) {
			util.SendMail([]byte(message))
			os.Exit(0)
		}
	}

	vsCodePPid, _ := util.GetProcessByName("Code.exe")

	fmt.Println(vsCodePPid)
	vsCode, err := os.FindProcess(vsCodePPid)

	if err == nil {

		deadline := time.Now().Add(time.Hour * 2)

		fmt.Println("SE PROGRAMA HASTA: ")
		fmt.Println(deadline)

		fmt.Println(vsCode.Wait())

		if time.Now().Before(deadline) {

			db, err := sql.Open("sqlite3", "./../database/cupons.db")

			if err != nil {
				fmt.Println(err)
			}

			statement, err := db.Prepare("INSERT INTO cupons (golosina, fecha) VALUES (?, ?)")
			if err != nil {
				fmt.Println(err)
			}

			statement.Exec(golosina, fecha)

			fmt.Println("No cumpliste con el timepo, aca iria el mail")

			util.SendMail([]byte(message))

		} else {

			fmt.Println("Cumpliste con el timepo, congratz!")
			time.Sleep(time.Second * 3)

		}

	} else {

		fmt.Println("No se encontró el proceso")
		time.Sleep(time.Second * 3)

	}

	fmt.Println("Fin del programa")
	time.Sleep(time.Minute * 1)
	os.Exit(0)

}
