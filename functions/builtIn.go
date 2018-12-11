package function

import (
	"context"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	cfg "github.com/alknopfler/alexa-profesor-guitarra/config"
	"log"
)

func Cancel(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("-- Cancel Intent --")
	response.SetStandardCard(cfg.CardTitle, cfg.SpeechCancel, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText(cfg.SpeechCancel)
	response.ShouldSessionEnd = true
	return

}

func Navigate(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("-- Navigate Intent --")
	response.SetStandardCard(cfg.CardTitle, cfg.SpeechNavigate, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText(cfg.SpeechNavigate)
	response.ShouldSessionEnd = false
	return
}

func Help(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("-- Help Intent --")
	response.SetStandardCard(cfg.CardTitle, cfg.SpeechHelp, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText(cfg.SpeechHelp)
	response.ShouldSessionEnd = false
	return
}

