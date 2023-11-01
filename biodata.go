package main

import (
	"fmt"
	"os"
)

type Friends struct {
	id      uint8
	name    string
	address string
	job     string
	reason  string
}

var friendList = []Friends{
	{1, "Arrayyan", "Gubeng", "Montir", "Karena Golang merupakan bahasa yang praktis dan sederhana dalam penulisan sintaksnya"},
	{2, "Wildan", "Rungkut", "Postman", "Karena Golang merupakan bahasa yang populer"},
	{3, "Fajri", "Semolowaru", "Barista", "Karena Golang merupakan bahasa yang mudah dipahami"},
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Tidak ditemukan nama [name]")
		return
	}

	for _, searchName := range args {
		for _, friend := range friendList {
			if friend.name == searchName {
				fmt.Println("Id:", friend.id)
				fmt.Println("Nama:", friend.name)
				fmt.Println("Alamat:", friend.address)
				fmt.Println("Pekerjaan:", friend.job)
				fmt.Println("Alasan:", friend.reason)
			}
		}
	}
}
