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

package api

import (
	"net/http"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	chatdomain "chat-go/internal/chat/domain"
	chathttp "chat-go/internal/chat/http"
	"chat-go/test/helpers"
	"chat-go/test/integration/framework"
)

var _ = ginkgo.Describe("Chat", ginkgo.Ordered, ginkgo.ContinueOnFailure, func() {
	var (
		httpClient helpers.HTTPClient
		chat       *chathttp.ChatDto
	)

	ginkgo.BeforeAll(func() {
		httpClient = framework.NewTestHTTPClient(fwk).WithTimeout(helpers.Timeout)

		chat = helpers.CreateChat(httpClient, "", helpers.AdminToken, &chathttp.CreateChatDto{
			Name: "Chat",
			Type: uint8(chatdomain.GroupChatType),
		})
	})

	ginkgo.AfterAll(func() {
		helpers.DeleteChat(httpClient, "", helpers.AdminToken, chat.ID)
	})

	ginkgo.Context("get chats endpoint", func() {
		ginkgo.It("should return valid response", func() {
			chats := helpers.GetChats(httpClient, "", helpers.AdminToken)
			gomega.Expect(chats.Items).To(gomega.HaveLen(1))
			gomega.Expect(chats.Count).To(gomega.Equal(uint64(1)))
			gomega.Expect(&chats.Items[0]).To(gomega.BeComparableTo(chat))
		})
	})

	ginkgo.Context("delete chat endpoint", func() {
		ginkgo.It("shouldn't delete not owned chat", func() {
			ginkgo.By("creating chat", func() {
				chat = helpers.CreateChat(httpClient, "", helpers.AdminToken, &chathttp.CreateChatDto{
					Name: "Chat",
					Type: uint8(chatdomain.GroupChatType),
				})
			})

			ginkgo.By("deleting chat", func() {
				helpers.DeleteChatWithStatus(httpClient, "", helpers.UserToken, chat.ID, http.StatusForbidden)
			})
		})
	})
	ginkgo.Context("update chat endpoint", func() {
		ginkgo.It("should return forbidden error when unauthorized user tries to update chat", func() {
			updateChatRequest := &chathttp.UpdateChatDto{
				Name: "Updated Group Chat",
			}

			updatedChat := helpers.UpdateChat(httpClient, "", helpers.UserToken, chat.ID, updateChatRequest, http.StatusForbidden)
			gomega.Expect(updatedChat).To(gomega.BeComparableTo(
				&chathttp.ChatDto{},
				cmpopts.IgnoreFields(chathttp.ChatDto{}, "CreatedAt", "UpdatedAt"),
			))
		})
	})
})
