/*
   Format of the json file used for saving configuration is as follows :-
   {
       "username" : "username",
       "password": "password_as_md5"
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

import (
	// "crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	// "../ovg"
	"github.com/cloudfoundry/gosigar"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
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

type projectT struct {
	id     string
	folder string
	cpu    int64
}

type configT struct {
	Username string
	Password string
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
	// store the username and password
	config.Username = *uname
	h := md5.New() // convert the password to md5
	io.WriteString(h, *passwd)
	config.Password = fmt.Sprintf("%x", h.Sum(nil))
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
	resp, err := http.Get("http://api.domain.com/work")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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
