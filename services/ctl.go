package services

//services.Mdl(args[0])
func (t *Srv) CreateCtlDefault(DestFileName string) {
	t.CreateDefaultGoFile("controllers", "default", "controllers", DestFileName)
}
func (t *Srv) CreateCtlCrud(DestFileName string) {
	t.CreateDefaultGoFile("controllers", "crud", "controllers", DestFileName)
}
