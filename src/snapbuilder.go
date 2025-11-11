package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	snapd_installer := exec.Command("snap", "install", "snapd", "--classic")

	err := snapd_installer.Run()

	if err != nil {

		fmt.Println(err)
	}

	lxd_installer := exec.Command("snap", "install", "lxd")

	err = lxd_installer.Run()

	if err != nil {

		fmt.Println(err)
	}
	user := os.Getenv("USER")
	lxd_group_add := exec.Command("sudo", "usermod", "-a", "-G", "lxd", user)

	err = lxd_group_add.Run()

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println("Log out and log in to apply group changes and then re-run this program!")

	lxd_init := exec.Command("sudo", "lxd", "init", "--auto")

	err = lxd_init.Run()

	if err != nil {

		fmt.Println(err)
	}

	build := exec.Command("snapcraft", "build")
	err = build.Run()

	if err != nil {

		fmt.Println(err)
	}

	pack := exec.Command("snapcraft", "pack")

	err = pack.Run()

	if err != nil {

		fmt.Println(err)
	}
}
