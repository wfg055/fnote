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

package repository

import (
	"context"

	"github.com/chenmingyong0423/fnote/backend/internal/config/repository/dao"
	"github.com/pkg/errors"
)

type IConfigRepository interface {
	FindByTyp(ctx context.Context, typ string) (any, error)
	Increase(ctx context.Context, field string) error
}

func NewConfigRepository(dao dao.IConfigDao) *ConfigRepository {
	return &ConfigRepository{
		dao: dao,
	}
}

var _ IConfigRepository = (*ConfigRepository)(nil)

type ConfigRepository struct {
	dao dao.IConfigDao
}

func (r *ConfigRepository) Increase(ctx context.Context, field string) error {
	return r.dao.Increase(ctx, field)
}

func (r *ConfigRepository) FindByTyp(ctx context.Context, typ string) (any, error) {
	config, err := r.dao.FindByTyp(ctx, typ)
	if err != nil {
		return nil, errors.WithMessage(err, "r.dao.FindByTyp failed")
	}
	return config.Props, nil
}
