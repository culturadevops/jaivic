package services

import (

	"github.com/culturadevops/jgt/jio"

)

func (t *Srv) CopyInitConfig( Config string) {
	origFolderName := "init/"+Config
	t.CopyTemplate(origFolderName, "main", "main.go")
	origFolderName = origFolderName + "/libs"
	jio.CreateFolder("libs")
	t.CopyTemplate(origFolderName, "libs", "libs/libs.go")
	t.CopyTemplate(origFolderName, "mysql", "libs/mysql.go")
}