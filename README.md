## IP 地址库

根据仓库[ip2region](https://github.com/lionsoul2014/ip2region)做了个api

```
docker build -t ip-parse .
```

```
docker run -it \
-p 8083:8083 \
--restart=always \
-d \
--name ip-parse \
ip-parse
```