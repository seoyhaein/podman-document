package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/libpod/define"
	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
)

// 일단 v4 에서 글러온 소스도 정상작동하니 method 와 function 을 1:1 비교해보자. -> 자료로 만들어 보자.
func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	if sock_dir == "" {
		sock_dir = "/var/run"
	}
	socket := "unix:" + sock_dir + "/podman/podman.sock"

	// Connect to Podman socket
	// TODO 다른 connection 도 살펴보자.
	ctx, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Pull Busybox image (Sample 1)
	fmt.Println("Pulling Busybox image...")
	_, err = images.Pull(ctx, "docker.io/busybox", &images.PullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Pull Fedora image (Sample 2)
	rawImage := "registry.fedoraproject.org/fedora:latest"
	fmt.Println("Pulling Fedora image...")
	_, err = images.Pull(ctx, rawImage, &images.PullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List images
	imageSummary, err := images.List(ctx, &images.ListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var names []string
	for _, i := range imageSummary {
		names = append(names, i.RepoTags...)
	}
	fmt.Println("Listing images...")
	fmt.Println(names)

	// Container create
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Terminal = true
	r, err := containers.CreateWithSpec(ctx, s, &containers.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container start
	fmt.Println("Starting Fedora container...")
	err = containers.Start(ctx, r.ID, &containers.StartOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = containers.Wait(ctx, r.ID, &containers.WaitOptions{
		Condition: []define.ContainerStatus{define.ContainerStateRunning},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container list
	var latestContainers = 1
	containerLatestList, err := containers.List(ctx, &containers.ListOptions{
		Last: &latestContainers,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Latest container is %s\n", containerLatestList[0].Names[0])

	// Container inspect
	ctrData, err := containers.Inspect(ctx, r.ID, &containers.InspectOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Container uses image %s\n", ctrData.ImageName)
	fmt.Printf("Container running status is %s\n", ctrData.State.Status)

	// Container stop
	fmt.Println("Stopping the container...")
	err = containers.Stop(ctx, r.ID, &containers.StopOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctrData, err = containers.Inspect(ctx, r.ID, &containers.InspectOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Container running status is now %s\n", ctrData.State.Status)
	return

}
