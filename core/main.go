// Package provides Discord bot's core functionality.
package core

import "github.com/bwmarrin/discordgo"

var (
	Session *discordgo.Session
)

func Start(session *discordgo.Session) {
	Session = session
	// set status
	session.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online", // online, idle, dnd or invisible
		Activities: []*discordgo.Activity{{
			Name:  "Visual Studio Code",
			State: "We be codin'",
			Type:  discordgo.ActivityTypeGame,
		}},
	})
}

func Stop() {
}
