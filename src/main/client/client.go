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
	"crypto/md5"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	// "path"
	"fmt"
)

var (
	uname   = flag.String("u", "username", "Enter your username")
	passwd  = flag.String("p", "password", "Enter your password")
	config  configT
	configF *os.File
)

const (
	configPath          = `~/Library/OVG/config.json`
	emptyConfigTemplate = `{
        "username" : "username",
        "password" : "password_as_md5",
        "projects" :
        [
            {
                "id"    : "project_id",
                "folder": "path/to/the/folder/containing/files",
                "cpu"   : cycles_spent on this project
            }
        ]
    }`
)

type projectT struct {
	id     string
	folder string
	cpu    int64
}

type configT struct {
	username string
	password string
	projects []projectT
}

func main() {
	// Parse the username and password from the flags
	flag.Parse()
	if *uname == "" || *passwd == "" {
		log.Fatal("Invalid username or password: ", *uname)
	}
	configF, err := os.Open(configPath) // only readable and writable by current user
	if err != nil {
		// file does not exist to create an empty one
		configF, err = os.Create(configPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	configF.WriteString(emptyConfigTemplate) // write out th empty template
	info, err := configF.Stat()
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, info.Size())
	configF.Read(b)
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}
	// store the username and password
	config.username = *uname
	h := md5.New() // convert the password to md5
	io.WriteString(h, *passwd)
	config.password = fmt.Sprintf("%x", h.Sum(nil))
	b, err = json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	_, err = configF.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
