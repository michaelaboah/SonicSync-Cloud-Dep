package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/michaelaboah/sonic-sync-cloud/graph/model"
	// "go.mongodb.org/mongo-driver/bson"
)

const ClientsDB = "clients-db"
const UserCol = "users"
const EquipDB = "equipment-inventory"
const ItemsCol = "items"

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {

  items := r.DB.Database(ClientsDB).Collection(UserCol)

  user := &model.User{
  	ID:    "0",
  	Name:  input.Name,
  	Email: input.Email,
  }	

  items.InsertOne(ctx, user)
  return user, nil
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput) (*model.Item, error) {
 
  var details model.CategoryDetails

  switch input.Category {
    case model.CategoryConsole: 
      details = input.Details.ConsoleInput 
  }


  item := &model.Item{
  	ID:           "",
  	CreatedAt:    input.CreatedAt,
  	UpdatedAt:    input.UpdatedAt,
  	Cost:         input.Cost,
  	Model:        input.Model,
  	Weight:       input.Weight,
  	Manufacturer: input.Manufacturer,
  	Category:     input.Category,
  	Details:      details,
  	Notes:        &input.Model,
  	Dimensions:   (*model.Dimension)(input.Dimensions),
  	PDFBlob:      input.PDFBlob,
  }

  r.items = append(r.items, item)
  return item, nil
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.ItemInput) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: UpdateItem - updateItem"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// FindByModel is the resolver for the find_by_model field.
func (r *queryResolver) FindByModel(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: FindByModel - find_by_model"))
}

// FindByID is the resolver for the find_by_id field.
func (r *queryResolver) FindByID(ctx context.Context) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: FindByID - find_by_id"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
