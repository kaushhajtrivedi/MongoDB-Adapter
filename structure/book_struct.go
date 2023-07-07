package structure
 import ("go.mongodb.org/mongo-driver/bson/primitive")

type Book struct {

	Age int32

	ID primitive.ObjectID `bson:"_id"`

	Name string

}
