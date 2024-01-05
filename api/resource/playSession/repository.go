package playSession

import (
	"context"
	"dnbbot-api/api/resource"
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
		coll: db.Collection("play-session"),
	}
}

func (r *Repository) FindAll() (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)

	cursor, err := r.coll.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &playSessions); err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByGuildAndUserId(guildId string, userId string) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)

	cursor, err := r.coll.Find(context.TODO(), bson.D{{"guild_id", guildId}, {"user_id", userId}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &playSessions); err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByGuildId(guildId string, date string, timeFilterStart time.Time, timeFilterEnd time.Time) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)
	filter := bson.D{{}}

	if !time.Time.IsZero(timeFilterStart) && !time.Time.IsZero(timeFilterEnd) {
		filter = bson.D{
			{"guild_id", guildId},
			{"$and",
				bson.A{
					bson.D{{"start_time", bson.D{{"$gte", timeFilterStart}}}},
					bson.D{{"start_time", bson.D{{"$lte", timeFilterEnd}}}},
				},
			},
		}
	} else if date != "" {
		filter = bson.D{{"guild_id", guildId}, {"date", date}}
	} else {
		filter = bson.D{{"guild_id", guildId}}
	}

	cursor, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &playSessions); err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByUserId(userId string) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)

	cursor, err := r.coll.Find(context.TODO(), bson.D{{"user_id", userId}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &playSessions); err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) Create(playSession *PlaySession) (*PlaySession, error) {
	result, err := r.coll.InsertOne(context.TODO(), playSession)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, &resource.DuplicateError{Message: "Session already exists for user, guild, time range"}
		} else {
			return nil, err
		}
	}

	playSession.ID = result.InsertedID.(primitive.ObjectID)

	return playSession, nil
}

func (r *Repository) Delete(playSessionId string) error {
	objID, err := primitive.ObjectIDFromHex(playSessionId)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(context.TODO(), bson.D{{"_id", objID}})

	return err
}
