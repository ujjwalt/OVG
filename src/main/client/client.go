/*
   Format of the json file used for saving configuration is as follows :-
   {
       "auth": "username:password",
       "projects" : [
           {
               "id": "project id"
               "folder": "path/to/the/folder/containing/files"
               "cpu": cycles_spent on this project
           }
       ]
   }

   if this is the first time then create the file else open it and load the values as a map. If username and password are passed then
   overwrite the stored values and use them.
*/
package main

/*
#include <unistd.h>
*/

import (
	// "crypto/md5"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os/exec"
	// "../ovg"
	"github.com/cloudfoundry/gosigar"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"syscall"
	"time"
)

var (
	uname          = flag.String("u", "username", "Enter your username")
	passwd         = flag.String("p", "password", "Enter your password")
	config         configT  // memory model of config.json
	configF        *os.File // file pointer to config.json
	currentUser, _ = user.Current()
	libPath        = currentUser.HomeDir + "/Library/OVG/" // path to the user's library
	configPath     = libPath + "config.json"               // path to the config file
)

const (
	workURL = "http://api.domain.com/work"
)

type projectT struct {
	id     string
	folder string
	cpu    int64
}

type configT struct {
	auth     string
	Projects []projectT
}

func main() {
	// Parse the username and password from the flags
	flag.Parse()
	if *uname == "" || *passwd == "" {
		log.Fatal("Invalid username or password: ", *uname)
	}
	// create the lib directory
	fatal(os.MkdirAll(libPath, 0700))
	fatal(os.Chdir(libPath))
	var err error
	configF, err = os.OpenFile(configPath, os.O_RDWR, 0644) // only readable and writable by current user
	if err != nil {
		configF, err = os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0644)
		fatal(err)
	} else {
		readConfig()
	}
	defer configF.Close()
	// store username:password encoded as base64
	config.auth = *uname + ":" + *passwd
	config.auth = base64.StdEncoding.EncodeToString([]byte(config.auth))
	defer writeConfig() // write out whenever you quit
	// Request work every time the free RAM reaches >= 25% of total RAM.
	// Keep requesting work whenever you get a notification for the system to be idle and start it off
	ch := make(chan time.Time, 10)
	notifyIdleSystem(ch)
	for t := range ch {
		// request work from the server
		go requestWork()
	}
}

func requestWork() {
	client := &http.Client{}
	resp, err := client.Get(workURL)
	req, err := http.NewRequest("GET", workURL, nil)
	req.Header.Add("Authorization", `Basic `+config.auth)
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	json.Unmarshal(body, m)
	files := m["files"]
	fileNames := m["fileNames"]
	var mainFile *os.File
	var mainFilePath string
	for i, f := range files {
		fh, err := os.Open(path.Join(os.TempDir(), m["id"], fileNames[i])) // tmp_dir/project_id/filename
		fatal(err)
		defer fh.Close()
		b, err := base64.StdEncoding.DecodeString(f)
		fh.Write(b)
		if fileNames[i] == "main" {
			mainFile = fh
			mainFilePath = path.Join(os.TempDir(), m["id"], "main")
		}
	}
	// fork off "main" and run it in a sandbox
	if pid := int(C.fork()); !pid {
		// child process
	}
}

func notifyIdleSystem(ch chan time.Time) {
	t := time.Tick(30 * time.Second)
	mem := new(sigar.Mem)
	for t1 := range t {
		mem.Get()
		if mem.Free >= 0.25*mem.Total {
			ch <- t1
		}
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeConfig() {
	b, err := json.Marshal(config)
	fatal(err)
	fatal(configF.Truncate(0))
	_, err = configF.WriteAt(b, 0)
	fatal(err)
	fatal(configF.Sync()) // write to stable storage
}

func readConfig() {
	info, err := configF.Stat()
	fatal(err)
	b := make([]byte, info.Size())
	_, err = configF.Read(b)
	fatal(err)
	fatal(json.Unmarshal(b, &config))
}
