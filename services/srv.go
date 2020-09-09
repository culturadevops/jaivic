package services

import (
	"strings"

	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

var VarSrv *Srv

type Srv struct {
}

func (t *Srv) aUpper(b string) string {
	b = strings.ToLower(b)
	b = strings.Title(b)
	return b
}
func (t *Srv) CreateMapDefault(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = strings.ToLower(fileName)
	MapForReplace["%EXPORTNAME%"] = t.aUpper(fileName)
	return MapForReplace
}
func (t *Srv) GetTempleteName(origFolderName string, filesVersionName string) string {
	return viper.GetString("homedir") + "/" + origFolderName + "/" + filesVersionName + ".stub"
}

func (t *Srv) CreateDefaultGoFile(origFolderName string, filesVersionName string, destFolderName string, destFileName string) {
	nameStruct := destFileName
	destFileName = destFileName + ".go"
	t.CreateDefaultFile(origFolderName, filesVersionName, destFolderName, nameStruct, destFileName)
}
func (t *Srv) CreateDefaultFile(origFolderName string, filesVersionName string, destFolderName string, nameStruct string, destFileName string) {
	jio.CreateFolder(origFolderName)
	MapForReplace := t.CreateMapDefault(nameStruct)
	newName := destFolderName + "/" + destFileName
	jio.NewFileforTemplate(newName, t.GetTempleteName(origFolderName, filesVersionName), MapForReplace)
}
func (t *Srv) CopyTemplate(origFolderName string, filesVersionName string, destFileName string) {
	jio.Copy(t.GetTempleteName(origFolderName, filesVersionName), destFileName)
}
func (t *Srv) UpdateFile(destFileName string, MapForReplace map[string]string) {
	jio.ChancarFile(destFileName, MapForReplace)
}
