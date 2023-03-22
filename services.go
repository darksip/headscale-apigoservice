package main

// returns all apikeys
func (c *Client) getApiKeys() (*ApiKeyData, error) {
	return getJsonData[ApiKeyData]("apikey", c)
}

// returns all apikeys after given expiration
func (c *Client) postApiKeysAfter(data *PostExpiration) (*ApiKey, error) {
	return postJsonData[ApiKey]("apikey", c, data)
}

// force apikey expiration
func (c *Client) postApiKeysExpire(data *PostApiKeyPrefix) (*ApiKey, error) {
	return postJsonData[ApiKey]("apikey/expire", c, data)
}

// get all machines (userId="") or machines from a given user
func (c *Client) getMachines(userId string) (*MachinesData, error) {
	if userId != "" {
		return getJsonData[MachinesData]("machine?user="+userId, c)
	}
	return getJsonData[MachinesData]("machine", c)
}

// delete machine
func (c *Client) deleteMachine(machineId string) (*ResponseError, error) {
	return deleteJsonData("machine/"+machineId, c)
}

// register machine
func (c *Client) postMachineRegister(userId string, key string) (*MachineData, error) {
	return postJsonData[MachineData, Empty]("machine?user="+userId, c, nil)
}

//get machine from Id
func (c *Client) getMachine(machineId string) (*MachineData, error) {
	return getJsonData[MachineData]("machine?machineId="+machineId, c)
}

//get routes for machine
func (c *Client) getMachineRoutes(machineId string) (*MachineData, error) {
	return getJsonData[MachineData]("machine/"+machineId+"/tags", c)
}

// force machine expiration
func (c *Client) postMachineExpire(machineId string) (*MachineData, error) {
	return postJsonData[MachineData, Empty]("machine/"+machineId+"/expire", c, nil)
}

// rename machine
func (c *Client) postMachineRename(machineId string, name string) (*MachineData, error) {
	return postJsonData[MachineData, Empty]("machine/"+machineId+"/rename/"+name, c, nil)
}

// change machine tags
func (c *Client) postMachineTags(machineId string, name string) (*MachineData, error) {
	return postJsonData[MachineData, Empty]("machine/"+machineId+"/rename/"+name, c, nil)
}

// set machine user' s
func (c *Client) postMachineUser(machineId string, userId string) (*MachineData, error) {
	return postJsonData[MachineData, Empty]("machine/"+machineId+"/user?user="+userId, c, nil)
}

// get all preAuhtKey or preauthkeys for a user
func (c *Client) getPreAuthKeys(userId string) (*PreAuthKeyData, error) {
	if userId != "" {
		return getJsonData[PreAuthKeyData]("preauthkey?user="+userId, c)
	}
	return getJsonData[PreAuthKeyData]("preauthkey", c)
}

// create new preAuthKey
func (c *Client) postPreauthkey(data *PostPreauthkey) (*PreAuthKeyData, error) {
	return postJsonData[PreAuthKeyData]("preauthkey", c, data)
}

//set machine tags for Id
func (c *Client) getRoutes() (*RoutesData, error) {
	return getJsonData[RoutesData]("routes", c)
}

// enable machine route
func (c *Client) postRouteEnable(routeId string) (*Empty, error) {
	return postJsonData[Empty, Empty]("routes/"+routeId+"/enable", c, nil)
}

// disable machine route
func (c *Client) postRouteDisable(routeId string) (*Empty, error) {
	return postJsonData[Empty, Empty]("routes/"+routeId+"/disable", c, nil)
}

// get all users
func (c *Client) getUsers() (*UsersData, error) {
	return getJsonData[UsersData]("user", c)
}

// get user by name TODO encodeurl name
func (c *Client) getUser(name string) (*UserData, error) {
	return getJsonData[UserData]("user/"+name, c)
}

// delete user
func (c *Client) deleteUser(name string) (*ResponseError, error) {
	return deleteJsonData("user/"+name, c)
}

// create new user
func (c *Client) postUser(data *PostUserName) (*UserData, error) {
	return postJsonData[UserData]("user", c, data)
}

// create new user
func (c *Client) postUserRename(oldname string, newname string) (*UserData, error) {
	return postJsonData[UserData, Empty]("user/"+oldname+"/rename/"+newname, c, nil)
}
