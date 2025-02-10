# my-ddns
use golang implement tencent domain ddns 

## config file
> Create a configuration file`config.yaml` and place it in the same directory as the program.
```yaml
app:
  sid: tencent SecretId
  skey: tencent SecretKey
  domain: you domain
  target: modify domain of prefix
```

## Future
- `only support tencent domain`
- Can automatically obtain the public network IP address of this machine and regularly modify the specified domain name.

## compile
> Compile to linux executable program ： open Terminal
```bash
set GOARCH=amd64
set GOOS=linux
go build -o myddns main.go
```

## build docker image
>` docker build -t myddns:0.1 .`
```bash
kingschan@server:/home/kingschan/myddns$ ls
config.yaml  Dockerfile  myddns

kingschan@server:/home/kingschan/myddns$ docker build -t myddns:0.1 .
[+] Building 16.0s (12/12) FINISHED                                                                                                                                                                                                                              
 => [internal] load .dockerignore                                                                                                                                                                                                                           0.0s
 => => transferring context: 2B                                                                                                                                                                                                                             0.0s
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                        0.0s
 => => transferring dockerfile: 437B                                                                                                                                                                                                                        0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                                                                                           15.3s
 => [1/7] FROM docker.io/library/alpine@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300                                                                                                                                             0.0s
 => [internal] load build context                                                                                                                                                                                                                           0.0s
 => => transferring context: 61B                                                                                                                                                                                                                            0.0s
 => CACHED [2/7] RUN echo ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone                                                                                                                                                        0.0s
 => [3/7] RUN mkdir -p /opt/myddns                                                                                                                                                                                                                          0.2s
 => [4/7] ADD  myddns /opt/myddns/myddns                                                                                                                                                                                                                    0.0s
 => [5/7] ADD  config.yaml /opt/myddns/config.yaml                                                                                                                                                                                                          0.0s
 => [6/7] WORKDIR /opt/myddns                                                                                                                                                                                                                               0.0s
 => [7/7] RUN chmod -R 755 /opt/myddns                                                                                                                                                                                                                      0.3s
 => exporting to image                                                                                                                                                                                                                                      0.1s
 => => exporting layers                                                                                                                                                                                                                                     0.1s
 => => writing image sha256:0f4e0807b70ee7902b9eefced16f87c6ed6487f75aef2ec873ffe48ca4d1cba3                                                                                                                                                                0.0s
 => => naming to docker.io/library/myddns:0.1                                                                                                                                                                                                                                                  0.0s

```

## run in docker
> `docker run -d  --name myddns --restart=always  myddns:0.1`
> > job done !


## enable linux service 
```
cp ./myddns.service /etc/systemd/system
systemctl daemon-reload
systemctl start myddns
systemctl status myddns
systemctl enable myddns
```