package mongo

import (
    "context"
    "time"

    "abc-user-backend/models"
    "abc-user-backend/repositories"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
    collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) repositories.UserRepository {
    return &userRepo{
        collection: db.Collection("users"),
    }
}

func (r *userRepo) GetAll(search string, limit, offset int) ([]models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{}
    if search != "" {
        filter = bson.M{
            "$or": []bson.M{
                {"name": bson.M{"$regex": search, "$options": "i"}},
                {"email": bson.M{"$regex": search, "$options": "i"}},
            },
        }
    }

    findOptions := options.Find()
    findOptions.SetLimit(int64(limit))
    findOptions.SetSkip(int64(offset))
    findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

    cursor, err := r.collection.Find(ctx, filter, findOptions)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []models.User
    if err := cursor.All(ctx, &users); err != nil {
        return nil, err
    }

    return users, nil
}


func (r *userRepo) GetByID(id string) (*models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepo) Create(user *models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user.ID = uuid.New().String()
    _, err := r.collection.InsertOne(ctx, user)
    return err
}

func (r *userRepo) Update(user *models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.UpdateOne(ctx,
        bson.M{"_id": user.ID},
        bson.M{"$set": user},
    )
    return err
}

func (r *userRepo) Delete(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

func (r *userRepo) IsEmailUnique(email string, excludeID string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"email": email}
    if excludeID != "" {
        filter["_id"] = bson.M{"$ne": excludeID}
    }

    count, err := r.collection.CountDocuments(ctx, filter)
    if err != nil {
        return false, err
    }

    return count == 0, nil
}

func (r *userRepo) Count(search string) (int64, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{}
    if search != "" {
        filter = bson.M{
            "$or": []bson.M{
                {"name": bson.M{"$regex": search, "$options": "i"}},
                {"email": bson.M{"$regex": search, "$options": "i"}},
            },
        }
    }

    count, err := r.collection.CountDocuments(ctx, filter)
    if err != nil {
        return 0, err
    }

    return count, nil
}

