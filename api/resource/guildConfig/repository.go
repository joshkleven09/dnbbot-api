package guildConfig

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	coll *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		coll: db.Collection("guild-config"),
	}
}

func (r *Repository) FindAll(externalGuildId string) (Models, error) {
	models := make([]*Model, 0)
	filter := bson.D{{}}

	if externalGuildId != "" {
		filter = bson.D{{"external_guild_id", externalGuildId}}
	}

	cursor, err := r.coll.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &models); err != nil {
		return nil, err
	}

	return models, nil
}

func (r *Repository) Create(model *Model) (int64, error) {
	filter := bson.D{{"external_guild_id", model.ExternalGuildId}}
	update := bson.D{{"$set", bson.D{
		{"guild_name", model.GuildName},
		{"default_channel", model.DefaultChannel},
		{"last_updated_at", time.Now()},
	}}}

	var findResult Model
	err := r.coll.FindOne(context.TODO(), filter).Decode(&findResult)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("here")
			_, err := r.coll.InsertOne(context.TODO(), model)
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
	}

	result, err := r.coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, nil
}

func (r *Repository) Delete(guildConfigId string) error {
	objID, err := primitive.ObjectIDFromHex(guildConfigId)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(context.TODO(), bson.D{{"_id", objID}})

	return err
}
