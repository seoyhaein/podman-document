## 공유메모리 (shared memory)

```
df -f 
```

먼저, 파일 시스템을 확인해보면 /dev/shm 의 사이즈를 확인 할 수 있다. /dev/shm 은 tmpfs 로 설정되어 있는 것을 확인할 수 있으며 이 디렉토리는 공유 메모리이고 물리 메모리의 반으로 설정되어 있다.

하지만 실제로 점유되고 있는 것은 아니고 사용한 만큼 소모 되는 형태이다. 실제 RAM 에 저장됨으로 reboot 하면 휘발성으로 저장된 데이터는 사라지게 된다.

그리고, 설정한 사이즈를 초과하게 되면 SWAP 영역으로 넘어가게 되고, 과도하게 사용할 경우 시스템이 작동 불능이 될 수 있다. (리부팅해야함.)

### pod 내에서 공유메모리 만들기 (podman-pod-create.md 참고)
```
# pod 를 생성해주면서 volume-from 을 통해서 data container 와 연결 시킨다.
podman pod create --name poddata --volumes-from=data

podman run -d --restart=always --name=pod-container01 docker.io/library/alpine:latest 
podman run -d --restart=always --name=pod-container02 docker.io/library/alpine:latest 

```
위와 같이 작성했을 때 두 컨테이너는 /dev/shm 에서 데이터를 고유하게 된다. 

### 생각할 문제( 초기 메모리 할당 설정의 어려움. 테라 단위로 메모리를 할당할 경우)
클라우드에서 사용할 경우는 메모리가 제한적일 경우, 대용량 메모리를 할당 해야하는 어려움이 있다. 이럴경우 어떻게 해야하는지 파악해야 한다. 

swap 영역으로 넘어갈 경우에 대한 스터디도 필요하다.

### 개발에 참고할 내용
기본적으로 휘발성이라는 특징있으며 공유메모리라는 특징을 가지고 있다. 따라서 ㄷ데이터 접근 속도가 빠르고, 컨테이너 상황에서 재사용시 데이터를 쉽게 제거 할 수 있다라는 장점을 가진다.(테스트 해봐야 함.)

유전체 분석시, 레퍼런스 파일이라던지 공용적으로 사용하는 대용량의 파일을 위치시키는 용도로 활용할 생각이다. 

### 참고 링크
[참고 1](https://blog.naver.com/ncloud24/221387977381)
[참고 2](https://github.com/containers/podman/issues/8181)