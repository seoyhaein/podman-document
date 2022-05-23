### 여러 dockerfile/containerfile 사용법 (seoy)
```
docker build -t api-server:latest . -f dev.Dockerfile
docker build -t api-server:latest . -f stg.Dockerfile
docker build -t api-server:latest . -f prod.Dockerfile

docker build -t <도커파일명>:<태그> . -f <사용할 Dockerfile>

```
