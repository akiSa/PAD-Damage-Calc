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
	for x, y := range msg.Orbs {
		for _, z := range y {
			if x != 5 && z[0] == 4 { //heal orbs can't tpa silly
				tpacount[x]++
				comboCount ++
			} else {
				if z[0] > 2 {
					comboCount ++
					fmt.Println(z)
					dmg[x] += 1 + (( z[0] -3) * 0.25) + (0.06 * z[1])
				}
			}
		}
	}

	if comboCount >= 1 {
		comboMulti = 1 + (( comboCount - 1)*0.25)
	} else {
		comboMulti = 0
	}
	fmt.Println("Combo Multiplier:", comboMulti)
	fmt.Println("TPA counter:", tpacount)
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
			//(# of 4 matches * (1.25 * 1.5^#num tpa awakenings) + multiplier without all the tpa's) * atk * combo multiplier
			//Get TPA out of the way
			teamD[x].Damage[0].Value = 0
			for _, y := range msg.Orbs[*team.Team[x].Element] {
				//Find the TPA matches
				if y[0] == 4  && numTpAwk > 0 {
					teamD[x].Damage[0].Value += (1.25 * math.Pow(1.5, numTpAwk) + (0.06 * y[1]))
				} else if y[0] == 4  && numTpAwk == 0 {
					teamD[x].Damage[0].Value += (1 + ((y[0] - 3) * 0.25 )) + (0.06 * y[1])
				}
			}

			teamD[x].Damage[0].Value += dmg[*team.Team[x].Element] // remember, dmg is the multiplier
			teamD[x].Damage[0].Value *= float64(team.Team[x].Stats.ATK) * comboMulti
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

			teamD[x].Damage[1].Value = 0
			for _, y := range msg.Orbs[*team.Team[x].Element2] {
				//Find the TPA matches
				if y[0] == 4 && numTpAwk > 0 {
					teamD[x].Damage[1].Value += (1.25 * math.Pow(1.5, numTpAwk) + (0.06 * y[1]))
				} else if y[0] == 4  && numTpAwk == 0 {
					teamD[x].Damage[1].Value += (1 + ((y[0] - 3) * 0.25 )) + (0.06 * y[1])
				}
			}

			teamD[x].Damage[1].Value += dmg[*team.Team[x].Element2]
			teamD[x].Damage[1].Value *= ( float64(team.Team[x].Stats.ATK) * subMulti ) * comboMulti

		}
		//Heal
		temp := 5
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

	//Active skill multiplier. (Strict multiplier.. if no active skill, put [ "type/elem", 1, 1 ]
	for x, _ := range teamD {
		switch msg.Active[0].(string) {
		case "type":
			if team.Team[x].Type == int(msg.Active[1].(float64)) || team.Team[x].Type2 == int(msg.Active[1].(float64)) {
				teamD[x].Damage[0].Value *= float64(msg.Active[2].(float64))
				teamD[x].Damage[1].Value *= float64(msg.Active[2].(float64))
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

	res = teamD
	return
}
