package main

import "github.com/msGo22/lesson3/worksheet/domains"

func main() {
	project := domains.NewProject("Lamazone", 5.99)
	teamLead := domains.NewTeamLead(project)
	teamLead.AddArchitecture("Adem")
	teamLead.AddDeveloper("Ali", 1)
	teamLead.AddDeveloper("Hakkı", 1)
	teamLead.AddDeveloper("Cengiz", 1)
	teamLead.AddDeveloper("Kemal", 2)
	teamLead.AddDeveloper("Cezmi", 3)
	teamLead.AddDeveloper("Nazmi", 3)
	teamLead.AddDeveloper("Kazım", 4)
	teamLead.AddDeveloper("Halil", 5)
	teamLead.AddDeveloper("Yusuf", 5)

	teamLead.Start()
	project.Statuses()
}
