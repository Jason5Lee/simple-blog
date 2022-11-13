package repository

import (
	"context"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/errors"
	"github.com/Jason5Lee/simple-blog/core/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Data for inserting into MongoDB.
type DBArticleInfo struct {
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Author  string `bson:"author"`
}

// Data for reading from MongoDB, with extra field "_id".
type DBArticle struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	Author  string             `bson:"author"`
}

// ArticleRepositoryMongoDB is a MongoDB implementation of ArticleRepository.
type ArticleRepositoryMongoDB struct {
	client *mongo.Client
}

const dbName = "simple-blog"
const collectionName = "articles"

// NewArticleRepositoryMongoDB creates a new ArticleRepositoryMongoDB connecting to a MongoDB.
func NewArticleRepositoryMongoDB(mongoUri string) (*ArticleRepositoryMongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}
	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return &ArticleRepositoryMongoDB{client: client}, nil
}

func (repo *ArticleRepositoryMongoDB) Create(ctx context.Context, article *data.ArticleInfo) (data.ArticleID, error) {
	insertResult, err := repo.client.Database(dbName).Collection(collectionName).InsertOne(ctx, DBArticleInfo{
		Title:   string(article.Title),
		Content: string(article.Content),
		Author:  string(article.Author),
	})
	if err != nil {
		return "", err
	}
	return data.ArticleID(insertResult.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (repo *ArticleRepositoryMongoDB) GetByID(ctx context.Context, id data.ArticleID) (*data.Article, error) {
	docID, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		// Invalid ID does not match any document, so we return ErrNotFound.
		return nil, errors.ErrNotFound
	}
	findResult := repo.client.Database(dbName).Collection(collectionName).FindOne(ctx, map[string]interface{}{"_id": docID})
	if err := findResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	var article DBArticle
	if err := findResult.Decode(&article); err != nil {
		return nil, err
	}
	return &data.Article{
		ID: id,
		// Assume the data in MongoDB is valid.
		ArticleInfo: data.ArticleInfo{
			Title:   data.ArticleTitle(article.Title),
			Content: data.ArticleContent(article.Content),
			Author:  data.ArticleAuthor(article.Author),
		},
	}, nil
}

func (repo *ArticleRepositoryMongoDB) GetAll(ctx context.Context) ([]*data.Article, error) {
	cursor, err := repo.client.Database(dbName).Collection(collectionName).Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	var articles []*DBArticle
	if err := cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	result := make([]*data.Article, len(articles))
	for i, article := range articles {
		result[i] = &data.Article{
			ID: data.ArticleID(article.ID.Hex()),
			// Assume the data in MongoDB is valid.
			ArticleInfo: data.ArticleInfo{
				Title:   data.ArticleTitle(article.Title),
				Content: data.ArticleContent(article.Content),
				Author:  data.ArticleAuthor(article.Author),
			},
		}
	}
	return result, nil
}

// Dropping the collection for integration testing.
func (repo *ArticleRepositoryMongoDB) Drop() error {
	return repo.client.Database(dbName).Collection(collectionName).Drop(context.Background())
}

func (repo *ArticleRepositoryMongoDB) Close() error {
	return repo.client.Disconnect(context.Background())
}

var _ repository.ArticleRepository = (*ArticleRepositoryMongoDB)(nil)
