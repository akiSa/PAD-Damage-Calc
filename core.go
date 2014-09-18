package main

import (
	"fmt"
)
//dmg = multiplier for the elements, 0-5 fire-heal
//TODO: Factor in damage enhance AND row enhance, just crossreference the awakening list, hardcode what each one does basically >_>, have it auto set # of enhanced orbs on web front end, user can change.
//Also, factor in active skill stuff
func damageResolve (team teamL/*[]lookup*/, teamD []teamDamage, dmg []float64, msg *orbs) (res []teamDamage) {
	var comboMulti float64
	var comboCount float64
	
	comboMulti, comboCount=0,0

	tpacount := []int{0,0,0,0,0}

	//two prong, ugh
	for _,y := range msg.Fire {
		if y == 4 { tpacount[0] ++ }
		comboCount ++
		dmg[0] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Water {
		if y == 4 { tpacount[1] ++ }
		comboCount ++
		dmg[1] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Wood {
		if y == 4 { tpacount[2] ++ }
		comboCount ++
		dmg[2] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Light {
		if y == 4 { tpacount[3] ++ }
		comboCount ++
		dmg[3] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Dark {
		if y == 4 { tpacount[4] ++ }
		comboCount ++
		dmg[4] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Heart {
		comboCount ++
		dmg[5] += 1 + (( y - 3)*0.25)
	}

	comboMulti = 1 + (( comboCount - 1)*0.25)
	fmt.Println("Combo Multiplier:", comboMulti)
	//factor in leaderskill last
	var subMulti float64
	//var lead, friend bool
	subMulti = 0
	for x,_ := range teamD {
		//Main attribute
		//lead,friend = false, false
		teamD[x].Damage[0].Element = team.Team[x].Element
		if team.Team[x].Element != nil {
			teamD[x].Damage[0].Value = dmg[*team.Team[x].Element] * float64(team.Team[x].Stats.ATK) * comboMulti
		}
		//Sub attribute
		if team.Team[x].Element2 != nil{
			teamD[x].Damage[1].Element = team.Team[x].Element2
		} else {
			teamD[x].Damage[1].Element = nil
		}
		if team.Team[x].Element == team.Team[x].Element2 { subMulti = 0.10 } else { subMulti = 0.30 }
		if team.Team[x].Element2 != nil{
			teamD[x].Damage[1].Value = dmg[*team.Team[x].Element2] * (float64(team.Team[x].Stats.ATK)*subMulti) * comboMulti
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
			if msg.LeaderSkill.Condition[1].(int) == team.Team[x].Type || msg.LeaderSkill.Condition[1].(int) == team.Team[x].Type2 {
				teamD[x].Damage[0].Value *= msg.LeaderSkill.ATK
				teamD[x].Damage[1].Value *= msg.LeaderSkill.ATK
				teamD[x].Damage[2].Value *= msg.LeaderSkill.RCV
			}
			
		case "elem":
			if teamD[x].Damage[0].Element != nil{
				if msg.LeaderSkill.Condition[1].(int) == *teamD[x].Damage[0].Element {
					teamD[x].Damage[0].Value *= msg.LeaderSkill.ATK
				}
			}
			if teamD[x].Damage[1].Element != nil{
				if msg.LeaderSkill.Condition[1].(int) == *teamD[x].Damage[1].Element {
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
			if msg.FLeaderSkill.Condition[1].(int) == team.Team[x].Type || msg.FLeaderSkill.Condition[1].(int) == team.Team[x].Type2 {
				teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
				teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
				teamD[x].Damage[2].Value *= msg.FLeaderSkill.RCV
			}
		case "elem":
			if teamD[x].Damage[0].Element != nil{
				if msg.FLeaderSkill.Condition[1].(int) == *teamD[x].Damage[0].Element {
					teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
				}
			}
			if teamD[x].Damage[1].Element != nil{
				if msg.FLeaderSkill.Condition[1].(int) == *teamD[x].Damage[1].Element {
					teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
				}
			}
		case "all", "default":
			teamD[x].Damage[0].Value *= msg.FLeaderSkill.ATK
			teamD[x].Damage[1].Value *= msg.FLeaderSkill.ATK
			teamD[x].Damage[2].Value *= msg.FLeaderSkill.RCV			
		}
		
	}

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

	//TWO PRONG ATTACk
	fmt.Print("TWO PRONG: ")
	fmt.Println(tpacount)
	for x, _ := range teamD {
		numawk := 0
		for _, awk := range team.Team[x].Awakenings {
			if awk == twoProng {
				numawk ++
			}
		}
		if numawk >= 1 {
			if team.Team[x].Element != nil{
				teamD[x].Damage[0].Value += ( 1 + ( float64(0.5) * float64(numawk) * float64(tpacount[*team.Team[x].Element])))
			}
			if team.Team[x].Element2 != nil {
				teamD[x].Damage[1].Value += ( 1 + ( float64(0.5) * float64(numawk) * float64(tpacount[*team.Team[x].Element2])))
			}
		}
	}
	
	res = teamD
	return
}
