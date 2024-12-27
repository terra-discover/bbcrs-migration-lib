package main

const (
	myPackage               = "lab.tog.co.id/bb/migration"
	modelMigrationsFilePath = "model_migrations.go"
)

func getListGenReq() (listGenReq ListGenModelMigrationRequest) {
	listGenReq = ListGenModelMigrationRequest{
		GenModelMigrationRequest{
			ModelsFilePath: "model",
			ImportAlias:    "model",
			ContentHeader:  "Model Section",
		},
	}

	return
}
