data Player = Player { player_id :: String, real_rating :: Int, glicko_rating :: Int } deriving Show

data Team = Team { team_id :: String, players :: [Player] } deriving Show

alice = Player "123" 10 12
bob   = Player "124" 5 13

teamA = Team "team_a" [alice]
teamB = Team "team_b" [bob]

calculateTeamRating :: Team -> Int
calculateTeamRating team = sum (map real_rating (players team))

updateRatings :: Team -> Int -> Team
updateRatings team change = 
    team { players = map updatePlayer (players team) }
  where
    updatePlayer p = p { glicko_rating = glicko_rating p + change }

calculateRatingSimple :: Team -> Team -> (Team, Team)
calculateRatingSimple team1 team2
  | calculateTeamRating team1 >= calculateTeamRating team2 = (updateRatings team1 25, updateRatings team2 (-25))
  | otherwise                                              = (updateRatings team2 25, updateRatings team1 (-25))
