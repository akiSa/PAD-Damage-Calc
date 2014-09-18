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

	tpacount := []int{}

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
	var lead, friend bool
	subMulti = 0
	for x,_ := range teamD {
		//Main attribute
		lead,friend = false, false
		teamD[x].Damage[0].Element = team.Team[x].Element
		teamD[x].Damage[0].Value = dmg[team.Team[x].Element] * float64(team.Team[x].Stats.ATK) * comboMulti
		//Sub attribute
		teamD[x].Damage[1].Element = team.Team[x].Element2
		if team.Team[x].Element == team.Team[x].Element2 { subMulti = 0.10 } else { subMulti = 0.30 }
		teamD[x].Damage[1].Value = dmg[team.Team[x].Element2] * (float64(team.Team[x].Stats.ATK)*subMulti) * comboMulti
		//Heal
		teamD[x].Damage[2].Element = 6
		teamD[x].Damage[2].Value = dmg[5] * float64(team.Team[x].Stats.RCV) * comboMulti
		//teamD
		//team[x]
		
		//Leader Skill
		if len(team.Team[0].LeaderSkill.Conditional) >= 1 {
			switch team.Team[0].LeaderSkill.Conditional[0] {
			case "type":
				//Apply only to conditional[1] types
				if team.Team[x].Type == team.Team[0].LeaderSkill.Conditional[1] ||
					team.Team[x].Type2 == team.Team[0].LeaderSkill.Conditional[1] {
					if !lead {				
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						lead = true
					}
					//I THINK it works like this, need to check, but I remember d/l batman enhancing his own dark attacks too
				}
			case "elem":
				//Apply only to conditional[1] elements
				if team.Team[x].Element == team.Team[0].LeaderSkill.Conditional[1] ||
					team.Team[x].Element2 == team.Team[0].LeaderSkill.Conditional[1] {
					
					if !lead {
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						lead = true
					}
				}
			default:
				//Apply to errbody
				if !lead {
					teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
					lead = true
				}
			}
		}
		if len(team.Team[0].LeaderSkill.Conditional2) >= 1 {
			switch team.Team[0].LeaderSkill.Conditional2[0] {
			case "type":
				//Apply only to conditional[1] types
				if team.Team[x].Type == team.Team[0].LeaderSkill.Conditional2[1] ||
					team.Team[x].Type2 == team.Team[0].LeaderSkill.Conditional2[1] {
					if !lead {				
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						lead = true
					}
					//I THINK it works like this, need to check, but I remember d/l batman enhancing his own dark attacks too
				}
			case "elem":
				//Apply only to conditional[1] elements
				if team.Team[x].Element == team.Team[0].LeaderSkill.Conditional2[1] ||
					team.Team[x].Element2 == team.Team[0].LeaderSkill.Conditional2[1] {
					
					if !lead {
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						lead = true
					}
				}
			default:
				//this'll never get called, but i'll leave it here just in case.
				if !lead {
					teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
					lead = true
				}
			}
		}
		
		//******************** FRIEND *********************
		if len(team.Team[5].LeaderSkill.Conditional) >= 1 {
			switch team.Team[5].LeaderSkill.Conditional[0] {
			case "type":
				//Apply only to conditional[1] types
				if team.Team[x].Type == team.Team[5].LeaderSkill.Conditional[1] ||
					team.Team[x].Type2 == team.Team[5].LeaderSkill.Conditional[1] {
					if !friend {				
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						friend = true
					}
				}
			case "elem":
				//Apply only to conditional[1] elements
				if team.Team[x].Element == team.Team[5].LeaderSkill.Conditional[1] ||
					team.Team[x].Element2 == team.Team[5].LeaderSkill.Conditional[1] {
					
					if !friend {
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						friend = true
					}
				}
			default:
				//Apply to errbody
				if !friend {
					teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
					friend = true
				}
			}
		}
		//FRIEND CONDITIONAL 2
		if len(team.Team[5].LeaderSkill.Conditional2) >= 1 {
			switch team.Team[5].LeaderSkill.Conditional2[0] {
			case "type":
				//Apply only to conditional[1] types
				if team.Team[x].Type == team.Team[5].LeaderSkill.Conditional2[1] ||
					team.Team[x].Type2 == team.Team[5].LeaderSkill.Conditional2[1] {
					if !friend {				
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						friend = true
					}
				}
			case "elem":
				//Apply only to conditional[1] elements
				if team.Team[x].Element == team.Team[5].LeaderSkill.Conditional2[1] ||
					team.Team[x].Element2 == team.Team[5].LeaderSkill.Conditional2[1] {
					
					if !friend {
						teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
						teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
						friend = true
					}
				}
			default:
				//Apply to errbody
				if !friend {
					teamD[x].Damage[0].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team.Team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team.Team[0].LeaderSkill.RCV
					friend = true
				}
			}
		}
	}

	fmt.Println(msg)
	//Row Multipliers
	//( 1 + ( 0.1 * n * r)) n = # rows, r = num awakenings
	for x, _ := range teamD {
		//for each in teamD, figure out how much the row multiplier affects., #rows = msg.Rows[element]
		//msg.Rows[teamD[x].Damage[0].Element] for main att // 1 for sub att
		teamD[x].Damage[0].Value *= (1 + (0.1 * float64(msg.Rows[teamD[x].Damage[0].Element]) * float64(team.Rows[teamD[x].Damage[0].Element])))
		teamD[x].Damage[1].Value *= (1 + (0.1 * float64(msg.Rows[teamD[x].Damage[1].Element]) * float64(team.Rows[teamD[x].Damage[0].Element])))
	} //test

	//Enhance orbs multiplier
	//(1 + ( 0.06 * n )) n = # enhanced orbs
	for x, _ := range teamD {
		teamD[x].Damage[0].Value *= (1 + (0.06 * float64(msg.Enhance[teamD[x].Damage[0].Element])))
		teamD[x].Damage[1].Value *= (1 + (0.06 * float64(msg.Enhance[teamD[x].Damage[1].Element])))
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
			if teamD[x].Damage[0].Element == int(msg.Active[1].(float64)){
				teamD[x].Damage[0].Value *= float64(msg.Active[2].(float64))
			}
			if teamD[x].Damage[1].Element == int(msg.Active[1].(float64)){
				teamD[x].Damage[1].Value *= float64(msg.Active[2].(float64))
			}
		}
	}

	//TWO PRONG ATTACk
	for x, _ := range teamD {
		numawk := 0
		for _, awk := range team.Team[x].Awakenings {
			if awk == twoProng {
				numawk ++
			}
		}
		if numawk >= 1 {
			teamD[x].Damage[0].Value += ( 1 + ( float64(0.5) * float64(numawk) * float64(tpacount[team.Team[x].Element])))
			teamD[x].Damage[1].Value += ( 1 + ( float64(0.5) * float64(numawk) * float64(tpacount[team.Team[x].Element2])))
		}
	}
	
	res = teamD
	return
}
