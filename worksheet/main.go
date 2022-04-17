package main

import "github.com/msGo22/lesson3/worksheet/domains"

func main() {
	level1 := domains.NewLevel(1)
	level2 := domains.NewLevel(2)
	arch := domains.NewArchitecture("Cengiz")
	developers := []*domains.Developer{
		domains.NewDeveloper("Ali", level1),
		domains.NewDeveloper("Kazım", level1),
		domains.NewDeveloper("Hakkı", level2),
	}
	teamLead := domains.NewTeamLead(arch, developers, level1, level2)
	project := domains.NewProject("Lamazon", 1.75)
	teamLead.Start(project)
	project.Statuses()
}
