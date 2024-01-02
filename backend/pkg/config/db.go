package config

import (
	"github.com/souravdev-eng/toktok-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	MONGO_URI := utils.GetEnvValue("MONGO_URI")

	clientOption := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(nil, clientOption)

	if err != nil {
		return nil, err
	}

	return client, nil
}
