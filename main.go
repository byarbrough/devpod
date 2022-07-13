package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
)

func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	if sock_dir == "" {
		sock_dir = "/tmp"
	}
	socket := "unix:" + sock_dir + "/podman/podman.sock"

	// Connect to Podman socket
	ctx, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(ctx)
	/*
		// Print podman info
		info, err := system.Info(ctx, &system.InfoOptions{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print(info)
	*/
}
