package main

import (
	"encoding/json"
//	"fmt"
	"strings"
	"strconv"
//	"math"
	
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"net/http"

	"log"
)

const (
	plusHP = 10
	plusATK = 5
	plusRCV = 3
)
var MonMap map[int]*Monster
var AwkMap map[int]*Awakenings
var LdrMap map[string]*LeaderSkill

func init () {
	MonMap = make(map[int]*Monster)
	AwkMap = make(map[int]*Awakenings)
	LdrMap = make(map[string]*LeaderSkill)
	var err error

	//Monsters
	var monjson []Monster
	monsters := read("monsters")
	err = json.Unmarshal(monsters, &monjson)
	if err != nil {panic(err)}

	for x, y := range monjson {
		MonMap[y.ID] = &monjson[x]
	}

	//Awakenings
	var awkjson []Awakenings
	awakenings := read("awakenings")
	err = json.Unmarshal(awakenings, &awkjson)
	if err != nil {panic(err)}

	for x,y := range awkjson {
		AwkMap[y.ID] = &awkjson[x]
	}

	//Leader Skills
	var ldrjson []LeaderSkill
	leaderskills := read("leader_skills")
	err = json.Unmarshal(leaderskills, &ldrjson)
	if err != nil {panic(err)}

	for x,y := range ldrjson {
		LdrMap[y.Name] = &ldrjson[x]
	}
	//Going to update files 1/24h
	//Take list of monsters, make map[monster_id int]Monster

	
	//MonsterList 
}

func teamHandler(res http.ResponseWriter, req *http.Request) {
	var err error;
	var msg *orbs
	var dmg []float64
	var teamD []teamDamage
	//var comboCount, comboMulti float64 
	log.Printf(req.URL.String())
	teamID, err := strconv.Atoi(strings.Split(req.URL.String(),"/")[2])
	if err != nil { panic(err) }

	//log.Println(teamID)
	
	ws, _ := websocket.Upgrade(res, req, nil, 1024, 1024)
	log.Printf("got websocket conn from %v\n", ws.RemoteAddr())

	team := TeamLookup(teamID)

	err = ws.WriteJSON(team)
	if err != nil { panic(err) }
	for {
		msg = new (orbs)
		//dmg = new ([]float64)
		dmg = make([]float64, 6, 6)
		//dmg = new ([6]float64)
		teamD = make ([]teamDamage, 6, 6)
		// comboCount = 0
		// comboMulti = 0
		
		if err = ws.ReadJSON(msg); err != nil {
			panic(err)
		}
		
		//Do da maff, return da damage, given team (which is team info)
		teamD =  damageResolve(team,teamD,dmg,msg)


		//teamD = damageResolve (team, teamD, dmg, comboMulti)

		ws.WriteJSON(teamD)
	}
	
}


func main () {
	//I wanna test some shit
	team := TeamLookup(77475)
	//var dmg [6]float64
	msg := new (orbs)
	
	dmg := make([]float64, 6, 6)
	teamD := make ([]teamDamage, 6, 6)

	msg.Light = []float64{ 3, 3 }
	msg.Fire = []float64{3}
	msg.Water = []float64{3}
	msg.Heart = []float64{3}

	teamD = damageResolve(team, teamD, dmg, msg)

	log.Println(u_PPJson(teamD, "", " "))
	log.Printf("Starting server")
	r := mux.NewRouter()
	r.HandleFunc("/team/{id:[0-9]+/}", teamHandler)
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))
	http.ListenAndServe(":8080", r)


	//fmt.Println(MonMap[752].Name)


	//fmt.Println(Lookup(1781703))

	//fmt.Println(u_PPJson(TeamLookup(77475),""," "))
	//fmt.Println(LdrMap[MonMap[752].LeaderSkill])
}

