package mongorapper

import (
	"context"
	"time"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoSession struct {
	client *mongo.Client
}

func MongoNew(param *config.Config) *MongoSession {
	mongoDns := config.MongoDns(&param.Mongo)
	clientOptions := options.Client().ApplyURI(mongoDns)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// defer client.Disconnect(context.TODO())
	if err != nil {
		mainGlobal.GVA_LOG.Error("Connected to MongoDB!!", zap.Error(err))
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		mainGlobal.GVA_LOG.Error("Ping MongoDB!!", zap.Error(err))
		return nil
	}

	session := MongoSession{
		client: client,
	}

	mainGlobal.GVA_LOG.Info("Connected to MongoDB!!")
	return &session
}

func (m *MongoSession) MongoInsert(mongoparm config.MongoParam, index_name string, document interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	collection := m.client.Database(mongoparm.DataBase).Collection(mongoparm.Collection)
	if index_name != "" {
		mod := mongo.IndexModel{
			Keys: bson.M{
				index_name: -1, // index in descending order
			},
			Options: options.Index().SetUnique(true),
		}
		collection.Indexes().CreateOne(ctx, mod)
	}

	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoSession) MongoGet(mongoparm config.MongoParam, limit int, offset int, filter interface{}) (list interface{}, total int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	collection := m.client.Database(mongoparm.DataBase).Collection(mongoparm.Collection)
	var findoptions options.FindOptions
	if limit > 0 {
		findoptions.SetLimit(int64(limit))
		findoptions.SetSkip(int64(offset))
	}
	total = 0
	cur, err := collection.Find(ctx, filter, &findoptions)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	var results []bson.M
	err = cur.All(ctx, &results)
	return results, total, err
}

// 关闭
func (m *MongoSession) Close() {
	err := m.client.Disconnect(context.TODO())
	if err != nil {
		mainGlobal.GVA_LOG.Error("Close MongoDB!!", zap.Error(err))
	}
	mainGlobal.GVA_LOG.Info("MongoDB closed!!")
}
