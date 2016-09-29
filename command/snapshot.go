package command

import (
	"fmt"
	"github.com/pshima/consul-snapshot/backup"
	"github.com/pshima/consul-snapshot/config"
	"github.com/pshima/consul-snapshot/consul"
	"path/filepath"
)

//SnapshotCommand for tacking backups to file
type SnapshotCommand struct {
	Meta
	Version string
}

//Run Get one time backup to file
func (c *SnapshotCommand) Run(args []string) int {
	c.UI.Info(fmt.Sprintf("v%v: Starting Consul Snapshot...", c.Version))
	c.UI.Info(fmt.Sprintf("Passed args are: %v", args))
	if len(args) < 1 {
		c.UI.Info("Please specify folder tio backup")
		return 1
	}
	destination := args[0]
	conf := &config.Config{Version: c.Version,
		TmpDir:       filepath.Dir(destination),
		SkipRemote:   true,
		BackupToFile: filepath.Base(destination)}
	client := &consul.Consul{Client: *consul.Client()}
	backup.DoWork(conf, client)
	return 0
}

// Synopsis of the command
func (c *SnapshotCommand) Synopsis() string {
	return "Makes single backup to file"
}

// Help for the command
func (c *SnapshotCommand) Help() string {
	return `
Usage: consul-snapshot snapshot path/to/snaphot.backup

Makes one shot backup to local folder.
`
}
