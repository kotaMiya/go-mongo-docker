package repository

import (
	"context"
	"go-mongo-docker/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Project repository functions
type ProjectRepository interface {
	GetProjects() ([]*entity.Project, error)
	CreateProject(*entity.Project) (*entity.Project, error)
	UpdateProject(*entity.Project, string) (*mongo.UpdateResult, error)
}

// Project repository structure has db
type projectRepository struct {
	db *mongo.Client
}

// NewProjectRepository returns "ProjectRepository"
func NewProjectRepository(db *mongo.Client) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

/**
 * @comment: returns all projects and todos which included in each project
 * @return : projects and todos
 */
func (p *projectRepository) GetProjects() ([]*entity.Project, error) {
	collection := p.db.Database("projects-db").Collection("projects")
	cur, err := collection.Find(context.Background(), bson.D{})
	avoidPanic(err)

	var results []*entity.Project

	for cur.Next(context.Background()) {
		var projcet *entity.Project
		err := cur.Decode(&projcet)
		avoidPanic(err)
		results = append(results, projcet)
	}

	return results, nil
}

// CreateProject() registers a project in db.
func (p *projectRepository) CreateProject(project *entity.Project) (*entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("projects-db").Collection("projects")
	_, err := collection.InsertOne(ctx, *project)
	avoidPanic(err)

	return project, nil
}

// UpdateProject() updates data of a project.
func (p *projectRepository) UpdateProject(project *entity.Project, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("projects-db").Collection("projects")

	filter := bson.M{"id": convertToInt(id)}
	update := bson.M{
		"$set": bson.M{
			"name":        *&project.Name,
			"description": *&project.Description,
			"color":       *&project.Color,
		}}

	result, err := collection.UpdateOne(ctx, filter, update)
	avoidPanic(err)

	return result, nil
}
