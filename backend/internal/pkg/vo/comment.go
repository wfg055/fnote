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

package vo

type LatestCommentVO struct {
	PostInfo4Comment
	Name       string `json:"name"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostInfo4Comment struct {
	// 文章 ID
	PostId string `json:"post_id"`
	// 文章标题字段
	PostTitle string `json:"post_title"`
}

type PostCommentVO struct {
	Id string `json:"id"`
	// 评论的内容
	Content string `json:"content"`
	// 评论的用户
	Name        string               `json:"username"`
	CommentTime int64                `json:"comment_time"`
	Replies     []PostCommentReplyVO `json:"replies,omitempty"`
}

type PostCommentReplyVO struct {
	Id        string `json:"id"`
	CommentId string `json:"comment_id"`
	// 回复的内容
	Content string `json:"content"`
	// 回复的用户
	Name string `json:"name"`
	// 被回复的回复 Id
	ReplyToId string `json:"reply_to_id"`
	// 被回复的用户
	ReplyTo string `json:"reply_to"`
	// 回复时间
	ReplyTime int64 `json:"reply_time"`
}
