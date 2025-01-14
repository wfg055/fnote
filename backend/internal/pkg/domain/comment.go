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

package domain

type LatestComment struct {
	PostInfo4Comment
	Name       string
	Content    string
	CreateTime int64
}

type CommentWithReplies struct {
	Comment
	Replies []CommentReply
}

type Comment struct {
	Id string
	// 文章信息
	PostInfo PostInfo4Comment
	// 评论的内容
	Content string
	// 用户信息
	UserInfo   UserInfo4Comment
	CreateTime int64
}

type CommentReply struct {
	ReplyId string
	// 回复内容
	Content string
	// 被回复的回复 Id
	ReplyToId string
	// 用户信息
	UserInfo UserInfo4Reply
	// 被回复用户的信息
	RepliedUserInfo UserInfo4Reply
	Status          CommentStatus
	CreateTime      int64
}

type UserInfo4Reply UserInfo4Comment

type PostInfo4Comment struct {
	// 文章 ID
	PostId string
	// 文章标题字段
	PostTitle string
}

type UserInfo4Comment struct {
	Name    string
	Email   string
	Ip      string
	Website string
}

type CommentStatus uint

const (
	// CommentStatusPending 审核中
	CommentStatusPending CommentStatus = iota
	// CommentStatusApproved 审核通过
	CommentStatusApproved
	// CommentStatusRejected 审核不通过
	CommentStatusRejected
)
