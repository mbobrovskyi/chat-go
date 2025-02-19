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

package domain

import (
	"time"

	"chat-go/internal/common/domain"
)

type MessageStatus uint8

func (ms MessageStatus) ToUint8() uint8 {
	return uint8(ms)
}

const (
	DraftMessageStatus  MessageStatus = 1
	UnreadMessageStatus MessageStatus = 2
	ReadMessageStatus   MessageStatus = 3
)

type Message struct {
	ID        uint64
	Text      string
	Status    MessageStatus
	ChatID    uint64
	CreatedBy uint64
	Creator   *domain.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
