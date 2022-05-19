package services

import (
	"mysqlbinlogparser/configs"
	"mysqlbinlogparser/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var SellsPerDayCollection *mongo.Collection = configs.GetCollection(configs.DB, "SellsPerDay")

//gets difference
func GetDiff() (*models.Difference, error) {
	return nil, nil
}

//Funcion que obtiene de la cache o en caso de no ser posible, del microservicio catalog, los datos de un cierto producto
// a partir de su ID.
