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

### TODO
- 4.x 와 3.x 코드 비교해서 4.x 코드 중심으로 bindins 관련 학습한다. 소스가 업데이트 됨으로 향후 4.x 가 apt install podman  으로 4.x 도 설치 가능하리라 예상.

### podman 기본사용법 (Docker 랑 사용법이 거의 동일 하다.)

```
// podman 이미지 가져오기
podman pull centos
// 이미지 run 시키기(컨테이너 실행시키기)
podman run -it --name cent centos /bin/sh

// container 가 실행중이고 /bin/bash 또는 /bin/sh 등의 bash 등이 실행되어 있다면 docker 와 같이 attach 로 접근할 수 있다.
podman attach cent

// volume mount 시키기
podman run -v /opt:/opt -it --name centPrint01 centos /bin/sh
```

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
### ContainerSecurityConfig
### ContainerCgroupConfig
### ContainerNetworkConfig
### ContainerResourceConfig
### ContainerHealthCheckConfig



