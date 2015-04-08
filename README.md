# Blackhole

Blackhole is a cleaner that remove untagged docker images. It makes your docker images are always clean without untagged images.

###Environment variable

* SLEEP_TIME(minute)

If you don't set SLEEP_TIME, default is 60(=1hour)

###Run

```
$ docker run -d --name <container name> -e SLEEP_TIME=<minute> -v /var/run/docker.sock:/var/run/docker.sock dongjujang/blackhole
```
