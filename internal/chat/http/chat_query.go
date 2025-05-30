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

type ChatQuery struct {
	IDs          []uint64 `query:"id" validate:"omitempty,dive,gte=0"`
	Types        []uint8  `query:"types" validate:"omitempty,dive,oneof=1 2"`
	CreatedByIDs []uint64 `query:"createdByIds" validate:"omitempty,dive,gte=0"`

	Search string `query:"search"`

	Limit  *uint64 `query:"limit"`
	Offset *uint64 `query:"offset"`

	Sort string `query:"sort"`
}
