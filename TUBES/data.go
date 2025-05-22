package main

type Team struct {
	ID        int
	Name      string
	Region    string
	Points    int
	Wins      int
	ScoreDiff int
}

type Match struct {
	ID       int 
	Team1ID  int
	Team2ID  int
	Date     string
	Location string
}

const MAX_TEAMS = 50
const MAX_MATCHES = 100

var Teams [MAX_TEAMS]Team
var TeamCount int

var Matches [MAX_MATCHES]Match
var MatchCount int
