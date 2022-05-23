## container stauts 확인하기
https://docs.podman.io/en/latest/Reference.html

Podman v2.0 RESTful API 에 집중한다.
Podman v2.0 RESTful API 는 크게 Docker 와 호완이 가능한 api 와 Pod 와 같은 podman 의 고유한 기능을 지원하는 Libpod API 로 구성된다.


### binding 시 api 차원에서 확인해줘야 한다.

### Status 세부 내용
- https://blog.naver.com/alice_k106/221310477844

하지만 위의 status 와 다르게 별도의 healthcheck 기능이 제공된다. 이 부분은 api 설계 또는 이용할때 참고해야한다.
https://docs.podman.io/en/latest/_static/api.html?version=v4.1#operation/ContainerExportLibpod

이게 binding 관련 소스 인듯
https://github.com/containers/podman/blob/main/pkg/bindings/containers/healthcheck.go