package dto

import "github.com/google/uuid"

type BlogResponseDto struct{
    Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	UserId uuid.UUID `json:"user_id"`
}