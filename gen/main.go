package main

func main() {
	listGenReq := getListGenReq()
	genModelMigrations(listGenReq, myPackage, modelMigrationsFilePath)
}
