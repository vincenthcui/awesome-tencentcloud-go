package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/go-git/go-git/v5"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

const (
	upstream = "https://github.com/TencentCloud/tencentcloud-cli.git"
)

var (
	workdir = flag.String("workdir", "tencentcloud/actions", "workdir to generate actions")
)

func genaction(ctx context.Context) error {
	defer wg.Done()

	tccli, err := os.MkdirTemp(os.TempDir(), "tccli")
	if err != nil {
		return err
	}
	defer func() {
		log.Println("clean up tccli directory")
		if err := os.RemoveAll(tccli); err != nil {
			log.Println("remove tccli directory failed: ", err.Error())
		}
	}()

	log.Println("cloning repository into ", tccli)
	if _, err = git.PlainCloneContext(ctx, tccli, false, &git.CloneOptions{URL: upstream, Progress: log.Writer()}); err != nil {
		return err
	}

	log.Println("walking into tccli")
	return filepath.WalkDir(path.Join(tccli, "tccli/services"), walks)
}

func walks(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !strings.HasSuffix(path, "api.json") {
		return nil
	}

	log.Println("processing ", path)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read file failed ", err.Error())
		return nil
	}

	mod := new(model)
	if err := json.Unmarshal(bytes, mod); err != nil {
		log.Println("unmarshal file failed: ", err.Error())
		return nil
	}

	ver := filepath.Join(*workdir, mod.Meta.ServiceShortName, mod.Meta.ApiVersion)
	if err = os.MkdirAll(ver, os.ModePerm); ignoreExists(err) != nil {
		log.Println("create directory failed", err.Error())
		return nil
	}

	f := file{
		service:   mod.Meta.ServiceShortName,
		instances: []instance{},
	}
	for actionName := range mod.Actions {
		f.instances = append(f.instances, instance{
			Service: mod.Meta.ServiceShortName,
			Version: mod.Meta.ApiVersion,
			Action:  actionName,
		})
	}

	sort.Sort(f.instances)
	return ioutil.WriteFile(filepath.Join(ver, "actions.go"), []byte(f.String()), os.ModePerm)
}

func ignoreExists(err error) error {
	if os.IsExist(err) {
		err = nil
	}
	return err
}
