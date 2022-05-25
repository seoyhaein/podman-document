# podman-document
## pdoman 의 설치와 binding 에 대한 설명을 기술한다.
podman 은 docker 와 호환된다. docker 에서 작성된 이미지 또한 podman 을 통해서 사용할 수 있다.

docker 에 비해서 여러가지 장점들이 있다.

### 필수 설치 및 필요사항 (2022.04.28 기준)
- Ubuntu 22.04 LTS 설치 필요
- golang 1.18 이상(1.16 이상 이긴한데, 1.16 일때 알수 없는 문제가 발생해서 그냥 1.18.1 로 업데이트 함.)
- goland 20.01 버전 사용하고 있었는데 golang 1.18 goroot 설정에서 문제가 발생하여 goland 22.01 버전으로 업그레이드 함.
- 현재(22.04.28) podman 을 ubuntu 에 최신버전으로 설치하면 version 이 3.4.4 이다. 하지만 현재 최신 버전은 4.x 이다.
- 바인딩시 코드에서 사용하는 버전과 설치되어 있는 podman 버전이 일치해야 한다. 만약 다를시 에러 발생한다.
- podman bindings 사용하여 개발시 반드시 go mod 사용해야 한다. 없으면 문제 발생.

### 관련 링크
- [podman 설치](https://podman.io/getting-started/installation#build-and-run-dependencies)
- [podman 설치 및 실행 예시](https://gochronicles.com/podman/)
- [podman bindings](https://podman.io/blogs/2020/08/10/podman-go-bindings.html)
- [podman trouble shooting](https://github.com/containers/podman/blob/main/troubleshooting.md)
- [podman Demo](https://github.com/containers/Demos)
- [podman restapi manual](https://docs.podman.io/en/latest/Reference.html)

### Volume mount & Bind mount
-[참고](https://www.daleseo.com/docker-volumes-bind-mounts/)

### podman 기본사용법 (Docker 랑 사용법이 거의 동일 하다.)

- 기본 사용법

```
// podman 이미지 가져오기
podman pull centos

// 이미지 리스트 확인 
// podman image list 와 동일
podman images


// 이미지 run 시키기(컨테이너 실행시키기)
podman run -it --name cent centos /bin/sh

// container 가 실행중이고 /bin/bash 또는 /bin/sh 등의 bash 등이 실행되어 있다면 docker 와 같이 attach 로 접근할 수 있다.
podman attach cent

// 그렇지 않을 경우
podman exec -it [container-name or container-id] /bin/bash

// bind mount 시키기
podman run -v /opt:/opt -it --name centPrint01 centos /bin/sh

```

- 모든 컨테이너 삭제

```

// 여기서 -q 옵션은 컨테이너 아이디를 한줄씩 출력해준다.
// 모든 컨테이너를 중지시킨다.
podman stop $(podman ps -a -q) 또는  podman stop $(podman ps -aq)

// 모두 중지된 컨테이너를 삭제한다.
podman rm $(podman ps -a -q)

// 위의 두과정을 거치지 않고 -f 을 주면 ㅎ강제로 중지시키고 삭제가 가능하다.
podman rm -f $(podman ps -a -q)

// 중지된 컨테이너만 삭제한다. (테스트 해보자.)
podman container prune

```
#### podman pull 또는 이미지를 가져올때 아래와 같은 문제가 발생할때
- 아래코드는 mariadb 이미지를 가져올때 발생했다.
- 이에 대한 해결책은 docker.io 를 붙이면 된다. [참고](https://url.kr/cgvhkx)

```
podman pull mariadb
> short-name "mariadb" did not resolve to an alias and no unqualified-search registries are defined in "/etc/containers/registries.conf"

podman pull docker.io/mariadb
```

#### podman run Detached mode
Detached mode: run the container in the background and print the new container ID. The default is false.

At any time you can run podman ps in the other shell to view a list of the running containers.

You can reattach to a detached container with podman attach.

- podman run --dt or podman run -d

### Pod 관련 (mesos container 관련해서도 한번 정리하자. https://mesos.apache.org/documentation/latest/)
- 쿠버네티스의 pod 참고, podman 의 pod 와 비슷함으로 [참고](https://kubernetes.io/ko/docs/concepts/workloads/pods/)
- https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml#enough_teasing__show_me_the_goods 읽고 지우기.

The podman generate kube command allows you to export your existing containers into Kubernetes Pod YAML. This YAML can then be imported into OpenShift or a Kubernetes cluster. The podman play kube does the opposite, it allows you to take a Kubernetes YAML and run it in Podman.

[여기서 가져옴-podman.io](https://podman.io/) 결룩 쿠버네틱스 pod 는 podman 의 pod 라는 의미 인거 같다. 좀더 살펴보자.


- pod 에 있는 컨테이너나 pod 는 podman 에서 삭제할 수 없다. 즉, podman rm -f 컨테이너ID 로 삭제할 수 없다.

모든 Podman 포드에는 인프라 컨테이너가 포함되어 있습니다. 이 컨테이너는 아무 작업도 수행하지 않지만 잠자기 상태로 전환됩니다. 그 목적은 포드와 연결된 네임스페이스를 보유하고 포드맨이 다른 컨테이너를 포드에 연결할 수 있도록 하는 것입니다. 이를 통해 POD 내에서 컨테이너를 시작 및 중지할 수 있으며 포드는 계속 실행됩니다. 기본 컨테이너가 포드를 제어하는 ​​것처럼 이는 불가능합니다. 기본 인프라 컨테이너는 k8s.gcr.io/pause이미지를 기반으로 합니다 . 달리 명시하지 않는 한 모든 포드에는 기본 이미지 기반 컨테이너가 있습니다.

Pod를 구성하는 대부분의 속성은 실제로 infra 컨테이너에 할당됩니다. 포트 바인딩, cgroup-parent 값 및 커널 네임스페이스는 모두 infra 컨테이너에 할당됩니다. 포드가 생성되면 이러한 속성이 인프라 컨테이너에 할당되고 변경할 수 없기 때문에 이를 이해하는 것이 중요합니다. 예를 들어 포드를 생성한 다음 나중에 새 포트를 바인딩하는 컨테이너를 추가하기로 결정하면 포드맨은 이를 수행할 수 없습니다. 새 컨테이너를 추가하기 전에 추가 포트 바인딩으로 포드를 다시 생성해야 합니다.

위의 다이어그램에서 각 컨테이너 위에 있는 상자를 확인하십시오. 이것은 컨테이너 모니터(conmon)입니다. 작은 C 프로그램이 하는 일은 컨테이너의 기본 프로세스를 감시하고 컨테이너가 죽으면 종료 코드를 저장하는 것입니다. 또한 나중에 첨부할 수 있도록 컨테이너의 tty를 열어둡니다. 이것은 podman이 분리 모드(백그라운드)에서 실행되도록 하여 podman은 종료할 수 있지만 conmon은 계속 실행됩니다. 각 컨테이너에는 고유한 conmon 인스턴스가 있습니다.

[podman pod create](https://docs.podman.io/en/latest/markdown/podman-pod-create.1.html)

- podman pod create --name test
- podman pod ls (pod 컨테이너 리스트)

[podman pod rm](https://docs.podman.io/en/latest/markdown/podman-pod-rm.1.html)

- podman pod rm test
- podman pod rm -fa (강제로(f) 모든(a) 컨테이너를 삭제) 



#### podman pod 에 컨테이너 넣는 두가지 방법 [참고](https://url.kr/fygb8s)
- 첫번째 방법

```
// 먼저 pod 를 test 라는 이름으로 생성한다.
podman pod create --name test 

//생성된 pod 를 확인 한다.
// ps, ls, list 는 동일하다.
podman ls

// --pod 옵션을 붙이면 pod 의 infra 컨테이너들도 나온다.

// 모든 컨테이너와 pod infra 컨테이너
podman ps -a --pod

// pod 컨테이너 만 나타난다.
podman ps --pod

// 이제 새로운 컨테이너늘 만들고 이것을 pod 에 연결 시킨다.
// detach mode 로 최신 alpine 컨테이너를 실행시키고 top 명령어를 실행시켰다. --pod 옵션으로 test 라는 이름의 pod 에 연결 시켰다.
podman run -dt --pod test docker.io/library/alpine:latest top

```
- 두번째 방법
- 아래 코드에서 --pod new:myapp_pod 를 보면 새로운 pod 를 myapp_pod 라고 지어주었다.
- The use of new: indicates to Podman that we want to create a new pod rather than attempt to assign the container to an existing pod.
```
podman run -d --restart=always --pod new:myapp_pod \
-e MYSQL_ROOT_PASSWORD="myrootpass"  \
-e MYSQL_DATABASE="wp-db"  \
-e MYSQL_USER="wp-user"  \
-e MYSQL_PASSWORD="w0rdpr3ss"  \
--name=wptest-db docker.io/mariadb

```

#### pod 관련 api 힌트 - 찾기 중
1. https://github.com/containers/podman/blob/d3903a85910979d8212028cf814574047015db58/libpod/runtime_pod.go
2. "github.com/containers/podman/v4/pkg/bindings/pods"
3. https://github.com/containers/podman/search?q=NewPod
4. github.com/containers/podman/v4/libpod -> libpod.Runtime.NewPod
5. https://github.com/containers/podman/blob/c3d871a3f6cc7a94c5e86782ba63e05cd1d2faeb/pkg/specgen/generate/pod_create.go

## podman binding api 정리
먼저,  specgen.NewSpecGenerator 함수를 통해서 SpecGenerator 정해준다. 이 SpecGenerator 에 저장된 정보를 통해서 컨테이너를 생성해준다. 

컨테이너 생성은 containers.CreateWithSpec 함수가 담당한다. 이후에 podman 사용시 명령어와 binding api는 거의 매칭이 되고 Restful Api 이다. 

예를 들어서 컨테이너 생성후 컨테이너가 존재할 경우는 podman start (컨테이너 ID or 컨테이너 Name) 은 binding api start 와 동일하다.

### SpecGenerator
### ContainerStorageConfig 
- Image string 컨테이너로 사용할 이미지 
- Rootfs string rootfs는 루트 파일 시스템을 뜻하며, 리눅스 파일 시스템을 미리 패키지화 해놓은 바이너리.
- Image 와 Rootfs 둘중하나는 설정되어 있어야 한다.(당연한 소리)
- ImageVolumeMode string image volume 을 어떻게 생성할지를 결정한다. optional 하며 설정을 하지 않거나 "anonymous" 로 설정하면  익명으로 설정한다. (create as anonymous volumes).
- 그외 "ignore" (do not create), "tmpfs" (create as tmpfs) 설정값이 존재한다.
- VolumesFrom []string 컨테이너의 volume 의 소스이다. * volume 좀더 파악해서 보완하자.
- Init bool 

Init은 리눅스 커널 부팅이 완료된 뒤 실행되는 첫 번째 프로세스다. 또한 동시에 Init은 커널이 직접 실행하는 유일한 프로세스다. 

따라서 Init은 부모 프로세스를 가지지 않는 유일한 프로세스인 동시에, Init을 제외한 나머지 모든 프로세스의 조상이 된다.

docker 기준 설명

docker run 수행 시 --init 옵션이 주어지지 않을 경우는 container 내에서 init process 를 별도로 기동하지 않는다. docker run 수행 시 넘겨준 command(/bin/bash)가 그대로 1 번 process 가 된다.
반대로 docker run 수행 시 --init 옵션이 주어질 경우, init process 를 container 구동 후 1 번 process 로 기동하게 된다. 

container 내에서 init process 를 1 번으로 구동한다는 것은 중요한 의미가 있다. 이는 child process 를 받아주어 resource 의 누수나 zombie process 의 생성 등을 방지하는 init system 의 역할을 container 내에서 수행한다는 뜻이기 때문이다.

init process 로 사용되는 default binary 는 /bin/docker-init 을 사용한다. (정확하게는 which docker-init 의 결과로 찾아지는 binary 를 사용) docker-init 은 container 외부에서 별도로 기동되거나 하는 process 가 아니다. 

container 내에서 첫 번째로 기동되어 마치 Host 에서의 init process 처럼 동작하도록 만들어진 프로그램이라고 생각하면 된다.

- InitPath string 위의 Init 이 true 이면 설정해줘야 하며 Init 바이너리의 위치가 기록된다. If not specified, the default set in the Libpod config will be used.
- Mounts []spec.Mount 컨테이너에 추가할 마운트들??? Image Volumes 과 VolumesFrom volumes 이 충돌할때 대체한다.???

spec.Mount 는 https://github.com/opencontainers/runtime-spec 에 정의 되어 있고 아래 코드는 https://github.com/opencontainers/runtime-spec/tree/main/specs-go 에서의 config.go 에 설정되어 있다.

```
// Mount specifies a mount for a container.
    type Mount struct {
	// Destination is the absolute path where the mount will be placed in the container.
	Destination string `json:"destination"`
	// Type specifies the mount kind.
	Type string `json:"type,omitempty" platform:"linux,solaris"`
	// Source specifies the source path of the mount.
	Source string `json:"source,omitempty"`
	// Options are fstab style mount options.
	Options []string `json:"options,omitempty"`
    }
```

- Volumes []*NamedVolume 

named volum은 Docker(Linux에서는 /var/lib/docker/volume/)가 관리하는 Host File System의 일부에 Data가 저장된다.

/specgen/volumes.go 에 있음.
```
    // NamedVolume holds information about a named volume that will be mounted into
    // the container.
    type NamedVolume struct {
	// Name is the name of the named volume to be mounted. May be empty.
	// If empty, a new named volume with a pseudorandomly generated name
	// will be mounted at the given destination.
	Name string
	// Destination to mount the named volume within the container. Must be
	// an absolute path. Path will be created if it does not exist.
	Dest string
	// Options are options that the named volume will be mounted with.
	Options []string
    }
```
- OverlayVolumes []*OverlayVolume [참고- 예전에 docker 할때 문제가 있어서 자료 조사 했었는데 잊어버림.](https://www.joinc.co.kr/w/man/12/docker/storage)

- // Image volumes bind-mount a container-image mount into the container.
- // Optional.
- ImageVolumes []*ImageVolume `json:"image_volumes,omitempty"`

위에서 Mounts []spec.Mount 가 대신 할 수 있다고 했다.

```
    // Devices are devices that will be added to the container.
	// Optional.
	Devices []spec.LinuxDevice `json:"devices,omitempty"`
	// DeviceCGroupRule are device cgroup rules that allow containers
	// to use additional types of devices.
	DeviceCGroupRule []spec.LinuxDeviceCgroup `json:"device_cgroup_rule,omitempty"`
	// IpcNS is the container's IPC namespace.
	// Default is private.
	// Conflicts with ShmSize if not set to private.
	// Mandatory.
	IpcNS Namespace `json:"ipcns,omitempty"`
	// ShmSize is the size of the tmpfs to mount in at /dev/shm, in bytes.
	// Conflicts with ShmSize if IpcNS is not private.
	// Optional.
	ShmSize *int64 `json:"shm_size,omitempty"`

```
- directory 설정

```
    // WorkDir is the container's working directory.
	// If unset, the default, /, will be used.
	// Optional.
	WorkDir string `json:"work_dir,omitempty"`
	// Create the working directory if it doesn't exist.
	// If unset, it doesn't create it.
	// Optional.
	CreateWorkingDir bool `json:"create_working_dir,omitempty"`
```

- 나머지는 그냥 이렇게 넣는다.

```
    // RootfsPropagation is the rootfs propagation mode for the container.
	// If not set, the default of rslave will be used.
	// Optional.
	RootfsPropagation string `json:"rootfs_propagation,omitempty"`
	// Secrets are the secrets that will be added to the container
	// Optional.
	Secrets []Secret `json:"secrets,omitempty"`
	// Volatile specifies whether the container storage can be optimized
	// at the cost of not syncing all the dirty files in memory.
	Volatile bool `json:"volatile,omitempty"`

```

### ContainerBasicConfig
- Name string 컨테이너 이름, 세팅 안되면 랜덤하게 세팅됨.
- Pod string Pod is the ID of the pod the container will join. [참고1](https://url.kr/nvfhsi) [참고2](https://url.kr/dfaq6m)
- Entrypoint []string // Entrypoint is the container's entrypoint. If not given and Image is specified, this will be populated by the image's configuration.
- Command []string // Command is the container's command. If not given and Image is specified, this will be populated by the image's configuration.
- EnvHost bool 호스트의 env (환경)이 컨테이너에 추가될지 결정.
- HTTPProxy bool // EnvHTTPProxy indicates that the http host proxy environment variables should be added to container.
- Env map[string]string // Env is a set of environment variables that will be set in the container.
- Terminal bool // Terminal is whether the container will create a PTY. PTY 는 원격접속을 의미한다. [참고3, TTY, PTY, PTS](https://url.kr/o9viub)
- Stdin bool // Stdin is whether the container will keep its STDIN open. 옵션 값인데, default 값이 무엇인지, 그리고 stdin 이 컨테이너에서 open 이 되면 어떤지 테스트 할것.
- Labels map[string]string //Labels are key-value pairs that are used to add metadata to containers.
- Annotations map[string]string // Annotations are key-value options passed into the container runtime that can be used to trigger special behavior. 예시 찾아보자.
- StopSignal *syscall.Signal // StopSignal is the signal that will be used to stop the container. Must be a non-zero integer below SIGRTMAX. If not provided, the default, SIGTERM, will be used. Will conflict with Systemd if Systemd is set to "true" or "always".
- [참고4, syscall.Signal,SIGRTMAX,SIGTERM](https://en.wikipedia.org/wiki/Signal_(IPC))
- StopTimeout *uint // StopTimeout is a timeout between the container's stop signal being sent and SIGKILL being sent. If not provided, the default will be used. If 0 is used, stop signal will not be sent, and SIGKILL will be sent instead.
- Timeout uint //// Timeout is a maximum time in seconds the container will run before main process is sent SIGKILL. If 0 is used, signal will not be sent. Container can run indefinitely
- LogConfiguration *LogConfig // LogConfiguration describes the logging for a container including driver, path, and options. 이건 테스트 해보야 할듯. 관련 예시 찾자.
- ConmonPidFile string  // ConmonPidFile is a path at which a PID file for Conmon will be placed. If not given, a default location will be used.
- [참고5, Conmon, An OCI container runtime monitor.](https://github.com/containers/conmon) 살펴보자.
- RawImageName string // RawImageName is the user-specified and unprocessed input referring to a local or a remote image. 살펴보자. 자료조사 필요
- RestartPolicy string // RestartRetries is the number of attempts that will be made to restart the container. Only available when RestartPolicy is set to "on-failure".
- [참고6, podman restart policy](https://asciinema.org/a/240306)
- [참고7, docker restart policy](https://url.kr/ic7s2v)
- RestartRetries *uint // RestartRetries is the number of attempts that will be made to restart the container. Only available when RestartPolicy is set to "on-failure".
- OCIRuntime string // OCIRuntime is the name of the OCI runtime that will be used to create the container. If not specified, the default will be used.
- Systemd string 	// Systemd is whether the container will be started in systemd mode. Valid options are "true", "false", and "always". "true" enables this mode only if the binary run in the container is /sbin/init or systemd. "always" unconditionally enables systemd mode.
- "false" unconditionally disables systemd mode. If enabled, mounts and stop signal will be modified. If set to "always" or set to "true" and conditionally triggered, conflicts with StopSignal. If not specified, "false" will be assumed.

- 나머지는 코드로 대체함.
```
// Determine how to handle the NOTIFY_SOCKET - do we participate or pass it through
	// "container" - let the OCI runtime deal with it, advertise conmon's MAINPID
	// "conmon-only" - advertise conmon's MAINPID, send READY when started, don't pass to OCI
	// "ignore" - unset NOTIFY_SOCKET
	SdNotifyMode string `json:"sdnotifyMode,omitempty"`
	// Namespace is the libpod namespace the container will be placed in.
	// Optional.
	Namespace string `json:"namespace,omitempty"`
	// PidNS is the container's PID namespace.
	// It defaults to private.
	// Mandatory.
	PidNS Namespace `json:"pidns,omitempty"`
	// UtsNS is the container's UTS namespace.
	// It defaults to private.
	// Must be set to Private to set Hostname.
	// Mandatory.
	UtsNS Namespace `json:"utsns,omitempty"`
	// Hostname is the container's hostname. If not set, the hostname will
	// not be modified (if UtsNS is not private) or will be set to the
	// container ID (if UtsNS is private).
	// Conflicts with UtsNS if UtsNS is not set to private.
	// Optional.
	Hostname string `json:"hostname,omitempty"`
	// Sysctl sets kernel parameters for the container
	Sysctl map[string]string `json:"sysctl,omitempty"`
	// Remove indicates if the container should be removed once it has been started
	// and exits
	Remove bool `json:"remove,omitempty"`
	// ContainerCreateCommand is the command that was used to create this
	// container.
	// This will be shown in the output of Inspect() on the container, and
	// may also be used by some tools that wish to recreate the container
	// (e.g. `podman generate systemd --new`).
	// Optional.
	ContainerCreateCommand []string `json:"containerCreateCommand,omitempty"`
	// PreserveFDs is a number of additional file descriptors (in addition
	// to 0, 1, 2) that will be passed to the executed process. The total FDs
	// passed will be 3 + PreserveFDs.
	// set tags as `json:"-"` for not supported remote
	// Optional.
	PreserveFDs uint `json:"-"`
	// Timezone is the timezone inside the container.
	// Local means it has the same timezone as the host machine
	// Optional.
	Timezone string `json:"timezone,omitempty"`
	// DependencyContainers is an array of containers this container
	// depends on. Dependency containers must be started before this
	// container. Dependencies can be specified by name or full/partial ID.
	// Optional.
	DependencyContainers []string `json:"dependencyContainers,omitempty"`
	// PidFile is the file that saves container process id.
	// set tags as `json:"-"` for not supported remote
	// Optional.
	PidFile string `json:"-"`
	// EnvSecrets are secrets that will be set as environment variables
	// Optional.
	EnvSecrets map[string]string `json:"secret_env,omitempty"`
	// InitContainerType describes if this container is an init container
	// and if so, what type: always or once
	InitContainerType string `json:"init_container_type"`
	// Personality allows users to configure different execution domains.
	// Execution domains tell Linux how to map signal numbers into signal actions.
	// The execution domain system allows Linux to provide limited support
	// for binaries compiled under other UNIX-like operating systems.
	Personality *spec.LinuxPersonality `json:"personality,omitempty"`

```

### ContainerSecurityConfig
- 몇번 이슈된적이 있는데 중요도는 낮으나, 잘 사용할려면 살펴두자.
```
// ContainerSecurityConfig is a container's security features, including
// SELinux, Apparmor, and Seccomp.
type ContainerSecurityConfig struct {
	// Privileged is whether the container is privileged.
	// Privileged does the following:
	// - Adds all devices on the system to the container.
	// - Adds all capabilities to the container.
	// - Disables Seccomp, SELinux, and Apparmor confinement.
	//   (Though SELinux can be manually re-enabled).
	// TODO: this conflicts with things.
	// TODO: this does more.
	Privileged bool `json:"privileged,omitempty"`
	// User is the user the container will be run as.
	// Can be given as a UID or a username; if a username, it will be
	// resolved within the container, using the container's /etc/passwd.
	// If unset, the container will be run as root.
	// Optional.
	User string `json:"user,omitempty"`
	// Groups are a list of supplemental groups the container's user will
	// be granted access to.
	// Optional.
	Groups []string `json:"groups,omitempty"`
	// CapAdd are capabilities which will be added to the container.
	// Conflicts with Privileged.
	// Optional.
	CapAdd []string `json:"cap_add,omitempty"`
	// CapDrop are capabilities which will be removed from the container.
	// Conflicts with Privileged.
	// Optional.
	CapDrop []string `json:"cap_drop,omitempty"`
	// SelinuxProcessLabel is the process label the container will use.
	// If SELinux is enabled and this is not specified, a label will be
	// automatically generated if not specified.
	// Optional.
	SelinuxOpts []string `json:"selinux_opts,omitempty"`
	// ApparmorProfile is the name of the Apparmor profile the container
	// will use.
	// Optional.
	ApparmorProfile string `json:"apparmor_profile,omitempty"`
	// SeccompPolicy determines which seccomp profile gets applied
	// the container. valid values: empty,default,image
	SeccompPolicy string `json:"seccomp_policy,omitempty"`
	// SeccompProfilePath is the path to a JSON file containing the
	// container's Seccomp profile.
	// If not specified, no Seccomp profile will be used.
	// Optional.
	SeccompProfilePath string `json:"seccomp_profile_path,omitempty"`
	// NoNewPrivileges is whether the container will set the no new
	// privileges flag on create, which disables gaining additional
	// privileges (e.g. via setuid) in the container.
	NoNewPrivileges bool `json:"no_new_privileges,omitempty"`
	// UserNS is the container's user namespace.
	// It defaults to host, indicating that no user namespace will be
	// created.
	// If set to private, IDMappings must be set.
	// Mandatory.
	UserNS Namespace `json:"userns,omitempty"`
	// IDMappings are UID and GID mappings that will be used by user
	// namespaces.
	// Required if UserNS is private.
	IDMappings *types.IDMappingOptions `json:"idmappings,omitempty"`
	// ReadOnlyFilesystem indicates that everything will be mounted
	// as read-only
	ReadOnlyFilesystem bool `json:"read_only_filesystem,omitempty"`
	// Umask is the umask the init process of the container will be run with.
	Umask string `json:"umask,omitempty"`
	// ProcOpts are the options used for the proc mount.
	ProcOpts []string `json:"procfs_opts,omitempty"`
	// Mask is the path we want to mask in the container. This masks the paths
	// given in addition to the default list.
	// Optional
	Mask []string `json:"mask,omitempty"`
	// Unmask is the path we want to unmask in the container. To override
	// all the default paths that are masked, set unmask=ALL.
	Unmask []string `json:"unmask,omitempty"`
}
```
### ContainerCgroupConfig
- 일단 생략

### ContainerNetworkConfig
- 일단 생략

### ContainerResourceConfig
- 아래 코드로 대신하는데, 일단 컨테이너 잡을 구성할때 Resource 에 대한 제한을 둬야 하기 때문에 실제로 잘 살펴봐야 한다.

```
// ContainerResourceConfig contains information on container resource limits.
type ContainerResourceConfig struct {
	// ResourceLimits are resource limits to apply to the container.,
	// Can only be set as root on cgroups v1 systems, but can be set as
	// rootless as well for cgroups v2.
	// Optional.
	ResourceLimits *spec.LinuxResources `json:"resource_limits,omitempty"`
	// Rlimits are POSIX rlimits to apply to the container.
	// Optional.
	Rlimits []spec.POSIXRlimit `json:"r_limits,omitempty"`
	// OOMScoreAdj adjusts the score used by the OOM killer to determine
	// processes to kill for the container's process.
	// Optional.
	OOMScoreAdj *int `json:"oom_score_adj,omitempty"`
	// Weight per cgroup per device, can override BlkioWeight
	WeightDevice map[string]spec.LinuxWeightDevice `json:"weightDevice,omitempty"`
	// IO read rate limit per cgroup per device, bytes per second
	ThrottleReadBpsDevice map[string]spec.LinuxThrottleDevice `json:"throttleReadBpsDevice,omitempty"`
	// IO write rate limit per cgroup per device, bytes per second
	ThrottleWriteBpsDevice map[string]spec.LinuxThrottleDevice `json:"throttleWriteBpsDevice,omitempty"`
	// IO read rate limit per cgroup per device, IO per second
	ThrottleReadIOPSDevice map[string]spec.LinuxThrottleDevice `json:"throttleReadIOPSDevice,omitempty"`
	// IO write rate limit per cgroup per device, IO per second
	ThrottleWriteIOPSDevice map[string]spec.LinuxThrottleDevice `json:"throttleWriteIOPSDevice,omitempty"`
	// CgroupConf are key-value options passed into the container runtime
	// that are used to configure cgroup v2.
	// Optional.
	CgroupConf map[string]string `json:"unified,omitempty"`
	// CPU period of the cpuset, determined by --cpus
	CPUPeriod uint64 `json:"cpu_period,omitempty"`
	// CPU quota of the cpuset, determined by --cpus
	CPUQuota int64 `json:"cpu_quota,omitempty"`
}
```
### ContainerHealthCheckConfig
- 관련 예시를 찾아보자.

```
// ContainerHealthCheckConfig describes a container healthcheck with attributes
// like command, retries, interval, start period, and timeout.
type ContainerHealthCheckConfig struct {
	HealthConfig *manifest.Schema2HealthConfig `json:"healthconfig,omitempty"`
}
```

### tts/pts [참고](https://codedragon.tistory.com/4211)


### 참고자료
https://eehoeskrap.tistory.com/245
https://minholee93.tistory.com/entry/Linux-Process-Status-PS
프로세스 세션 리더 와 프로세스 그룹 리더  https://blueyikim.tistory.com/89

PROCESS STATE CODES
Here are the different values that the s, stat and state output
specifiers (header "STAT" or "S") will display to describe the state of
a process:

               D    uninterruptible sleep (usually IO)
               R    running or runnable (on run queue)
               S    interruptible sleep (waiting for an event to complete)
               T    stopped by job control signal
               t    stopped by debugger during the tracing
               W    paging (not valid since the 2.6.xx kernel)
               X    dead (should never be seen)
               Z    defunct ("zombie") process, terminated but not reaped by
                    its parent

       For BSD formats and when the stat keyword is used, additional
       characters may be displayed:

               <    high-priority (not nice to other users)
               N    low-priority (nice to other users)
               L    has pages locked into memory (for real-time and custom IO)
               s    is a session leader
               l    is multi-threaded (using CLONE_THREAD, like NPTL pthreads
                    do)
               +    is in the foreground process group

[참고, 그냥 심심할때](https://www.samsungsds.com/kr/insights/docker.html)

꼭 읽어보기(정리전)
https://mkdev.me/posts/the-tool-that-really-runs-your-containers-deep-dive-into-runc-and-oci-specifications

https://chhanz.github.io/container/2020/09/22/podman-build-flask-example-app/

네임스페이스
https://linuxtut.com/en/61f1291f6ee804531328/
https://www.44bits.io/ko/keyword/linux-namespace

기초
https://docs.oracle.com/en/learn/storage_podman_containers/#introduction
https://phoenixnap.com/kb/podman-tutorial

containerfile
https://meta.stackoverflow.com/questions/407966/generalize-dockerfile-to-containerfile-for-now-and-the-future
https://www.mankier.com/5/Containerfile

alpine
https://blog.naver.com/PostView.nhn?blogId=ki630808&logNo=222149370156

cmd vs entrypoint
https://bluese05.tistory.com/77
