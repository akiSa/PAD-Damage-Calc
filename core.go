package main

import (
	"fmt"
	"math"
)
//dmg = multiplier for the elements, 0-5 fire-heal
//TODO: Factor in damage enhance AND row enhance, just crossreference the awakening list, hardcode what each one does basically >_>, have it auto set # of enhanced orbs on web front end, user can change.
//Also, factor in active skill stuff
func damageResolve (team teamL/*[]lookup*/, teamD []teamDamage, dmg []float64, msg *orbs) (res []teamDamage) {
	//NOTE, dmg = multiplier for that colour.
	var comboMulti float64
	var comboCount float64
	
	comboMulti, comboCount=0,0

	tpacount := []float64{0,0,0,0,0}

	//two prong, ugh
	for _,y := range msg.Fire {
		if y == 4 {
			tpacount[0] ++
		} else if y > 2 {
		//if y == 0 { continue } (don't do any calcs)
		//Or better yet, if y > 2 (do damage stuff and combocount++)

			comboCount ++
			dmg[0] += 1 + (( y - 3)*0.25)
		}
	}
	for _,y := range msg.Water {
		if y == 4 {
			tpacount[1] ++
		} else if y > 2 {
			comboCount ++
			dmg[1] += 1 + (( y - 3)*0.25)
		}
	}
	for _,y := range msg.Wood {
		if y == 4 {
			tpacount[2] ++
		} else if y > 2 {
			comboCount ++
			dmg[2] += 1 + (( y - 3)*0.25)
		}
	}
	for _,y := range msg.Light {
		if y == 4 {
			tpacount[3] ++
		} else if y > 2 {
			comboCount ++
			dmg[3] += 1 + (( y - 3)*0.25)
		}
	}
	for _,y := range msg.Dark {
		if y == 4 {
			tpacount[4] ++
		} else if y > 2 {
			comboCount ++
			dmg[4] += 1 + (( y - 3)*0.25)
		}
	}
	for _,y := range msg.Heart {
		if y > 2 {
			comboCount ++
			dmg[5] += 1 + (( y - 3)*0.25)
		}
	}

	if comboCount >= 1 {
		comboMulti = 1 + (( comboCount - 1)*0.25)
	} else {
		comboMulti = 0
	}
	fmt.Println("Combo Multiplier:", comboMulti)
	//factor in leaderskill last
	var subMulti float64
	//var lead, friend bool
	subMulti = 0
	var numTpAwk float64
	
	for x,_ := range teamD {
		//Main attribute
		//lead,friend = false, false
		numTpAwk = 0
		teamD[x].Damage[0].Element = team.Team[x].Element
		for _, awk := range team.Team[x].Awakenings {
			if awk == twoProng {
				numTpAwk ++
			}
		}
		if team.Team[x].Element != nil {
			
			teamD[x].Damage[0].Value = ((tpacount[*team.Team[x].Element] * (1.25 * math.Pow(1.5, numTpAwk))) + dmg[*team.Team[x].Element]) * float64(team.Team[x].Stats.ATK) * comboMulti
		}
		//Sub attribute
		if team.Team[x].Element2 != nil{
			teamD[x].Damage[1].Element = team.Team[x].Element2
		} else {
			teamD[x].Damage[1].Element = nil
		}

		//Multiplier for sub element, 0.1 for same element, 0.3 if they differ.
		if team.Team[x].Element2 != nil{
			if *team.Team[x].Element == *team.Team[x].Element2 {
				subMulti = 0.10
			} else { subMulti = 0.30 }
			
			teamD[x].Damage[1].Value = ((tpacount[*team.Team[x].Element2] * (1.25 * math.Pow(1.5, numTpAwk))) + dmg[*team.Team[x].Element2]) * ( float64(team.Team[x].Stats.ATK)*subMulti ) * comboMulti
		}
		//Heal
		temp := 6
		teamD[x].Damage[2].Element = &temp
		teamD[x].Damage[2].Value = dmg[5] * float64(team.Team[x].Stats.RCV) * comboMulti
		//teamD
		//team[x]
		
		//Leader Skill
		//Re-implementing leaderskill
		switch msg.LeaderSkill.Condition[0].(string) {
		case "type":
			if msg.LeaderSkill.Condition[1].(float64) == float64(team.Team[x].Type) || msg.LeaderSkill.Condition[1].(float64) == float64(team.Team[x].Type2) {
				teamD[x].Damage[0].Value *= msg.LeaderSkill.ATK
				teamD[x].Damage[1].Value *= msg.LeaderSkill.ATK
				teamD[x].Damage[2].Value *= msg.LeaderSkill.RCV
			}
			
		case "elem":
			if teamD[x].Damage[0].Element != nil{
				if msg.LeaderSkill.Condition[1].(float64) == float64(*teamD[x].Damage[0].Element) {
					teamD[x].Damage[0].Value *= msg.LeaderSkill.ATK
				}
			}
			if teamD[x].Damage[1].Element != nil{
				if msg.LeaderSkill.Condition[1].(float64) == float64(*teamD[x].Damage[1].Element) {
					teamD[x].Damage[1].Value *= msg.LeaderSkill.ATK
				}
			}
		case "all", "default":
			teamD[x].Damage[0].Value *= msg.LeaderSkill.ATK
			teamD[x].Damage[1].Value *= msg.LeaderSkill.ATK
			teamD[x].Damage[2].Value *= msg.LeaderSkill.RCV			
		}

		//Friend leader skill!
		switch msg.FLeaderSkill.Condition[0].(string) {
		case "type":
			if msg.FLeaderSkill.Condition[1].(float64) == float64(team.Team[x].Type) || msg.FLeaderSkill.Condition[1].(float64) == float64(team.Team[x].Type2) {
				teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
				teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
				teamD[x].Damage[2].Value *= msg.FLeaderSkill.RCV
			}
		case "elem":
			if teamD[x].Damage[0].Element != nil{
				if msg.FLeaderSkill.Condition[1].(float64) == float64(*teamD[x].Damage[0].Element) {
					teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
				}
			}
			if teamD[x].Damage[1].Element != nil{
				if msg.FLeaderSkill.Condition[1].(float64) == float64(*teamD[x].Damage[1].Element) {
					teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
				}
			}
		case "all", "default":
			teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
			teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
			teamD[x].Damage[2].Value *= msg.FLeaderSkill.RCV			
		}
		
	}

	//fmt.Println(teamD[0].Damage[0].Value)

	fmt.Println(msg)
	//Row Multipliers
	//( 1 + ( 0.1 * n * r)) n = # rows, r = num awakenings
	for x, _ := range teamD {
		//for each in teamD, figure out how much the row multiplier affects., #rows = msg.Rows[element]
		//msg.Rows[teamD[x].Damage[0].Element] for main att // 1 for sub att
		if teamD[x].Damage[0].Element != nil{
			teamD[x].Damage[0].Value *= (1 + (0.1 * float64(msg.Rows[*teamD[x].Damage[0].Element]) * float64(team.Rows[*teamD[x].Damage[0].Element])))}
		if teamD[x].Damage[1].Element != nil{
			teamD[x].Damage[1].Value *= (1 + (0.1 * float64(msg.Rows[*teamD[x].Damage[1].Element]) * float64(team.Rows[*teamD[x].Damage[0].Element])))
		}
	} //test

	//Enhance orbs multiplier
	//(1 + ( 0.06 * n )) n = # enhanced orbs
	for x, _ := range teamD {
		if teamD[x].Damage[0].Element != nil{
			teamD[x].Damage[0].Value *= (1 + (0.06 * float64(msg.Enhance[*teamD[x].Damage[0].Element])))
		}
		if teamD[x].Damage[1].Element != nil{
			teamD[x].Damage[1].Value *= (1 + (0.06 * float64(msg.Enhance[*teamD[x].Damage[1].Element])))
		}
	}


	//Active skill multiplier. (Strict multiplier.. if no active skill, put [ "type/elem", 1, 1 ]
	for x, _ := range teamD {
		switch msg.Active[0].(string) {
		case "type":
			if team.Team[x].Type == msg.Active[1].(int) || team.Team[x].Type2 == msg.Active[1].(int) {
				teamD[x].Damage[0].Value *= float64(msg.Active[2].(int))
				teamD[x].Damage[1].Value *= float64(msg.Active[2].(int))
			}
		case "elem":
			if teamD[x].Damage[0].Element != nil{
				if *teamD[x].Damage[0].Element == int(msg.Active[1].(float64)){
					teamD[x].Damage[0].Value *= float64(msg.Active[2].(float64))
				}
			}
			if teamD[x].Damage[1].Element != nil{
				if *teamD[x].Damage[1].Element == int(msg.Active[1].(float64)){
					teamD[x].Damage[1].Value *= float64(msg.Active[2].(float64))
				}
			}
		}
	}


	fmt.Println(teamD[0].Damage[0].Value)
	res = teamD
	return
}
