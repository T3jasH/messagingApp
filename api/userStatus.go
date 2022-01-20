package api

import (
	"strconv"
)

/*
Redis model:

key			field 	value
userStatus  userId	status

status values-> online offline typing
*/

func (app *App) updateStatus(userId uint, status string){
	_, err := app.Cache.HSet("userStatus", strconv.Itoa(int(userId)), status).Result()
	if(err != nil){
		app.logger(err)
	}
}

func (app *App) getStatus(userId  string) string {
	status, err := app.Cache.HGet("userStatus", userId).Result()
	if(err != nil){
		app.logger(err)
		return ""
	}
	return status
}

func (app *App) goingOffline(userId uint){
	// Update cache and then broadcast it to all contacts who are online
	app.updateStatus(userId, "offline")
	streamData := StreamData{
		Type: "USR_STAT",
		UserStatus: UserStatus{
			UserId: userId,
			Status: "offline",
		},
	}
	app.sendStatusUpdate(streamData)
}