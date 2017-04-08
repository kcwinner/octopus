package main

import (
	"log"
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/kcwinner/octopus/client"
)

var (
	app        = kingpin.New("octopus", "A command-line octopus application.")
	debug      = app.Flag("debug", "Enable debug mode.").Bool()
	server     = app.Flag("server", "URL to the server.").Required().String()
	apikey     = app.Flag("apikey", "Apikey to the server.").Required().String()
	formvalues = app.Flag("formvalues", "Form values for the promotion.").PlaceHolder("KEY:VALUE").StringMap()

	promoteRelease  = app.Command("promote-release", "Promote a release.")
	project         = promoteRelease.Arg("project", "Project to promote.").Required().String()
	fromEnvironment = promoteRelease.Arg("from", "Environment to promote from.").Required().String()
	toEnvironment   = promoteRelease.Arg("to", "Environment to promote to.").Required().String()

	getProject  = app.Command("get-project", "Gets a project.")
	projectName = getProject.Arg("project", "Project to promote.").Required().String()
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	val, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	if *debug {
		log.Println("Debug enabled")
		log.Println("Server:", *server)
		log.Println("APIkey:", *apikey)
	}

	octopus, err := client.NewOctopusClient(*debug, *server, *apikey)
	if err != nil {
		log.Fatal("Error creating Octopus client:", err)
	}

	switch kingpin.MustParse(val, err) {

	case promoteRelease.FullCommand():
		if octopus.Debug {
			log.Printf("Promoting release:\n\tProject: %s\n\tFrom: %s\n\tTo: %s", *project, *fromEnvironment, *toEnvironment)
		}

		err = octopus.PromoteRelease(*project, *fromEnvironment, *toEnvironment, *formvalues)
		if err != nil {
			log.Fatal(err)
		}

	case getProject.FullCommand():
		if octopus.Debug {
			log.Println("Getting Project:", *projectName)
		}
		result, err := octopus.GetProjectByName(*projectName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}
}
