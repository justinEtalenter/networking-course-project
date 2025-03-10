package models

type Board struct {
	Id                 int         `json:"id,omitempty"`
	ProjectId          int         `json:"projectId"`
	OwnerId            int         `json:"ownerId,omitempty"`
	DefaultPermissions *Permission `json:"defaultPermissions,omitempty"`
	Datetimes          *Datetimes  `json:"datetimes,omitempty"`
	Title              string      `json:"title"`
}

type UpdateBoard struct {
	DefaultPermissions *UpdatePermission `json:"defaultPermissions,omitempty"`
	Datetimes          *UpdateDatetimes  `json:"datetimes,omitempty"`
	Title              *string           `json:"title"`
}
