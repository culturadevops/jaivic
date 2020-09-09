package services

func (t *Srv) CreateMidDefault(DestFileName string) {
	t.CreateDefaultGoFile("middlewares", "default", "middlewares", DestFileName)
	t.AddMidToMain(DestFileName)
}
func (t *Srv) CreateMidDefaultToRoutes(DestFileName string, DestRoutes string) {
	t.CreateDefaultGoFile("middlewares", "default", "middlewares", DestFileName)
	t.AddMidToRts(DestFileName, "routes/"+DestRoutes)
}
