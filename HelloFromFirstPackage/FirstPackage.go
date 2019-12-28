package main

import (
	"flag"
	"fmt"
	firstPackage "github.com/Renat11111/GoLangBasic/firstPacage"
)

func main() {

	var (
		userid   = flag.String("U", "", "login_id")
		password = flag.String("P", "", "password")
		server   = flag.String("S", "localhost", "server_name[\\instance_name]")
		database = flag.String("d", "", "db_name")
	)

	fmt.Println(firstPackage.HelloFromShahinozavrik())
	fmt.Println(userid)
	fmt.Println(password)
	fmt.Println(server)
	fmt.Println(database)

}
