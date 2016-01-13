package main

import (
	"fmt"
	"log"
	// "reflect"
	// "testing"
	"time"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	currentContainers := []docker.APIContainers{}

	foundContainersChan := make(chan []docker.APIContainers)
	quit := make(chan bool)

	client := getDockerClient()

	go readContainers(client, foundContainersChan)

	go func() {
		time.Sleep(15 * time.Second)
		quit <- true
	}()

	for {
		select {
		case conts := <-foundContainersChan:
			if !equalContainerLists(conts, currentContainers) {
				fmt.Println("Not equal", conts, currentContainers)

				currentContainers = conts

				for _, cont := range conts {
					fmt.Println("ID: ", cont.ID[:10], "Name: ", cont.Names,
						"Status ", cont.Status)
				}
			}

		case <-quit:
			fmt.Println("oi")
			return

		case <-time.After(2 * time.Second):
			fmt.Println("containers read timeout")
		}

	}

}

func equalContainerLists(old []docker.APIContainers, new []docker.APIContainers) bool {
	if len(old) != len(new) {
		return false
	}

	for idx, _ := range old {
		if old[idx].ID != new[idx].ID {
			return false
		}
	}

	return true
}

func readContainers(client *docker.Client, ch chan []docker.APIContainers) {
	for {
		conts, err := client.ListContainers(docker.ListContainersOptions{})

		if err != nil {
			fmt.Println("error getting containers", err)
		}

		// fmt.Println("sending read containers")
		ch <- conts

		time.Sleep(1 * time.Second)
	}
}

func getDockerClient() *docker.Client {
	client, err := docker.NewClientFromEnv()

	if err != nil {
		fmt.Println("Client env init error", err)
		log.Panic(err)
	}

	return client
}

/*
func TestListImages(t *testing.T) {
	images := ListImages(testClient
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
*/
