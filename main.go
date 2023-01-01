/*
Copyright © 2022 Bishal Pandit <bishalpandit17@gmail.com>
*/
package main

import (
	"os"
	"log"
	"github.com/bishalpandit/taskbox/cmd"
)

func main() {
	t, err := os.Open("./data/tasks.json")

	if err != nil {
		if err := os.Mkdir("data", os.ModePerm); err != nil {
			log.Fatal(err)
		}
		f, err := os.Create("./data/tasks.json");
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		defer t.Close();
	}

	cmd.Execute()
}
