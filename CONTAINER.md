## podman pod create --share
### podman 의 pod 에서 namespaces, ipc, net, pid, user, 그리고 uts 가 shared / isolated 된다는 의미?

[참고 - podman pod create 옵션](https://github.com/containers/podman/blob/main/docs/source/markdown/podman-pod-create.1.md)

- --share 옵션 에 대한 설명

```
####--share=namespace A comma delimited list of kernel namespaces to share. 
If none or "" is specified, no namespaces will be shared. The namespaces to choose from are ipc, net, pid, user, uts.

--Option to specify the namespace to share in the pod. 
--Specify multiple comma-separated lists. 
--None or nothing is shared if an empty set is specified. 
--Can be specified from 5 namespaces: ipc, net, ipd, user, uts.

```

- 네임스페이스에 대해서
https://www.44bits.io/ko/keyword/linux-namespace

- network namespace [참고](https://url.kr/h2bm1f) 
- ip netns 네트워크 네임스페이스 확인하기, 처음에 해당 명령어를 입력했을때는 아무것도 나오지 않을 것이다. 
- First, check the network namespace that exists on the OS with the following command. Nothing is output. 
- It means that the network namespace has not been created yet.


* 두가지를 생각해보자, 메모리를 공유하는 방안과 volume 을 공유하는 방안을 생각해보자.


Using gRPC for (local) inter-process communication
-- https://docs.microsoft.com/ko-kr/aspnet/core/grpc/interprocess?view=aspnetcore-6.0

#### 일단 이부분은 추후에 살펴보는데 podman binding 콛에서 소켓 받는 부분을 참고해도 도움될듯하다. podman binding 이 경우는 소켓연결한후 restful 방식으로 진행하는데
#### 여기서 소켓 받는 부분을 참고 하고 그 후 grpc 로 접근하는 방법을 생각해봐도 될듯하다. 단순히 아이디어 차원에서 가능할지 모르겠다.

https://www.mpi-hd.mpg.de/personalhomes/fwerner/research/2021/09/grpc-for-ipc/


