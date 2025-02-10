// Copyright 2025 Mykhailo Bobrovskyi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contract

import (
	"context"

	"mbobrovskyi/chat-go/internal/common/domain"
)

type UserServiceContractImpl struct {
	userService UserService
}

func (c *UserServiceContractImpl) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	return c.userService.GetUser(ctx, id)
}

func (c *UserServiceContractImpl) GetUsers(ctx context.Context, filter *domain.UserFilter) ([]domain.User, uint64, error) {
	return c.userService.GetUsers(ctx, filter)
}

func NewUserServiceContractImpl(userService UserService) *UserServiceContractImpl {
	return &UserServiceContractImpl{userService: userService}
}
