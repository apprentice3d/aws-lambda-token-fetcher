package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/apprentice3d/forge-api-go-client/oauth"
	"os"
)


var forgeClient oauth.TwoLeggedAuthenticator


func getToken() (oauth.Bearer, error){
	return forgeClient.Authenticate("viewables:read")
}


func main() {

	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
	forgeClient = oauth.NewTwoLeggedClient(clientID, clientSecret)

	lambda.Start(getToken)
}
