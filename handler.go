package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	cfg "github.com/alknopfler/alexa-profesor-guitarra/config"
	f "github.com/alknopfler/alexa-profesor-guitarra/functions"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
)


var a = &alexa.Alexa{ApplicationID: cfg.AppID, RequestHandler: &ProfesorGuitarra{}, IgnoreTimestamp: true}

// ProfesorGuitarra struct for request from the ProfesorGuitarra skill.
type ProfesorGuitarra struct{}

// Handle processes calls from Lambda
func Handle(ctx context.Context, requestEnv *alexa.RequestEnvelope) (interface{}, error) {
	return a.ProcessRequest(ctx, requestEnv)
}

// OnSessionStarted called when a new session is created.
func (h *ProfesorGuitarra) OnSessionStarted(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {
	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
		//Can be usefull to login internally with the end service
	f.InitChordsTable()
	return nil
}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *ProfesorGuitarra) OnLaunch(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {
	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnLaunch, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText(cfg.SpeechOnLaunch)

	response.ShouldSessionEnd = false
	response.SetRepromptText(cfg.SpeechNavigate)

	return nil
}

// OnIntent called with a reqeust is received of type IntentRequest
func (h *ProfesorGuitarra) OnIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {
	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)

	switch request.Intent.Name {

	case cfg.AcordeIntent:
		f.AcordeIntent(context, request, session, aContext, response)

	case cfg.Cancel,cfg.Stop:
		f.Cancel(context, request, session, aContext, response)
		response.ShouldSessionEnd = true
		return nil

	case cfg.Navigate:
		f.Navigate(context, request, session, aContext, response)

	case cfg.Help:
		f.Help(context, request, session, aContext, response)

	default:
		log.Println("Entra por default Intent")
		f.Navigate(context, request, session, aContext, response)
	}
	log.Println("antes del nil onIntent")
	return nil
}

// OnSessionEnded called with a reqeust is received of type SessionEndedRequest
func (h *ProfesorGuitarra) OnSessionEnded(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {
	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
	response.ShouldSessionEnd = true
	return nil
}


func main() {
	lambda.Start(Handle)
}