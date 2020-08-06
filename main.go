package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func CreateUser(args sdk.Arguments) error {
	log.Println("CreateUser has been started!")
	log.Println("Update!")
	log.Println("Update2!")
	log.Println("Update3!")
	log.Println("Update4!")
	log.Println("Update5!")
	log.Println("Update6!")
	log.Println("Update7!")
	log.Println("Update8!")
	log.Println("Update9!")
	log.Println("Update10!")
	log.Println("Update11!")
	log.Println("Update12!")
	log.Println("Update13!")
	log.Println("Update14!")
	log.Println("Update15!")
	log.Println("Update16!")
	log.Println("Update17!")
	log.Println("Update18!")
	log.Println("Update19!")
	log.Println("Update20!")
	log.Println("Update21!")
	log.Println("Update22!")
	log.Println("Update23!")
	log.Println("Update24!")
	log.Println("Update25!")
	log.Println("Update26!")
	log.Println("Update27!")	
	log.Println("Update28!")
	log.Println("Update29!")	
	log.Println("Update30!")		
	
	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateUser has been finished!")
	return nil
}

func MigrateDB(args sdk.Arguments) error {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("MigrateDB has been finished!")
	return nil
}

func CreateNamespace(args sdk.Arguments) error {
	log.Println("CreateNamespace has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateNamespace has been finished!")
	return nil
}

func CreateDeployment(args sdk.Arguments) error {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateDeployment has been finished!")
	return nil
}

func CreateService(args sdk.Arguments) error {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateService has been finished!")
	return nil
}

func CreateIngress(args sdk.Arguments) error {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateIngress has been finished!")
	return nil
}

func Cleanup(args sdk.Arguments) error {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("Cleanup has been finished!")
	return nil
}

func main() {
	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
