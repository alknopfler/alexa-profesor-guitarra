package function

import (
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	cfg "github.com/alknopfler/alexa-profesor-guitarra/config"
	"context"
	"log"
)

func AcordeIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("-- Acorde Intent --")
	if request.DialogState != "COMPLETED" {
		log.Println("Get into dialog to confirm name 'acordeIntent'")
		response.AddDialogDirective("Dialog.Delegate", "", "", &request.Intent)
		response.ShouldSessionEnd = false

	} else {

		response.SetStandardCard(cfg.CardTitle, cfg.SpeechCancel, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText("El acorde elegido es: " + request.Intent.Slots["acorde"].Value)
		response.ShouldSessionEnd = true
		return
	}
}
