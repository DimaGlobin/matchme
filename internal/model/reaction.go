package model

type Reaction struct {
	ReactionType string `json:"reaction_type"`
	ReactionId   uint64 `json:"reaction_id"`
	MatchId      uint64 `json:"match_id"`
}
