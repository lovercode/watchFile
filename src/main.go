package main

import (
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"log"
	"os/exec"
	"wathcher/src/conf"
)

const configName = "conf/conf.toml"

var watcher *fsnotify.Watcher

func init() {
	InitWatcher()

}

func InitWatcher() {
	log.Println(toml.DecodeFile("conf/conf.toml", &conf.Config))
	watcher, _ = fsnotify.NewWatcher()
	for k, _ := range conf.Config.Files {
		log.Println("add:", k, watcher.Add(k))
	}
	log.Println("add:", configName, watcher.Add(configName))
	log.Println("config is :", conf.Config)

}
func execCmd(file string) {
	cmd := exec.Command("/bin/bash", "-c", file)
	res, _ := cmd.Output()
	log.Println("run finish:", string(res))
}

func main() {
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					watcher.Add(event.Name)
				}
				if event.Name == configName {
					InitWatcher()
				} else {
					go func() {
						log.Println("file [", event.Name, "] op [", event.Name, "] , run :", conf.Config.Files[event.Name])
						for _, v := range conf.Config.Files[event.Name][event.Op.String()] {
							go execCmd(v)

						}
					}()

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-done

}
