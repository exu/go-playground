package main

import (
	"golang.org/x/exp/inotify"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func valid(fileName string) bool {
	matched, err := regexp.MatchString(".+Test.php", fileName)

	log.Println(matched, err)

	return matched
}

type PhpUnit struct {
}

func (u *PhpUnit) runPhpunit(fileName string) {
	proc := exec.Command("php", "bin/phpunit-lts.phar", "-c", "app/", fileName)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	proc.Start()

	err := proc.Wait()
	ps := proc.ProcessState

	if err != nil {
		log.Printf("pid %d failed with %s", ps.Pid(), ps.String())
		os.Exit(1)
	} else {
		log.Printf("Phpunit completed")
	}
}

func main() {
	path := os.Args[1]
	log.Println("tests path", path)

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// @todo find way to do it recursively
	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}

	pu := &PhpUnit{}

	for {
		select {
		case ev := <-watcher.Event:
			modified := inotify.IN_MODIFY
			if ev.Mask&modified == modified && valid(ev.Name) {
				pu.runPhpunit(ev.Name)
			} else {
				log.Println("event:", ev)
			}

		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
