package main

import (
	"fmt"
	"os/exec"
    "os"
    "log"
)

func main(){
	file, err := os.Create("ext4dist.txt")
     if err != nil {
       log.Fatal(err)
     }
    cmd := exec.Command("sudo","ext4dist-bpfcc")
    cmd.Stdin = os.Stdin
    cmd.Stdout = file
    cmd.Stderr = os.Stderr
    err1 := cmd.Run(); 
    if err1 != nil {
        fmt.Printf("Got error: %s\n", err.Error())
    }

	

}