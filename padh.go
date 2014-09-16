package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	"math"
)

type lookup struct {
	//note, put images in here too
	Name string
	ID int `json:"id"`
	Element int `json:"element"`
	Element2 int `json:"element2"`
	Type int `json:"type"`
	Type2 int `json:"type2"`
	CurrAwaken int `json:"current_awakening"`
	Awakenings []int `json:"awakenings"`
	LeaderSkill struct {
		Name string `json:"name"`
		HP float64 `json:"hp"`
		ATK float64 `json:"atk"`
		RCV float64 `json:"rcv"`
		Conditional [2]interface{} `json:"condition"`  //annoying -_-, no condition if this is empty
		Conditional2 [2]interface{} `json:"condition2"`
	} `json:"leader_skill"`
	
	Stats struct {
		Level int `json:"level"`
		HP int `json:"hp"`
		ATK int `json:"atk"`
		RCV int `json:"rcv"`
	} `json:"stats"`
}

func TeamLookup (ID int) (res []lookup) {
	var teamj PADHTeam
	team := PADHGet("team", ID)
	err := json.Unmarshal(team, &teamj)
	if err != nil { panic(err) }
	
	res = append(res, Lookup(teamj.Leader))
	//res = append(res, Lookup(teamj.Sub1))
	// for x := 1; x < 5; x++ {
	// 	fmt.Println(teamj[fmt.Sprintf("sub%d", x)])
	// }
	
	s := reflect.ValueOf(&teamj).Elem()
	st := s.Type()
	for x := 1; x < 5; x ++ {
		for i:= 0; i < s.NumField(); i++ {
			if st.Field(i).Name == fmt.Sprintf("Sub%d", x) {
				//res = append(re
				res = append(res, Lookup(s.Field(i).Interface().(int)))
				//fmt.Println(s.Field(i).Interface())
			}
		}		
	}

	//Friend leader time

	//friend := make (lookup)
	friend := lookup {
		Name: MonMap[teamj.FLead].Name,
		ID: teamj.FLead,
		CurrAwaken: teamj.FAwaken,

	}
	
	friend.Stats.Level = teamj.FLevel
	friend.ID = MonMap[friend.ID].ID
	friend.Awakenings = MonMap[friend.ID].Awakenings
	friend.Element = MonMap[friend.ID].Element
	friend.Element2 = MonMap[friend.ID].Element2
	friend.Type = MonMap[friend.ID].Type
	friend.Type2 = MonMap[friend.ID].Type2
	friend.LeaderSkill.Name = LdrMap[MonMap[friend.ID].LeaderSkill].Name
	if len(LdrMap[MonMap[friend.ID].LeaderSkill].Data) >= 1 {
		friend.LeaderSkill.HP = LdrMap[MonMap[friend.ID].LeaderSkill].Data[0].(float64)
		friend.LeaderSkill.ATK = LdrMap[MonMap[friend.ID].LeaderSkill].Data[1].(float64)
		friend.LeaderSkill.RCV = LdrMap[MonMap[friend.ID].LeaderSkill].Data[2].(float64)
	}
	if len(LdrMap[MonMap[friend.ID].LeaderSkill].Data) >= 4 {
		friend.LeaderSkill.Conditional = LdrMap[MonMap[friend.ID].LeaderSkill].Data[3].([2]interface{})
	}

	friend.Stats.HP = int(float64(MonMap[friend.ID].HPMin) +
		float64( MonMap[friend.ID].HPMax - MonMap[friend.ID].HPMin ) *
		( math.Pow(float64(friend.Stats.Level - 1 ) / float64(MonMap[friend.ID].MaxLevel - 1), MonMap[friend.ID].HPScale  ) )) + (teamj.FHP * plusHP)

	friend.Stats.ATK = int(float64(MonMap[friend.ID].ATKMin) +
		float64( MonMap[friend.ID].ATKMax - MonMap[friend.ID].ATKMin ) *
		( math.Pow(float64(friend.Stats.Level - 1 ) / float64(MonMap[friend.ID].MaxLevel - 1), MonMap[friend.ID].HPScale  ) )) + (teamj.FATK * plusATK)

	friend.Stats.RCV = int(float64(MonMap[friend.ID].RCVMin) +
		float64( MonMap[friend.ID].RCVMax - MonMap[friend.ID].RCVMin ) *
		( math.Pow(float64(friend.Stats.Level - 1 ) / float64(MonMap[friend.ID].MaxLevel - 1), MonMap[friend.ID].HPScale  ) )) + (teamj.FRCV * plusRCV)

	res = append(res, friend)
	return
}


