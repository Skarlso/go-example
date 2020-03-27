package main

import (
	"log"
	"time"

	sdk "github.com/Skarlso/gosdk/v2"
)

func CreateUser(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("CreateUser has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateUser has been finished!")
	return nil, nil
}

func MigrateDB(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("MigrateDB has been finished!")
	oKey := "AWESOME_KEY"
	oValue := "YES"
	log.Println("Output will be: ", oKey, oValue)
	o := &sdk.Output{
		Key:   oKey,
		Value: oValue,
	}
	out := make([]*sdk.Output, 0)
	out = append(out, o)
	return out, nil
}

func CreateNamespace(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("CreateNamespace has been started!")
	log.Println("Got arguments: ", args)
	log.Println("Values:")
	for _, v := range args {
		log.Print("Value: ", v.Value)
	}
	log.Println()

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateNamespace has been finished!")
	return nil, nil
}

func CreateDeployment(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateDeployment has been finished!")
	return nil, nil
}

func CreateService(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateService has been finished!")
	return nil, nil
}

func CreateIngress(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateIngress has been finished!")
	return nil, nil
}

func Cleanup(args sdk.Arguments) (sdk.Outputs, error) {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("Cleanup has been finished!")
	return nil, nil
}

func main() {
	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
