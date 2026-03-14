package main

import "fmt"

type Player struct {
  Id string
  RealRating int
  GlickoRating int
}

type Team struct {
  Id string
  Players []Player
}

func getTeamRealRating(t Team) int {
  total := 0
  for _, player := range t.Players {
    total += player.RealRating
  }
  return total
}

func game(t1, t2 Team) []Player {
  res := []Player{}
  if (getTeamRealRating(t1) >= getTeamRealRating(t1)) {
    for _, player := range t1.Players {
      res = append(res, Player{Id: player.Id, RealRating: player.RealRating, GlickoRating: player.GlickoRating + 25})
    }
    for _, player := range t2.Players {
      res = append(res, Player{Id: player.Id, RealRating: player.RealRating, GlickoRating: player.GlickoRating - 25})
    }
  } else {
    for _, player := range t1.Players {
      res = append(res, Player{Id: player.Id, RealRating: player.RealRating, GlickoRating: player.GlickoRating - 25})
    }
    for _, player := range t2.Players {
      res = append(res, Player{Id: player.Id, RealRating: player.RealRating, GlickoRating: player.GlickoRating + 25})
    }
  }
  return res
}


func main() {
  fmt.Println("start")

  teamA := Team{
    Id: "Radiant",
    Players: []Player{
      Player{Id: "a_pos1", RealRating: 4000, GlickoRating: 4200},
      Player{Id: "a_pos2", RealRating: 4500, GlickoRating: 3800},
      Player{Id: "a_pos3", RealRating: 2000, GlickoRating: 3500},
      Player{Id: "a_pos4", RealRating: 3000, GlickoRating: 3500},
      Player{Id: "a_pos5", RealRating: 3000, GlickoRating: 3500},
    },
  }
  teamB := Team{
    Id: "Dire",
    Players: []Player{
      Player{Id: "b_pos1", RealRating: 5000, GlickoRating: 5200},
      Player{Id: "b_pos2", RealRating: 5500, GlickoRating: 4800},
      Player{Id: "b_pos3", RealRating: 3000, GlickoRating: 4500},
      Player{Id: "b_pos4", RealRating: 4000, GlickoRating: 4500},
      Player{Id: "b_pos5", RealRating: 4000, GlickoRating: 4500},
    },
  }


  fmt.Printf("%v\n", teamA)
  fmt.Printf("%v\n", teamB)

  fmt.Println("--- first match ---")
  newPlayers := game(teamA, teamB)

  fmt.Printf("%v\n", newPlayers)

  
}
