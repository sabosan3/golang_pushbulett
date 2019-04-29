package lib

import (
	gopb "github.com/xconstruct/go-pushbullet"
)

/*
createMessage return string
*/
func createMessage(contents ...string) string {
	message := ""
	for _, c := range contents {
		message += c
		message += "\n"
	}
	return message
}

/*
PostMessage return
*/
func PostMessage(token string, title string, contents ...string) {

	logger := InitLogging()

	pb := gopb.New(token)
	devices, err := pb.Devices()
	if err != nil {
		logger.Printf("error: failed to get devices")
		logger.Printf(err.Error())
	}

	message := createMessage(contents...)

	err = pb.PushNote(devices[0].Iden, title, message)
}
