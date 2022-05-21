### 정리하기.

### 시작하기 전
먼저 외부에서 접근하기 위해서는 restful api 를 호출하는 형태를 가져야 한다. 이를 위해서는 podman 을 홀로? 사용할 경우 docker 와 달리 service 실행 할 필요가 없다.(데몬 시작)
하지만, 외부에서 접근할 경우 그것이 프로세스 차원이던 인스턴스 차원이던 접근하기 위해서는 podman 은 docker 의 데몬을 살려두는 것처럼(server) podman 을 서버로서 실행시켜줘야 하기때문에 docker와 비슷하게 서비스 또는 socket 가 살아 있어야 한다.

이것은, 현재 apt (ubuntu 기준) install 로 4.x 로 설치가 가능하지 않음으로 github 에서 소스를 받아와서 빌드 후 이후 메뉴얼하게 진행해야 한다. 이 과정에 대한 것은 추후 문서로 남긴다.

### 생각하기 
현재까지 파악한 바로는 binding 관련 api 들은 https://github.com/containers/podman/tree/main/pkg 여기에 모여 지고 있는 듯 하다.
그리고 pod 관련해서는 https://github.com/containers/podman/tree/main/pkg/bindings/pods 여기에서 작성되고 있는 것 같지만, 세부적으로는 구현이 안되어 있는 듯한 느낌을 받았다.

즉, 구현이 안되어 있다라는 뜻은 client 또는 remote 관점에서 접근하기 위한 api 가 충분히 만들어 지지 않았다라는 느낌을 받았다. 하지만 이것을 처리하는 서버측 api 는 개발이 이미 끝난듯한 느낌을 받았다.

아직, 테스트 전이기 때문에 잘못된 정보 일 수도 있지만, 일단 현재까지로서는 이렇게 판단을 했다.

구현에 대한 restful api 에 대한 상세 설명은 아래 레퍼런스에서 파악할 수 있다.

https://docs.podman.io/en/latest/Reference.html

--volume-from 의 경우 컨테이너와 pod 가 마운트 한 경우인데, 해당 컨테이너의 경우는 pod 밖의 컨테이너 인가? 

왜냐하면 pod 에 들어가기 위해서는 컨테이너 형태로 존재해야 하는데 pod 를 만들때 컨테이너ID 또는 name 을 요구하기 때문에 pod 밖의 컨테이너로 의심함.

그렇다면, pod 외부에 data 컨테이너를 두고, 이것을 pod 내의 컨테이너가 이러한 데이터를 공유하는 것인가?

### 처리 상황
#### pod crate --volume-from 을 통해서 pod 와 컨테이너간의 volume 공유를 테스트 진행할 예정임.

먼저 alpine:latest 기준으로 Containerfile 을 제작해주고 자료 확인을 위해서 bash 와 nano 를 설치 해준다.
- 자주 잊어버려서 기억용으로 기록해둔다.
- data.Conatainerfile
```
FROM alpine:latest
RUN apk update && apk add --no-cache bash nano
RUN mkdir -p /opt/data
RUN chmod a+rw /opt/data
COPY ./usage.md /opt/data

```
data.Conatainerfile 로 이동후
```
podman build -t seoy-data:latest . -f data.Containerfile

# 생성 이미지 확인
podman images

# data container 를 detach mode 로 restart=always 로 하고 command 는 /bin/bash 로 해주고 디버깅을 위해서 접근하기 위해 -it 옵션을 추가 해주었다.
podman run -d --restart=always --name=data -it seoy-data /bin/bash

# 컨테이너의 현재 상태를 확인한다.
podman ps

# pod 를 생성해주면서 volume-from 을 통해서 data container 와 연결 시킨다.
podman pod create --name poddata --volumes-from=data

# 현재까지 infra container 만 생성되었다. 이제 pod 내에 컨테이너를 연결 시킨다.
podman run -it --name=pod-container01 --pod poddata docker.io/library/alpine:latest /bin/sh

# 해당 컨테이너에 마운트되지 않았다. 
# pod 를 조사 해보자
podman info --debug
podman pod inspect poddata 


```

### 문제점 인식하기
- https://docs.oracle.com/en/learn/storage_podman_containers/#using-volumes-with-containers

일단 위와 같이 작성했을 때 mount 가 되지 않는다. 그 이유를 추측해보면, 해당 소스 컨테이너에 volume 이 없다고 판단이 된다. 
한번 volume 여부를 살펴보자

```
podman inspect -f '{{.Mounts}}' data

```

살펴보면 null 값이다. 만약 volume 을 create 해주면 어떻게 되는지 살펴본다.