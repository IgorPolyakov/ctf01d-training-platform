// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package models

import (
	"time"
)

// Defines values for UserRequestRole.
const (
	UserRequestRoleAdmin  UserRequestRole = "admin"
	UserRequestRoleGuest  UserRequestRole = "guest"
	UserRequestRolePlayer UserRequestRole = "player"
)

// Defines values for UserResponseRole.
const (
	UserResponseRoleAdmin  UserResponseRole = "admin"
	UserResponseRoleGuest  UserResponseRole = "guest"
	UserResponseRolePlayer UserResponseRole = "player"
)

// GameRequest defines model for GameRequest.
type GameRequest struct {
	// Description A brief description of the game
	Description *string `json:"description,omitempty"`

	// EndTime The end time of the game
	EndTime time.Time `json:"end_time"`

	// StartTime The start time of the game
	StartTime time.Time `json:"start_time"`
}

// GameResponse defines model for GameResponse.
type GameResponse struct {
	// Description A brief description of the game
	Description *string `json:"description,omitempty"`

	// EndTime The end time of the game
	EndTime time.Time `json:"end_time"`

	// Id Unique identifier for the game
	Id string `json:"id"`

	// StartTime The start time of the game
	StartTime time.Time `json:"start_time"`
}

// ResultRequest defines model for ResultRequest.
type ResultRequest struct {
	// GameId Identifier of the game this result is for
	GameId string `json:"game_id"`

	// Rank The rank achieved by the team in this game
	Rank int `json:"rank"`

	// Score The score achieved by the team
	Score int `json:"score"`

	// TeamId Identifier of the team this result belongs to
	TeamId string `json:"team_id"`
}

// ResultResponse defines model for ResultResponse.
type ResultResponse struct {
	// GameId Identifier of the game this result is for
	GameId string `json:"game_id"`

	// Id Unique identifier for the result entry
	Id string `json:"id"`

	// Rank The rank achieved by the team in this game
	Rank int `json:"rank"`

	// Score The score achieved by the team
	Score int `json:"score"`

	// TeamId Identifier of the team this result belongs to
	TeamId string `json:"team_id"`
}

// ServiceRequest defines model for ServiceRequest.
type ServiceRequest struct {
	// Author Author of the service
	Author string `json:"author"`

	// Description A brief description of the service
	Description *string `json:"description,omitempty"`

	// IsPublic Boolean indicating if the service is public
	IsPublic bool `json:"is_public"`

	// LogoUrl URL to the logo of the service
	LogoUrl *string `json:"logo_url,omitempty"`

	// Name Name of the service
	Name string `json:"name"`
}

// ServiceResponse defines model for ServiceResponse.
type ServiceResponse struct {
	// Author Author of the service
	Author string `json:"author"`

	// Description A brief description of the service
	Description *string `json:"description,omitempty"`

	// Id Unique identifier for the service
	Id string `json:"id"`

	// IsPublic Boolean indicating if the service is public
	IsPublic bool `json:"is_public"`

	// LogoUrl URL to the logo of the service
	LogoUrl *string `json:"logo_url,omitempty"`

	// Name Name of the service
	Name string `json:"name"`
}

// TeamRequest defines model for TeamRequest.
type TeamRequest struct {
	// AvatarUrl URL to the team's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description A brief description of the team
	Description *string `json:"description,omitempty"`

	// Name Name of the team
	Name string `json:"name"`

	// SocialLinks JSON string containing social media links of the team
	SocialLinks *string `json:"social_links,omitempty"`

	// University University or institution the team is associated with
	University *string `json:"university,omitempty"`
}

// TeamResponse defines model for TeamResponse.
type TeamResponse struct {
	// AvatarUrl URL to the team's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description A brief description of the team
	Description *string `json:"description,omitempty"`

	// Id Unique identifier for the team
	Id string `json:"id"`

	// Name Name of the team
	Name string `json:"name"`

	// SocialLinks JSON string containing social media links of the team
	SocialLinks *string `json:"social_links,omitempty"`

	// University University or institution the team is associated with
	University *string `json:"university,omitempty"`
}

// UserRequest defines model for UserRequest.
type UserRequest struct {
	// AvatarUrl URL to the user's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Password User password
	Password *string `json:"password,omitempty"`

	// Role The role of the user (admin, player or guest)
	Role *UserRequestRole `json:"role,omitempty"`

	// Status Status of the user (active, disabled)
	Status *string `json:"status,omitempty"`

	// UserName The name of the user
	UserName *string `json:"user_name,omitempty"`
}

// UserRequestRole The role of the user (admin, player or guest)
type UserRequestRole string

// UserResponse defines model for UserResponse.
type UserResponse struct {
	// AvatarUrl URL to the user's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Id The unique identifier for the user
	Id *string `json:"id,omitempty"`

	// Role The role of the user (admin, player or guest)
	Role *UserResponseRole `json:"role,omitempty"`

	// Status Status of the user (active, disabled)
	Status *string `json:"status,omitempty"`

	// UserName The name of the user
	UserName *string `json:"user_name,omitempty"`
}

// UserResponseRole The role of the user (admin, player or guest)
type UserResponseRole string

// GetApiUniversitiesParams defines parameters for GetApiUniversities.
type GetApiUniversitiesParams struct {
	// Term Optional search term to filter universities by name.
	Term *string `form:"term,omitempty" json:"term,omitempty"`
}

// PostApiUsersLoginJSONBody defines parameters for PostApiUsersLogin.
type PostApiUsersLoginJSONBody struct {
	Password *string `json:"password,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// CreateGameJSONRequestBody defines body for CreateGame for application/json ContentType.
type CreateGameJSONRequestBody = GameRequest

// UpdateGameJSONRequestBody defines body for UpdateGame for application/json ContentType.
type UpdateGameJSONRequestBody = GameRequest

// CreateResultJSONRequestBody defines body for CreateResult for application/json ContentType.
type CreateResultJSONRequestBody = ResultRequest

// CreateServiceJSONRequestBody defines body for CreateService for application/json ContentType.
type CreateServiceJSONRequestBody = ServiceRequest

// UpdateServiceJSONRequestBody defines body for UpdateService for application/json ContentType.
type UpdateServiceJSONRequestBody = ServiceRequest

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = TeamRequest

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.
type UpdateTeamJSONRequestBody = TeamRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = UserRequest

// PostApiUsersLoginJSONRequestBody defines body for PostApiUsersLogin for application/json ContentType.
type PostApiUsersLoginJSONRequestBody PostApiUsersLoginJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UserRequest