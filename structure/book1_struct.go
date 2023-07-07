package structure
 import ("go.mongodb.org/mongo-driver/bson/primitive")

type Book1 struct {

	ID primitive.ObjectID `bson:"_id"`

}
