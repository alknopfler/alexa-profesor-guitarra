package function

import (
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	cfg "github.com/alknopfler/alexa-profesor-guitarra/config"
	"context"
)

func AcordeIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	response.SetStandardCard(cfg.CardTitle, cfg.SpeechCancel, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText("El acorde elegido es: "+request.Intent.Slots["acorde"].Value+request.Intent.Slots["tipo"].Value)
	response.ShouldSessionEnd = true
	return

}