//server code to manage the project communication interface

package main

import (
   "crypto/md5"
   "flag"
   "fmt"
   "io"
   "log"
   "os"
   "os/user"
)

var (
   uname = flag.String("u", "username", "Enter your username")
   passwd = flag.String("p", "password", "Enter your password")
    currentUser, _ = user.Current()
   libPath = currentUser.HomeDir + "/Library/OVG/"
   configPath = libPath + "config.json"
   err error // generic error
)

type projectT struct {
   id string
   folder string
   cpu int64
}

