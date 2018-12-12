package function

import (
	"errors"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	cfg "github.com/alknopfler/alexa-profesor-guitarra/config"
	"context"
	"log"
	"strings"
)

type ChordDefinition struct {
	image string
	sound string
	text  string
}

var Chord = map[string]ChordDefinition{}

func AcordeIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("-- Acorde Intent --")

	if request.DialogState != "COMPLETED" {
		log.Println("Get into dialog to confirm name 'acordeIntent'")
		response.AddDialogDirective("Dialog.Delegate", "", "", &request.Intent)
		response.ShouldSessionEnd = false

	} else {

		acorde, err:= getAcorde(request)

		if err != nil{
			response.SetSimpleCard(cfg.CardTitle, cfg.SpeechAcordeNotFound)
			response.SetOutputText(cfg.SpeechAcordeNotFound)
			response.ShouldSessionEnd = true
			return
		}
		response.SetStandardCard(cfg.CardTitle, acorde.text, acorde.image, acorde.image)
		response.SetOutputText(cfg.SpeechAcorde)
		response.AddAudioPlayer("AudioPlayer.Play","REPLACE_ALL", acorde.sound, acorde.sound,0)
		response.SetRepromptText("Qué quieres hacer ahora?")
		return
	}
}

func getAcorde(request *alexa.Request) (ChordDefinition,error){

	if request.Intent.Slots["acorde"].Value != ""{
		acorde := strings.ToLower(request.Intent.Slots["acorde"].Value)
		if val, ok := Chord[acorde]; ok {
			return val,nil
		}
		return ChordDefinition{}, errors.New("Not found key in map")
	}
	return ChordDefinition{}, errors.New("Not acorde recorded")
}

func InitChordTable(){
	Chord["do mayor"] = ChordDefinition{cfg.ImageCmaj, cfg.SoundCmaj, cfg.SpeechCmaj}
	Chord["do menor"] = ChordDefinition{cfg.ImageCmin, cfg.SoundCmin, cfg.SpeechCmin}
	Chord["re mayor"] = ChordDefinition{cfg.ImageDmaj, cfg.SoundDmaj, cfg.SpeechDmaj}
	Chord["re menor"] = ChordDefinition{cfg.ImageDmin, cfg.SoundDmin, cfg.SpeechDmin}
	Chord["mi mayor"] = ChordDefinition{cfg.ImageEmaj, cfg.SoundEmaj, cfg.SpeechEmaj}
	Chord["mi menor"] = ChordDefinition{cfg.ImageEmin, cfg.SoundEmin, cfg.SpeechEmin}
	Chord["fa mayor"] = ChordDefinition{cfg.ImageFmaj, cfg.SoundFmaj, cfg.SpeechFmaj}
	Chord["fa menor"] = ChordDefinition{cfg.ImageFmin, cfg.SoundFmin, cfg.SpeechFmin}
	Chord["sol mayor"] = ChordDefinition{cfg.ImageGmaj, cfg.SoundGmaj, cfg.SpeechGmaj}
	Chord["sol menor"] = ChordDefinition{cfg.ImageGmin, cfg.SoundGmin, cfg.SpeechGmin}
	Chord["la mayor"] = ChordDefinition{cfg.ImageAmaj, cfg.SoundAmaj, cfg.SpeechAmaj}
	Chord["la menor"] = ChordDefinition{cfg.ImageAmin, cfg.SoundAmin, cfg.SpeechAmin}
	Chord["si mayor"] = ChordDefinition{cfg.ImageBmaj, cfg.SoundBmaj, cfg.SpeechBmaj}
	Chord["si menor"] = ChordDefinition{cfg.ImageBmin, cfg.SoundBmin, cfg.SpeechBmin}
}