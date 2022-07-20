### 정리전 (서버용, 클라이언트용)

아래 링크에서 후반부에서 빌드 참고.(서버 설치)
- https://podman.io/getting-started/installation

빌드가 다 되면, 바이러니들을 /usr/local/bin 옮겨놓는다.

폴더에 대한 설명: https://jacking75.github.io/OS_linux_dir_kind/
https://wookiist.dev/10

바인딩할려면(바인딩도 결국 클라이언트이니까???) 소켓을 활성화 시켜야 한다.(클라이언트 설치)
아래 링크에서,
"systemctl --user enable --now podman.socket"
이 부분을 참고하자. rootless 임.

- https://github.com/containers/podman/blob/main/docs/tutorials/remote_client.md
- https://github.com/containers/podman/tree/main/pkg/bindings

그런다음에
systemctl start podman.socket (rootful)

또는 

systemctl start --user podman.socket (rootless)

기타, 3.x 설치후 여러 경험을 지운 후 문제점들.
apt purge 로 지워서 세팅들을 모두 날려야 한다. 안그럼 예전게 남아 에러 발생

User-selected graph driver "overlay" overwritten by graph driver "vfs" from database - delete libpod local files to resolve
발생시

https://github.com/containers/podman/issues/5114
그냥 db 를 지우거나 세팅을 지우면 됨.

- sudo rm -rf ~/.local/share/containers/
- rm /var/lib/containers/storage/libpod/bolt_state.db

### podman 상태
- podman --remote info
- https://www.redhat.com/sysadmin/container-information-podman

### etc
직접적인 예제 코드는 없고 일단 테스트 코드를 살펴보자 - 3.x 에는 해당 기능이 없다. 최신 버번에서 확인되는데 4.x 를 설치해야하는가???
https://sourcegraph.com/github.com/containers/podman/-/blob/test/e2e/pod_create_test.go?L14

### 하 정말.. 날려버린 12시간...
- https://docs.podman.io/en/latest/_static/api.html?version=v4.1
- podman system service -t 5000 & 

```
podman system service --help
Run API service

Description:
  Run an API service

Enable a listening service for API access to Podman commands.


Usage:
  podman system service [options] [URI]

Examples:
  podman system service --time=0 unix:///tmp/podman.sock
  podman system service --time=0 tcp://localhost:8888

Options:
      --cors string   Set CORS Headers
  -t, --time uint     Time until the service session expires in seconds.  Use 0 to disable the timeout (default 5)

```

