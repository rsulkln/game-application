package gameserviceclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Cleint struct {
	// game service url //game-app/game
	URL string
}

func (c Cleint) TotalScore(PlayerID uint) int {
	response, _ := http.Get(c.URL + "/total-score?player_id=" + strconv.Itoa(int(PlayerID)))

	type TotalScoreResponse struct {
		TotalScore int `json:"total_score"`
	}

	var tsr TotalScoreResponse
	data, rErr := io.ReadAll(response.Body)
	if rErr != nil {
		fmt.Errorf("error to read body hhtp request: %w", rErr)

	}
	json.Unmarshal(data, &tsr)

	return tsr.TotalScore
}
