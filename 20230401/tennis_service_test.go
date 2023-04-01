package tennis

import (
	"fmt"
	"math"
	"testing"
)

type TennisService struct {
	player1Score int
	player2Score int
}

func (s TennisService) GetScore() string {
	scoreMap := map[int]string{
		0: "love",
		1: "fifteen",
		2: "thirty",
		3: "forty",
	}
	if s.player1Score >= 3 && s.player2Score >= 3 {

		if s.player1Score == s.player2Score {
			return "deuce"
		}

		score := math.Abs(float64(s.player1Score - s.player2Score))
		if (s.player1Score > 3 || s.player2Score > 3) && score >= 1 {
			if s.player1Score > s.player2Score {
				return "joey win"
			}
			return "andy win"
		}
	}
	if s.player1Score != s.player2Score {
		return fmt.Sprintf("%s %s", scoreMap[s.player1Score], scoreMap[s.player2Score])
	}
	return fmt.Sprintf("%s all", scoreMap[s.player1Score])
}

func (s *TennisService) SetPlayer2Score(score int) {
	s.player2Score = score
}

func (s *TennisService) SetPlayer1Score(score int) {
	s.player1Score = score
}

func SetScore(player1, player2 int) TennisService {
	service := TennisService{}
	service.SetPlayer1Score(player1)
	service.SetPlayer2Score(player2)
	return service
}
func TestTennisService_GetScore(t *testing.T) {
	t.Run("Should Equal love all", func(t *testing.T) {
		service := SetScore(0, 0)
		if score := service.GetScore(); score != "love all" {
			t.Errorf("Should Equal love all, but get %s", score)
		}
	})

	t.Run("Should Equal fifteen love", func(t *testing.T) {
		service := SetScore(1, 0)
		if score := service.GetScore(); score != "fifteen love" {
			t.Errorf("Should Equal fifteen love, but get %s", score)
		}
	})

	t.Run("Should Equal thirty love", func(t *testing.T) {
		service := SetScore(2, 0)
		if score := service.GetScore(); score != "thirty love" {
			t.Errorf("Should Equal thirty love, but get %s", score)
		}
	})
	t.Run("Should Equal forty love", func(t *testing.T) {
		service := SetScore(3, 0)
		if score := service.GetScore(); score != "forty love" {
			t.Errorf("Should Equal forty love, but get %s", score)
		}
	})
	t.Run("Should Equal love fifteen", func(t *testing.T) {
		service := SetScore(0, 1)
		if score := service.GetScore(); score != "love fifteen" {
			t.Errorf("Should Equal love fifteen, but get %s", score)
		}
	})
	t.Run("Should Equal love thirty", func(t *testing.T) {
		service := SetScore(0, 2)
		if score := service.GetScore(); score != "love thirty" {
			t.Errorf("Should Equal love thirty, but get %s", score)
		}
	})
	t.Run("Should Equal love forty", func(t *testing.T) {
		service := SetScore(0, 3)
		if score := service.GetScore(); score != "love forty" {
			t.Errorf("Should Equal love forty, but get %s", score)
		}
	})
	t.Run("Should Equal fifteen all", func(t *testing.T) {
		service := SetScore(1, 1)
		if score := service.GetScore(); score != "fifteen all" {
			t.Errorf("Should Equal fifteen all, but get %s", score)
		}
	})
	t.Run("Should Equal thirty all", func(t *testing.T) {
		service := SetScore(2, 2)
		if score := service.GetScore(); score != "thirty all" {
			t.Errorf("Should Equal thirty all, but get %s", score)
		}
	})
	t.Run("Should Equal deuce", func(t *testing.T) {
		service := SetScore(3, 3)
		if score := service.GetScore(); score != "deuce" {
			t.Errorf("Should Equal deuce, but get %s", score)
		}
	})
	t.Run("Should Equal joey win", func(t *testing.T) {
		service := SetScore(4, 3)
		if score := service.GetScore(); score != "joey win" {
			t.Errorf("Should Equal joey win, but get %s", score)
		}
	})
	t.Run("Should Equal andy win", func(t *testing.T) {
		service := SetScore(3, 4)
		if score := service.GetScore(); score != "andy win" {
			t.Errorf("Should Equal andy win, but get %s", score)
		}
	})
	t.Run("Should Equal joey win", func(t *testing.T) {
		service := SetScore(3, 5)
		if score := service.GetScore(); score != "andy win" {
			t.Errorf("Should Equal joey win, but get %s", score)
		}
	})
}
