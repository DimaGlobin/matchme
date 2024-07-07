package model

import "time"

type Like struct {
	Id           uint64    `gorm:"primaryKey;column:like_id"`
	LikingId     uint64    `gorm:"column:liking_id"`
	LikedId      uint64    `gorm:"column:liked_id"`
	CreationDate time.Time `gorm:"column:creation_date"`
}

func (Like) TableName() string {
	return "likes"
}

type Dislike struct {
	Id           uint64    `gorm:"primaryKey;column:dislike_id"`
	DislikingId  uint64    `gorm:"column:disliking_id"`
	DislikedId   uint64    `gorm:"column:disliked_id"`
	CreationDate time.Time `gorm:"column:creation_date"`
}

func (Dislike) TableName() string {
	return "dislikes"
}

type Match struct {
	MatchId      uint64    `gorm:"primaryKey;column:match_id"`
	User1Id      uint64    `gorm:"column:user1"`
	User2Id      uint64    `gorm:"column:user2"`
	CreationDate time.Time `gorm:"column:creation_date"`
}

func (Match) TableName() string {
	return "matches"
}

type DislikeResp struct {
	ReactionType string `json:"reaction_type"`
	ReactionId   uint64 `json:"reaction_id"`
}

type LikeResp struct {
	ReactionType string `json:"reaction_type"`
	ReactionId   uint64 `json:"reaction_id"`
	MatchId      uint64 `json:"match_id"`
	LikesLeft    *int   `json:"likes_left,omitempty"`
}
