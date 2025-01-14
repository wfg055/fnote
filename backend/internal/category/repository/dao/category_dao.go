// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"context"

	"github.com/chenmingyong0423/go-mongox"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category struct {
	Id         string   `bson:"_id"`
	Name       string   `bson:"name"`
	Route      string   `bson:"route"`
	Tags       []string `bson:"tags"`
	CreateTime int64    `bson:"create_time"`
	UpdateTime int64    `bson:"update_time"`
}

type ICategoryDao interface {
	GetAll(ctx context.Context) ([]*Category, error)
	GetCategoryByName(ctx context.Context, name string) (*Category, error)
}

var _ ICategoryDao = (*CategoryDao)(nil)

func NewCategoryDao(db *mongo.Database) *CategoryDao {
	return &CategoryDao{
		coll: mongox.NewCollection[Category](db.Collection("categories")),
	}
}

type CategoryDao struct {
	coll *mongox.Collection[Category]
}

func (d *CategoryDao) GetCategoryByName(ctx context.Context, name string) (*Category, error) {
	category, err := d.coll.Finder().Filter(bson.M{"name": name}).FindOne(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Find a category failed, [name=%s]", name)
	}
	return category, nil
}

func (d *CategoryDao) GetAll(ctx context.Context) ([]*Category, error) {
	result, err := d.coll.Finder().Find(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Find all categories failed failed")
	}
	return result, nil
}
