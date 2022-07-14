## buildh 의 사용방법 정리.

[메뉴얼 참고](https://github.com/containers/buildah/blob/main/docs/tutorials/01-intro.md)

- 해당 메뉴얼은 fedor 계열을 중심으로 설명하고 있다. 향후 centos 8 이상으로 가거나 할때 필요할 수 는 있지만, 돈이 들어서? 생각해봐야 한다.
- ubuntu 기반으로 동일하게 작성함.
- 설치는 debootstrap 으로 진행함. [설치 메뉴얼](https://manpages.ubuntu.com/manpages/jammy/man8/debootstrap.8.html)
- 우분투 version 을 확인한다. [링크1](https://www.delftstack.com/ko/howto/linux/how-to-check-the-version-of-ubuntu/)
- 아래 코드는 debian bullseye 기준으로 설치 하는 설명 화면임.

```
# create an empty container with [scratch]
root@dlp:~# newcontainer=$(buildah from scratch)
root@dlp:~# buildah containers
CONTAINER ID  BUILDER  IMAGE ID     IMAGE NAME                       CONTAINER NAME
dd2673f1bad6     *     fe3c5de03486 docker.io/library/debian:latest  debian-working-container
6ad9282ae0cc     *     fe3c5de03486 docker.io/library/debian:latest  debian-working-container-1
d72663213be1     *                  scratch                          working-container

# mount [scratch] container
root@dlp:~# scratchmnt=$(buildah mount $newcontainer)
root@dlp:~# echo $scratchmnt
/var/lib/containers/storage/overlay/aa8ddf9a3fc5efa5653da2b5d091cfc21759e18c7db1945caf30ce42f98319f6/merged

# install packages to [scratch] container
root@dlp:~# apt -y install debootstrap
root@dlp:~# debootstrap bullseye $scratchmnt
# unmount
root@dlp:~# buildah umount $newcontainer
7dfacaa35c7523d2b82d6f957aad28a8c917ee4d95ebb1d6a1992843f4377788

# run container
root@dlp:~# buildah run $newcontainer bash
root@d72663213be1:/#
root@d72663213be1:/# cat /etc/os-release
PRETTY_NAME="Debian GNU/Linux 11 (bullseye)"
NAME="Debian GNU/Linux"
VERSION_ID="11"
VERSION="11 (bullseye)"
VERSION_CODENAME=bullseye
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"

root@d72663213be1:/# exit
# add images
root@dlp:~# buildah commit $newcontainer debian-basic:latest
Getting image source signatures
Copying blob ac373c890fc4 done
Copying config 3cd7af0ea4 done
Writing manifest to image destination
Storing signatures
3cd7af0ea41a46ab0cad5499fda0a20f3b89f4869774f0735c8b1e393e50b571

root@dlp:~# buildah images
REPOSITORY                 TAG      IMAGE ID       CREATED         SIZE
localhost/debian-basic     latest   3cd7af0ea41a   2 minutes ago   297 MB
localhost/my-debian        latest   7aa4a609ff69   2 hours ago     181 MB
docker.io/library/debian   latest   fe3c5de03486   11 days ago     129 MB

# test to run a container
root@dlp:~# podman run localhost/debian-basic /bin/echo "Hello my debian"
Hello my debian
```

- 지금 나의 현재 버번은 jammy 임으로 이 버전으로 진행한다.

- 현재, https://github.com/containers/buildah/blob/main/docs/tutorials/01-intro.md 에서 scratch 에서 bash 설치하고 진행하는 부분에서 막혔음.
- 조금 수정해서 alpine:latest 에서 bash 설치해서 진행하는 방향으로 접근하자.

### api 를 통해서 이미지 구성 중요
- https://github.com/containers/buildah/blob/main/docs/tutorials/04-include-in-your-build-tool.md

### install
[설치](https://github.com/containers/buildah/blob/main/install.md)

- podman 과 겹치는 모듈들이 발견이 되었다. runc 같은 경우는 일단 설치가 되어 있는데, 이것은 한번 buildah 설치 시 살펴보도록 하고, 세부적인 설치는 메뉴얼을 좀더 숙달해야 할 것 같다.



