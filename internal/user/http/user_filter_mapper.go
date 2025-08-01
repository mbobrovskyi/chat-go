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

package http

import (
	"chat-go/internal/common/domain"
	"chat-go/internal/common/errors"
	"chat-go/internal/common/http"
	"chat-go/internal/user/constants"
)

var userSortFields = []string{
	"id",
	"email",
	"username",
	"role",
	"firstName",
	"lastName",
}

func UserFilterFromQuery(query UserQuery) (domain.UserFilter, error) {
	sort, err := http.SortFromDto(query.Sort, userSortFields)
	if err != nil {
		return domain.UserFilter{}, errors.NewBadRequestError(
			constants.UserDomain, err, map[string]any{"sort": query.Sort})
	}

	return domain.UserFilter{
		IDs:       query.IDs,
		Emails:    query.Emails,
		Usernames: query.Usernames,
		Search:    query.Search,
		Limit:     query.Limit,
		Offset:    query.Offset,
		Sort:      sort,
	}, nil
}
