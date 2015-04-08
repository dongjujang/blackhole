package main

import (
  "github.com/samalba/dockerclient"
  "log"
  "time"
  "os"
  "strconv"
)

func RemoveImages() {
  docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
  if err != nil {
    log.Fatal(err)
  }

  images, err := docker.ListImages()
  if err != nil {
    log.Fatal(err)
  }
  
  for _, c := range images {
    tag := string(c.RepoTags[0])
    if (tag == "<none>:<none>") {
      id := string(c.Id[0:5])
      _, err := docker.RemoveImage(id)
      if err != nil {
        log.Fatal(err)
      }
      log.Println(c.Id + " deleted")
    }
  }
}

func main() {
  sleep_time_str := os.Getenv("SLEEP_TIME")
  var sleep_time int

  if sleep_time_str == "" {
    sleep_time = 60
  } else {
    sleep_time_int, err := strconv.Atoi(sleep_time_str)
    if err != nil {
      log.Fatal(err)
    }
    sleep_time = sleep_time_int
  }
  
  for true {
    RemoveImages()
    time.Sleep(time.Duration(sleep_time) * time.Minute)
  }
}

