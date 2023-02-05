package common

import "Douyin/service/user"

// 该文件为demo自带，修改时请注意调用关系

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            uint      `json:"id,omitempty"`
	Author        user.Info `json:"author"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount uint      `json:"favorite_count,omitempty"`
	CommentCount  uint      `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         uint      `json:"id,omitempty"`
	User       user.Info `json:"user"`
	Content    string    `json:"content,omitempty"`
	CreateDate string    `json:"create_date,omitempty"`
}

type Message struct {
	Id         uint   `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     uint   `json:"user_id,omitempty"`
	ToUserId   uint   `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId uint   `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
