package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MehraB832/olivia_core/locales"
	"github.com/MehraB832/olivia_core/training"

	"github.com/MehraB832/olivia_core/dashboard"

	"github.com/MehraB832/olivia_core/util"

	"github.com/gookit/color"

	"github.com/MehraB832/olivia_core/network"

	"github.com/MehraB832/olivia_core/server"
)

var neuralNetworkInstances = map[string]network.Network{}

func main() {
	apiPort := flag.String("port", "8080", "The port for the API and WebSocket.")
	localesToReTrain := flag.String("re-train", "", "The locale(s) to re-train.")
	flag.Parse()

	// If the locales flag isn't empty then retrain the given models
	if *localesToReTrain != "" {
		reTrainLocales(*localesToReTrain)
	}

	// Print the Olivia ASCII text
	oliviaASCIIArt := string(util.ReadFile("res/olivia-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(oliviaASCIIArt))

	// Create the authentication token
	dashboard.Authenticate()

	for _, locale := range locales.Locales {
		util.SerializeMessages(locale.Tag)

		neuralNetworkInstances[locale.Tag] = training.CreateNeuralNetwork(
			locale.Tag,
			false,
		)
	}

	// Get port from environment variables if present
	if os.Getenv("PORT") != "" {
		*apiPort = os.Getenv("PORT")
	}

	// Start the server
	server.Serve(neuralNetworkInstances, *apiPort)
}

// reTrainLocales re-trains the given locales
func reTrainLocales(localesToReTrain string) {
	// Iterate locales by separating them by comma
	for _, localeToReTrain := range strings.Split(localesToReTrain, ",") {
		trainingFilePath := fmt.Sprintf("res/locales/%s/training.json", localeToReTrain)
		err := os.Remove(trainingFilePath)

		if err != nil {
			fmt.Printf("Cannot re-train %s model.", localeToReTrain)
			return
		}
	}
}
