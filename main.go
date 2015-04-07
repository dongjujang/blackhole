package main

import (
  "github.com/samalba/dockerclient"
  "log"
  "time"
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
  for true {
    RemoveImages()
    time.Sleep(1 * time.Hour)
  }
}

