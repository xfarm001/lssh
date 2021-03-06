package main

import (
	"fmt"
	"os"
	"os/user"
	"sort"

	arg "github.com/alexflint/go-arg"
	"github.com/blacknon/lssh/check"
	"github.com/blacknon/lssh/conf"
	"github.com/blacknon/lssh/list"
	"github.com/blacknon/lssh/ssh"
)

type CommandOption struct {
	FilePath string `arg:"-f,help:config file path"`
	Exec     string `arg:"-e,help:exec_command"`
}

func main() {
	// Exec Before Check
	check.OsCheck()
	check.DefCommandExistCheck()

	// Set default value
	usr, _ := user.Current()
	defaultConfPath := usr.HomeDir + "/.lssh.conf"

	// get Command Option
	var args struct {
		CommandOption
	}

	// Default Value
	args.FilePath = defaultConfPath
	args.Exec = ""
	arg.MustParse(&args)

	// set option value
	configFile := args.FilePath
	execRemoteCmd := args.Exec

	// Get List
	listConf := conf.ConfigCheckRead(configFile)

	// Get Server Name List (and sort List)
	nameList := conf.GetNameList(listConf)
	sort.Strings(nameList)

	// View List And Get Select Line
	selectServer := list.DrawList(nameList, listConf)

	if selectServer == "ServerName" {
		fmt.Println("Server not selected.")
		os.Exit(0)
	}

	// Connect SSH
	ssh.ConnectSsh(selectServer, listConf, execRemoteCmd)
}
