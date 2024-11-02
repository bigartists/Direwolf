package main

import (
	"fmt"
	"github.com/bigartists/Direwolf/client"
	"github.com/bigartists/Direwolf/internal/module/user"
)

func main() {
	db := client.ProvideDB()
	iUserRepo := user.ProvideUserRepo(db)

	ret, err := iUserRepo.FindUserByUsername("rh1")
	fmt.Println("ret=", ret)
	fmt.Println("err=", err)
}
