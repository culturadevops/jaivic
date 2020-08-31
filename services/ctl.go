package services

import (
	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

var VarCtl *Ctl

type Ctl struct {
}

func (t *Ctl) CreateMap(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = fileName
	return MapForReplace
}

//services.Mdl(args[0])
func (t *Ctl) Create(fileName string, version string) {
	FolderBase := "controllers"
	jio.CreateFolder(FolderBase)
	MapForReplace := t.CreateMap(fileName)
	newName := FolderBase + "/" + fileName + ".go"

	jio.NewFileforTemplate(newName, FolderBase+viper.GetString("default.templatefolder")+version+".stub", MapForReplace)

}
