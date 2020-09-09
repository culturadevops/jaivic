package services

func (t *Srv) CreateRtsDefault(DestFileName string) {
	t.CreateDefaultGoFile("routes", "default", "routes", DestFileName)
	t.AddRts(DestFileName)
}
func (t *Srv) CreateRtsApi(DestFileName string) {
	t.CreateDefaultGoFile("routes", "api", "routes", DestFileName)
	t.AddRts(DestFileName)
}
