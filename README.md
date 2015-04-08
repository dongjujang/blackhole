# blackhole

Remove untagged images

###ENV

SLEEP_TIME(minute)
if you don't set env, default is 60(=1hour)

    docker run -d --name <container name> -e SLEEP_TIME=<minute> -v /var/run/docker.sock:/var/run/docker.sock dongjujang/blackhole
