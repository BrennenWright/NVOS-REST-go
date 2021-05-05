// Description:
//   This script xports the current configuration of an NTO or GSC device
//   to a .ata file.
//   The script exports the configuration of several hosts simultaneously to
//   multiple files by creating one thread per host.

package main

import (
	"os"
	"strings"
	"time"

	nto "github.com/BrennenWright/NVOS-REST-go"
)

func main() {

	//argv = sys.argv[1:]
	var username = "admin"
	var password = "admin"
	//var host = ""
	//var hosts_file = ""
	//var config_file = ""
	var port = 8000
	path, _ := os.Getwd()

	var hosts = [2]string{"10.36.16.14", "10.36.16.100"}
	for _, host := range hosts {
		//run in a separate thread to max performance
		go func(target string) {
			//create a new nto system and export its full config
			os.MkdirAll(path+"/Backups", 0700)
			hostNPB := nto.New(target, username, password, port)
			hostNPB.ExportConfig("Full Config Scripted Backup", "FULL_BACKUP", path+"/Backups/"+strings.Replace(target, ".", "_", 3)+"_config"+time.Now().Format("20060102150405")+".ata")
		}(host)
	}

}
