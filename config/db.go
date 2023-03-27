package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}

func DBinstance() (client *mongo.Client) {

	user := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	conn := fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, pass, host, port)

	replicaSetQueryString := "replicaSet=replica-set"
	tlsQueryString := ""
	//var tlsConfig *tls.Config

	if os.Getenv("GIN_MODE") == "release" {
		replicaSetQueryString = "replicaSet=rs0"
		//tlsQueryString = "&tls=true"

		//// configure tls
		//var filename = "rds-combined-ca-bundle.pem"
		//tlsConfig := new(tls.Config)
		//certs, err := ioutil.ReadFile(filename)

		//if err != nil {
		//fmt.Println("Failed to read CA file")
		//return
		//}

		//tlsConfig.RootCAs = x509.NewCertPool()
		//ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

		//if !ok {
		//fmt.Println("Failed to append CA file")
		//return
		//}

	}

	conn = fmt.Sprintf("%s?%s%s", conn, replicaSetQueryString, tlsQueryString)

	fmt.Printf("Attempting connection with: %s", conn)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPIOptions)
	//if tlsConfig != nil {
	//fmt.Println("Successfully set TLS config")
	//clientOptions.SetTLSConfig(tlsConfig)
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Connecting to MongoDB ...")
	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
	fmt.Println("Initialising indexes ...")
	// initialise indexes
	InitIndexes(client)
	fmt.Println("Success!")
	return client
}

func InitIndexes(client *mongo.Client) {

	// transactions_transactions_-1 index
	transactionCollection := OpenCollection(client, "transactions")

	transactionIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "transaction_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}
	transactionIndexCreated, err := transactionCollection.Indexes().CreateOne(context.Background(), transactionIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	// campaigns_campaigns_-1 index
	campaignCollection := OpenCollection(client, "campaigns")

	campaignIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "campaign_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	campaignIndexCreated, err := campaignCollection.Indexes().CreateOne(context.Background(), campaignIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	// cards_cards_-1 index
	cardCollection := OpenCollection(client, "cards")

	cardIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "card_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	cardIndexCreated, err := cardCollection.Indexes().CreateOne(context.Background(), cardIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created Transaction Index %s\n", transactionIndexCreated)
	fmt.Printf("Created Campaign Index %s\n", campaignIndexCreated)
	fmt.Printf("Created Card Index %s\n", cardIndexCreated)
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("loyalty").Collection(collectionName)

	return collection
}
