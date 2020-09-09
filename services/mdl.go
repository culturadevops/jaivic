package services

func (t *Srv) CreateMdlDefault(DestFileName string) {
	t.CreateMdl(DestFileName, "default")
}
func (t *Srv) CreateMdlCrud(DestFileName string) {
	t.CreateMdl(DestFileName, "crud")
}
func (t *Srv) CreateMdl(DestFileName string, verFileName string) {
	t.CreateDefaultGoFile("models", verFileName, "models", DestFileName)
}
