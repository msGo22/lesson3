package main

import "github.com/msGo22/lesson3/worksheet/domains"

func main() {
	project := domains.NewProject("Lamazon", 3)
	teamLead := domains.NewTeamLead(project)
	teamLead.AddArchitecture("Adem")
	teamLead.AddDeveloper("Ali", 1)
	teamLead.AddDeveloper("Hakkı", 1)
	teamLead.AddDeveloper("Cengiz", 1)
	teamLead.AddDeveloper("Kemal", 2)
	teamLead.AddDeveloper("Cezmi", 3)

	teamLead.Start()
	project.Statuses()
}
