package models

type Tag struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type TagListResponse struct {
	Data []Tag `json:"data"`
}

type TicketTagsRequest struct {
	Tags []string `json:"tags"`
}

type TicketMergeRequest struct {
	TargetTicketID string `json:"targetTicketId"`
}

type FollowRequest struct {
	Follow bool `json:"follow"`
}