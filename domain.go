package main

import "time"

type ApiKey struct {
	ID         string    `json:"id"`
	Prefix     string    `json:"prefix"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	LastSeen   time.Time `json:"lastSeen"`
}

type Machine struct {
	ID                   string     `json:"id"`
	MachineKey           string     `json:"machineKey"`
	NodeKey              string     `json:"nodeKey"`
	DiscoKey             string     `json:"discoKey"`
	IPAddresses          []string   `json:"ipAddresses"`
	Name                 string     `json:"name"`
	User                 User       `json:"user"`
	LastSeen             time.Time  `json:"lastSeen"`
	LastSuccessfulUpdate time.Time  `json:"lastSuccessfulUpdate"`
	Expiry               time.Time  `json:"expiry"`
	PreAuthKey           PreAuthKey `json:"preAuthKey"`
	CreatedAt            time.Time  `json:"createdAt"`
	RegisterMethod       string     `json:"registerMethod"`
	ForcedTags           []string   `json:"forcedTags"`
	InvalidTags          []string   `json:"invalidTags"`
	ValidTags            []string   `json:"validTags"`
	GivenName            string     `json:"givenName"`
	Online               bool       `json:"online"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type PreAuthKey struct {
	User       string    `json:"user"`
	ID         string    `json:"id"`
	Key        string    `json:"key"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Used       bool      `json:"used"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	ACLTags    []string  `json:"aclTags"`
}

type Route struct {
	ID         string    `json:"id"`
	Machine    Machine   `json:"machine"`
	Prefix     string    `json:"prefix"`
	Advertised bool      `json:"advertised"`
	Enabled    bool      `json:"enabled"`
	IsPrimary  bool      `json:"isPrimary"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

type ApiKeyData struct {
	ApiKeys []ApiKey `json:"apiKeys"`
}

type RoutesData struct {
	Routes []Route `json:"routes"`
}

type MachinesData struct {
	Machines []Machine `json:"machines"`
}

type MachineData struct {
	Machine Machine `json:"machine"`
}

type UsersData struct {
	Users []User `json:"users"`
}

type UserData struct {
	User User `json:"user"`
}

type PreAuthKeyData struct {
	PreAuthKey PreAuthKey `json:"preAuthKey"`
}

type Data interface {
	RoutesData | MachinesData | UsersData | UserData | MachineData | PreAuthKeyData | ApiKeyData | ApiKey | Empty
}

type PostExpiration struct {
	Expiration time.Time `json:"expiration"`
}

type PostApiKeyPrefix struct {
	Prefix string `json:"prefix"`
}

type PostMachineTags struct {
	Tags []string `json:"tags"`
}

type PostUserName struct {
	Name string `json:"name"`
}

type PostPreauthkey struct {
	User       string    `json:"user"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Expiration time.Time `json:"expiration"`
	ACLTags    []string  `json:"aclTags"`
}

type Empty interface{}

type RequestData interface {
	Empty | PostApiKeyPrefix | PostExpiration | PostMachineTags | PostPreauthkey | PostUserName
}

type ResponseError struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}

type ErrorDetail interface{}
