package config

import "os"

const(
	AppID					= "amzn1.ask.skill.99434a76-7b33-4725-8a5f-72cb85e0e470"

	CardTitle 				= "Profesor de Guitarra"
	ImageSmall 				= "https://raw.githubusercontent.com/alknopfler/alexa-cenas/master/images/icon_108_A2Z.png"
	ImageLong 				= "https://raw.githubusercontent.com/alknopfler/alexa-cenas/master/images/icon_512_A2Z.png"

	QueryIntent				= "searchIntent"
	ComerIntent				= "comerIntent"
	CenarIntent				= "cenarIntent"
	CopasIntent				= "copasIntent"
	Cancel					= "AMAZON.CancelIntent"
	Stop					= "AMAZON.StopIntent"
	Navigate				= "AMAZON.NavigateHomeIntent"
	Help					= "AMAZON.HelpIntent"
	)

var (
	ApiKey					= os.Getenv("GEOAPIKEY")
)

