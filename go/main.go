package main

import "fmt"

type Player struct {
  Id string
  TrueSkill int
  GlickoRating int
}

type Team struct {
  Id string
  Players []Player
}

func getTeamTrueSkill(t Team) int {
  total := 0
  for _, player := range t.Players {
    total += player.TrueSkill
  }
  return total
}

func game(t1, t2 Team) []Player {
  res := []Player{}
  if (getTeamTrueSkill(t1) >= getTeamTrueSkill(t1)) {
    for _, player := range t1.Players {
      res = append(res, Player{Id: player.Id, TrueSkill: player.TrueSkill, GlickoRating: player.GlickoRating + 25})
    }
    for _, player := range t2.Players {
      res = append(res, Player{Id: player.Id, TrueSkill: player.TrueSkill, GlickoRating: player.GlickoRating - 25})
    }
  } else {
    for _, player := range t1.Players {
      res = append(res, Player{Id: player.Id, TrueSkill: player.TrueSkill, GlickoRating: player.GlickoRating - 25})
    }
    for _, player := range t2.Players {
      res = append(res, Player{Id: player.Id, TrueSkill: player.TrueSkill, GlickoRating: player.GlickoRating + 25})
    }
  }
  return res
}


func main() {
  fmt.Println("start")

  teamA := Team{
    Id: "Radiant",
    Players: []Player{
      Player{Id: "a_pos1", TrueSkill: 4000, GlickoRating: 4200},
      Player{Id: "a_pos2", TrueSkill: 4500, GlickoRating: 3800},
      Player{Id: "a_pos3", TrueSkill: 2000, GlickoRating: 3500},
      Player{Id: "a_pos4", TrueSkill: 3000, GlickoRating: 3500},
      Player{Id: "a_pos5", TrueSkill: 3000, GlickoRating: 3500},
    },
  }
  teamB := Team{
    Id: "Dire",
    Players: []Player{
      Player{Id: "b_pos1", TrueSkill: 5000, GlickoRating: 5200},
      Player{Id: "b_pos2", TrueSkill: 5500, GlickoRating: 4800},
      Player{Id: "b_pos3", TrueSkill: 3000, GlickoRating: 4500},
      Player{Id: "b_pos4", TrueSkill: 4000, GlickoRating: 4500},
      Player{Id: "b_pos5", TrueSkill: 4000, GlickoRating: 4500},
    },
  }


  fmt.Printf("%v\n", teamA)
  fmt.Printf("%v\n", teamB)

  fmt.Println("--- first match ---")
  newPlayers := game(teamA, teamB)

  fmt.Printf("%v\n", newPlayers)

  
}
