package services

import (
	"github.com/culturadevops/jgt/jio"
)

var VarMdl *Mdl

type Mdl struct {
}

func (t *Mdl) CreateMap(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = fileName
	return MapForReplace
}
func (t *Mdl) GetModelTempleteName(version string) string {
	return "C:/Users/jaivi/.config/jaivic/models/" + version + ".stub"
}

//services.Mdl(args[0])
func (t *Mdl) Create(fileName string, version string) {
	FolderBase := "models"

	jio.CreateFolder(FolderBase)
	MapForReplace := t.CreateMap(fileName)
	newName := FolderBase + "/" + fileName + ".go"

	jio.NewFileforTemplate(newName, t.GetModelTempleteName(version), MapForReplace)

}
