package main

import (
	"fmt"
	"os/exec"
)

func test1() {
	fmt.Println("Go kubers!")
}
func main() {
	cmd := exec.Command("minikube", "start")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Minikube started sucessfully")
		fmt.Println(string(output))
	}
}
