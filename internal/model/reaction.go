package model

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
