package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Setup mongo 클라이언트를 초기화한다.
func Setup(ctx context.Context, address string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//cancel은 설정이 종료될 때 호출된다.(지연 호출)
	defer cancel()

	/*mongo.NewClient is deprecated
	client, err := mongo.NewClient(options.Client().ApplyURI(address))*/
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	/*client.Connect(ctx) is deprecated
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	*/
	return client, nil
}