func Lookup (ID int) (res lookup) {
	//fmt.Println ("Looking up monster ID:", ID)
	var monj PADHMonster
	mon := PADHGet("monster",ID)
	err := json.Unmarshal(mon, &monj)
	if err != nil { panic(err) }

	//fmt.Println(monj)

	//MonID := monj.Monster
	res.ID = monj.Monster
	//fmt.Println(u_PPJson(MonMap[res.ID], "", " "))
	res.Name = MonMap[res.ID].Name
	
	//fmt.Println(LdrMap[ MonMap[MonID].LeaderSkill])
	res.CurrAwaken = monj.CurrAwaken
	
	res.ID = MonMap[res.ID].ID
	res.Awakenings = MonMap[res.ID].Awakenings
	res.Element = MonMap[res.ID].Element
	res.Element2 = MonMap[res.ID].Element2
	res.Type = MonMap[res.ID].Type
	res.Type2 = MonMap[res.ID].Type2
	res.LeaderSkill.Name = LdrMap[MonMap[res.ID].LeaderSkill].Name
	if len(LdrMap[MonMap[res.ID].LeaderSkill].Data) >= 1 {
		res.LeaderSkill.HP = LdrMap[MonMap[res.ID].LeaderSkill].Data[0].(float64)
		res.LeaderSkill.ATK = LdrMap[MonMap[res.ID].LeaderSkill].Data[1].(float64)
		res.LeaderSkill.RCV = LdrMap[MonMap[res.ID].LeaderSkill].Data[2].(float64)
	}
	if len(LdrMap[MonMap[res.ID].LeaderSkill].Data) >= 4 {
		res.LeaderSkill.Conditional = LdrMap[MonMap[res.ID].LeaderSkill].Data[3].([2]interface{})
		if len(LdrMap[MonMap[res.ID].LeaderSkill].Data) >= 5 {
			res.LeaderSkill.Conditional2 = LdrMap[MonMap[res.ID].LeaderSkill].Data[4].([2]interface{})
		}
	}

	//1 + (98 * (e^(  ( ln(x/y)/2.5 )  )) = Z  (Z = level, X = current exp, Y = max exp)
	res.Stats.Level = int( 1 + (98 * (math.Pow(math.E, math.Log(float64(monj.CurrXP) / float64(MaxLevelExp(MonMap[res.ID].XPCurve, MonMap[res.ID].MaxLevel)))/2.5)))  )

	res.Stats.HP = int(float64(MonMap[res.ID].HPMin) +
		float64( MonMap[res.ID].HPMax - MonMap[res.ID].HPMin ) *
		( math.Pow(float64(res.Stats.Level - 1 ) / float64(MonMap[res.ID].MaxLevel - 1), MonMap[res.ID].HPScale  ) )) + (monj.PlusHP * plusHP)

	res.Stats.ATK = int(float64(MonMap[res.ID].ATKMin) +
		float64( MonMap[res.ID].ATKMax - MonMap[res.ID].ATKMin ) *
		( math.Pow(float64(res.Stats.Level - 1 ) / float64(MonMap[res.ID].MaxLevel - 1), MonMap[res.ID].HPScale  ) )) + (monj.PlusATK * plusATK)

	res.Stats.RCV = int(float64(MonMap[res.ID].RCVMin) +
		float64( MonMap[res.ID].RCVMax - MonMap[res.ID].RCVMin ) *
		( math.Pow(float64(res.Stats.Level - 1 ) / float64(MonMap[res.ID].MaxLevel - 1), MonMap[res.ID].HPScale  ) )) + (monj.PlusRCV * plusRCV)
	
	return res
}
