package mongodb

import (
	"api-service/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mgr *MongoDBManager) PostDoc(payload interface{}) error {
	loggerService := logger.GetLogger()
	collection := mgr.Instance
	opts := options.InsertOne()
	filter, err := collection.InsertOne(context.TODO(), payload, opts)
	if err != nil {
		loggerService.Errorln("[Create:substate] Error in update reason : ", err)
		return err
	}
	loggerService.Infoln("Successfully inserted doc ", filter.InsertedID)
	return nil
}
func (mgr *MongoDBManager) GetDoc(id string) (map[string]interface{}, error) {
	loggerService := logger.GetLogger()
	collection := mgr.Instance
	opts := options.FindOne()
	filter := collection.FindOne(context.TODO(), bson.M{"_id": id}, opts)
	payload := make(map[string]interface{})
	if err := filter.Decode(&payload); err != nil {
		return nil, err
	}
	loggerService.Infoln("Successfully Fetched data from DB")
	return payload, nil
}

func (mgr *MongoDBManager) PutDoc(payload interface{}, id string) error {
	loggerService := logger.GetLogger()
	collection := mgr.Instance
	opts := options.Update()
	opts.SetUpsert(true)
	filter, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": &payload}, opts)
	if err != nil {
		loggerService.Errorln("Error in update reason : ", err)
		return err
	}
	if filter.MatchedCount == 0 {
		loggerService.Infoln("document was  Added ")
	}
	if filter.MatchedCount == 1 {
		loggerService.Infoln("document was  updated  ")
	}
	return nil
}

func (mgr *MongoDBManager) DeleteDoc(id string) error {
	loggerService := logger.GetLogger()
	collection := mgr.Instance
	opts := options.Delete()
	filter, err := collection.DeleteOne(context.Background(), bson.M{"_id": id}, opts)
	if err != nil {
		loggerService.Errorln("Error in update reason : ", err)
		return err
	}
	if filter.DeletedCount >= 1 {
		loggerService.Infoln("Successfully deleted the document")
	}
	return nil
}
