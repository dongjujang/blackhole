# Blackhole

Blackhole is a cleaner that remove untagged docker images

###Environment variable

* SLEEP_TIME(minute)

If you don't set SLEEP_TIME, default is 60(=1hour)

###RUN

```
$ docker run -d --name <container name> -e SLEEP_TIME=<minute> -v /var/run/docker.sock:/var/run/docker.sock dongjujang/blackhole
```
