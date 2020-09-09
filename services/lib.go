package services

func (t *Srv) CreateLibDefault(DestFileName string) {
	t.CreateDefaultGoFile("libs", "default", "libs", DestFileName)
}
