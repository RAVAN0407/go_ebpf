package main

import (
	"fmt"
	"os/exec"
	"time"
	"github.com/shirou/gopsutil/v3/process"
	"log"

)
func command1(){
	//cmd,err:= exec.Command("sudo","ext4dist-bpfcc")
	out:= exec.Command("sudo","ext4dist-bpfcc")
	err :=out.Start()
	fmt.Println(" code is here ")
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(" code is there ")
    fmt.Println(out)

	// ctx, cancel := context.WithTimeout(context.Background(),  10000*time.Millisecond)
	// defer cancel()

	// if err := exec.CommandContext(ctx,"ext4dist-bpfcc" ).Run(); err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(ctx)

	// 	// This will fail after 100 milliseconds. The 5 second sleep
	// 	// will be interrupted.
	// }

	// fmt.Println("there was no error")


}

func KillProcess(name string) error {
    processes, err := process.Processes()
    if err != nil {
        return err
    }
    for _, p := range processes {
        n, err := p.Name()
        if err != nil {
            return err
        }
        if n == name {
            return p.Kill()
        }
    }
    return fmt.Errorf("process not found")
}
func main(){
	command1()
	fmt.Println(" hi ")
	time.Sleep(5 * time.Second)
	fmt.Println(" hello ")
	KillProcess("ext4dist-bpfcc")
	fmt.Println("output")
  

	

}