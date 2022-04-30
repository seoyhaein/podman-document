package main

import (
	"context"
	"fmt"
	"github.com/containers/podman/v3/pkg/specgen"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func main() {
	fmt.Println("podman test 시작")

	// Get Podman socket location
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	if sock_dir == "" {
		sock_dir = "/var/run"
	}
	socket := "unix:" + sock_dir + "/podman/podman.sock"

	// Connect to Podman socket
	ctx, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	centos := "cent"
	// container 를 생성해주는데 여기서 bind mount 관련 설정이 있는 듯하다.
	// https://sourcegraph.com/github.com/containers/podman/-/blob/pkg/specgenutil/specgenutil_test.go?L16 참고해서 일단 test 를 해보자.
	// https://sourcegraph.com/search?q=context:global+specgen.NewSpecGenerator&patternType=literal
	s := specgen.NewSpecGenerator(centos, false)
	s.Terminal = true
	r, err := containers.CreateWithSpec(ctx, s, &containers.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container start
	fmt.Println("centos latest start")

	err = containers.Start(ctx, r.ID, &containers.StartOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*	path, err1 := containers.Mount(ctx, centos, &containers.MountOptions{})

		if err1 != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(path)
		}*/

}
