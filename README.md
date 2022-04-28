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
- [podman bindings](https://podman.io/blogs/2020/08/10/podman-go-bindings.html)
- [podman trouble shooting](https://github.com/containers/podman/blob/main/troubleshooting.md)
- [podman Demo](https://github.com/containers/Demos)

### TODO
- 4.x 와 3.x 코드 비교해서 4.x 코드 중심으로 bindins 관련 학습한다. 소스가 업데이트 됨으로 향후 4.x 가 apt install podman  으로 4.x 도 설치 가능하리라 예상.
- 설치된 version 과 바인딩 코드와 버전 체크하는 api 제작 해야할듯.

