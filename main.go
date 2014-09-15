package main

import (
	"encoding/json"
	"fmt"
	"math"
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

func main () {
	// r := mux.NewRouter()
	// r.HandleFunc("/ws", remoteHandler)
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))
	// http.ListenAndServe(":8080", r)


	//fmt.Println(MonMap[752].Name)

	fmt.Println(Lookup(1781703))
	//fmt.Println(LdrMap[MonMap[752].LeaderSkill])
}

type lookup struct {
	ID int `json:"id"`
	CurrAwaken int `json:"current_awakening"`
	Awakenings []int `json:"awakenings"`
	LeaderSkill struct {
		HP float64 `json:"hp"`
		ATK float64 `json:"atk"`
		RCV float64 `json:"rcv"`
		Conditional [2]interface{} `json:"condition"`  //annoying -_-, no condition if this is empty
	} `json:"leader_skill"`
	
	Stats struct {
		Level int `json:"level"`
		HP int `json:"hp"`
		ATK int `json:"atk"`
		RCV int `json:"rcv"`
	} `json:"stats"`
}
//Will return: { id, awakenings, leaderskill data, stats stuff }
func MaxLevelExp(Curve int, MLvl int) (res int) {
	res = int(float64(Curve) * (  math.Pow( ((float64(MLvl) - 1) * 50/49), 2.5 )  ))

	return res
}
func Lookup (ID int) (res lookup) {
	var monj PADHMonster
	mon := getMon(ID)
	err := json.Unmarshal(mon, &monj)
	if err != nil { panic(err) }

	MonID := monj.Monster
	res.CurrAwaken = monj.CurrAwaken
	
	res.ID = MonMap[MonID].ID
	res.Awakenings = MonMap[MonID].Awakenings
	res.LeaderSkill.HP = LdrMap[MonMap[MonID].LeaderSkill].Data[0].(float64)
	res.LeaderSkill.ATK = LdrMap[MonMap[MonID].LeaderSkill].Data[1].(float64)
	res.LeaderSkill.RCV = LdrMap[MonMap[MonID].LeaderSkill].Data[2].(float64)
	if len(LdrMap[MonMap[MonID].LeaderSkill].Data) >= 4 {
		res.LeaderSkill.Conditional = LdrMap[MonMap[MonID].LeaderSkill].Data[3].([2]interface{})
	}


	//Stats
	//fmt.Println(MaxLevelExp(MonMap[MonID].XPCurve, MonMap[MonID].MaxLevel))
	
	res.Stats.Level = int(1 + ( 98 * ( math.Pow(math.E, (math.Log( float64(monj.CurrXP/ MaxLevelExp(MonMap[MonID].XPCurve, MonMap[MonID].MaxLevel)) ) / 2.5) ) ) ))

	res.Stats.HP = int(float64(MonMap[MonID].HPMin) +
		float64( MonMap[MonID].HPMax - MonMap[MonID].HPMin ) *
		( math.Pow(float64(res.Stats.Level - 1 ) / float64(MonMap[MonID].MaxLevel - 1), MonMap[MonID].HPScale  ) ))
	// res.Stats.HP = float64(MonMap[MonID].HPMin) +
	// 	math.Pow( ( float64(MonMap[MonID].HPMax) - float64(MonMap[MonID].HPMin) ), MonMap[MonID].HPScale) *
	// 	math.Pow( ( float64(res.Stats.Level - 1) / float64(MonMap[MonID].MaxLevel -1
	
	return res
}












