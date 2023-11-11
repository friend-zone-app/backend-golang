package scalar

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MarshalMyCustomObjectIDScalar marshals primitive.ObjectID to a GraphQL scalar
func MarshalMyCustomObjectIDScalar(objectID primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		idString := objectID.Hex()
		w.Write([]byte(strconv.Quote(idString)))
	})
}

// UnmarshalMyCustomObjectIDScalar unmarshals GraphQL scalar to primitive.ObjectID
func UnmarshalMyCustomObjectIDScalar(v interface{}) (primitive.ObjectID, error) {
	switch v := v.(type) {
	case string:
		objectID, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return primitive.ObjectID{}, err
		}
		return objectID, nil
	default:
		return primitive.ObjectID{}, fmt.Errorf("%T is not a string (hex representation of ObjectID)", v)
	}
}