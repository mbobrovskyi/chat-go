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
	"mbobrovskyi/chat-go/internal/common/http"
	"mbobrovskyi/chat-go/internal/user/domain"
)

func TokenToDto(token domain.Token) TokenDto {
	return TokenDto{
		Token:   token.Token,
		ExpIn:   token.ExpIn.Milliseconds() / 1000,
		ExpTime: token.ExpTime,
		User:    http.UserToDto(token.User),
	}
}
