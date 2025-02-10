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
	"context"

	"mbobrovskyi/chat-go/internal/common/repository"
)

type MessageRepo interface {
	GetMessages(ctx context.Context, filter *MessageFilter) ([]Message, error)
	GetMessagesCount(ctx context.Context, filter *MessageFilter) (uint64, error)
	CreateMessage(ctx context.Context, message Message, tx repository.Tx) (*Message, error)
	UpdateMessageStatus(
		ctx context.Context,
		messageIDs []uint64,
		messageStatus MessageStatus,
		tx repository.Tx,
	) error
}
