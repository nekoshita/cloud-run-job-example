package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Print("hello cloud run job!")
	log.Print(os.Args)
	log.Printf("CLOUD_RUN_TASK_INDEX = %s", os.Getenv("CLOUD_RUN_TASK_INDEX"))
	log.Printf("CLOUD_RUN_TASK_COUNT = %s", os.Getenv("CLOUD_RUN_TASK_COUNT"))
	log.Printf("FOO = %s", os.Getenv("FOO"))
	log.Printf("BAR = %s", os.Getenv("BAR"))

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		log.Printf("[%d] I'm doing something!", i)
	}
	log.Print("Your task has been successfully completed!!!")
}
