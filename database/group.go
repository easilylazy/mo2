package database

import (
	"context"
	"github.com/Monkey-Mouse/mo2/mo2utils/mo2errors"
	"github.com/Monkey-Mouse/mo2/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const groupStr = "group"

var GroupCol = GetCollection(groupStr)

// UpsertGroup upsert group
func UpsertGroup(group model.Group) mo2errors.Mo2Errors {
	if _, err := GroupCol.UpdateOne(context.TODO(), bson.M{"_id": group.ID}, bson.M{"$set": bson.M{
		"owner_id":       group.OwnerID,
		"access_manager": group.AccessManager,
	}}, options.Update().SetUpsert(true)); err != nil {
		return mo2errors.InitError(err)
	}
	return mo2errors.InitNoError("upsert success")
}

// FindGroup find group
func FindGroup(id primitive.ObjectID) (group model.Group, mErr mo2errors.Mo2Errors) {
	if err := GroupCol.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			mErr.InitCode(mo2errors.Mo2NotFound)
			return
		}
		mErr.InitError(err)
		return
	}
	mErr.InitNoError("find one")
	return
}

// DeleteGroupByID  delete group by id
func DeleteGroupByID(id primitive.ObjectID) mo2errors.Mo2Errors {
	if _, err := GroupCol.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return mo2errors.InitError(err)
	}
	return mo2errors.InitNoError("delete success")
}

// DeleteGroupByOwnerID  delete groups by owner id
func DeleteGroupByOwnerID(id primitive.ObjectID) mo2errors.Mo2Errors {
	if res, err := GroupCol.DeleteMany(context.TODO(), bson.M{"owner_id": id}); err != nil {
		return mo2errors.InitError(err)
	} else {
		log.Println(res)
		return mo2errors.InitNoError("delete success %v", res)
	}
}
