package main

import (
	"fmt"
	"github.com/src-d/go-git/clients/http"
	git "gopkg.in/src-d/go-git.v2"
)

func main() {
	au := http.NewBasicAuth("user-build", ";msxasdsdFJ,4Cg]")
	r, err := git.NewAuthenticatedRemote("https://bitbucket.org/myuser/tools", au)
	if err != nil {
		panic(err)
	}
	if err = r.Connect(); err != nil {
		panic(err)
	}
	refs := r.Refs()
	for k, v := range refs {
		fmt.Printf("ref %s, sha %s\n", k, v)
	}
}
