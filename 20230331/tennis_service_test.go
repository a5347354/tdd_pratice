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

func (s *TennisService) SetPlayer1Score(score int) {
	s.player1Score = score
}

func (s *TennisService) GetScore() string {
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
	}
	if s.player1Score > 3 || s.player2Score > 3 {
		score := math.Abs(float64(s.player1Score - s.player2Score))
		if score > 1 && s.player1Score > s.player2Score {
			return "joey win"
		}
		return "andy win"
	}
	if s.player1Score == s.player2Score {
		return fmt.Sprintf("%s all", scoreMap[s.player1Score])
	}
	return fmt.Sprintf("%s %s", scoreMap[s.player1Score], scoreMap[s.player2Score])
}

func (s *TennisService) SetPlayer2Score(score int) {
	s.player2Score = score
}

func TestTennisService(t *testing.T) {
	service := TennisService{}
	t.Run("Should Equal love all", func(t *testing.T) {
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(0)
		if score := service.GetScore(); score != "love all" {
			t.Errorf("Shoud Equal love all")
		}
	})

	t.Run("Should Equal fifteen love", func(t *testing.T) {
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(0)
		if score := service.GetScore(); score != "fifteen love" {
			t.Errorf("Shoud Equal fifteen love")
		}
	})

	t.Run("Should Equal fifteen all", func(t *testing.T) {
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(1)
		if score := service.GetScore(); score != "fifteen all" {
			t.Errorf("Shoud Equal fifteen all, but %s", score)
		}
	})

	t.Run("Should Equal thirty fifteen", func(t *testing.T) {
		service.SetPlayer1Score(2)
		service.SetPlayer2Score(1)
		if score := service.GetScore(); score != "thirty fifteen" {
			t.Errorf("Shoud Equal thirty fifteen, but %s", score)
		}
	})

	t.Run("Should Equal forty fifteen", func(t *testing.T) {
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(1)
		if score := service.GetScore(); score != "forty fifteen" {
			t.Errorf("Shoud Equal forty fifteen, but %s", score)
		}
	})

	t.Run("Should Equal joey win", func(t *testing.T) {
		service.SetPlayer1Score(4)
		service.SetPlayer2Score(1)
		if score := service.GetScore(); score != "joey win" {
			t.Errorf("Shoud Equal joey win, but %s", score)
		}
	})

	t.Run("Should Equal deuce", func(t *testing.T) {
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(3)
		if score := service.GetScore(); score != "deuce" {
			t.Errorf("Shoud Equal deuce, but %s", score)
		}
	})

	t.Run("Should Equal Andy Win", func(t *testing.T) {
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(4)
		if score := service.GetScore(); score != "andy win" {
			t.Errorf("Shoud Equal andy win, but %s", score)
		}

	})

	t.Run("Should Equal Andy Win", func(t *testing.T) {
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(5)
		if score := service.GetScore(); score != "andy win" {
			t.Errorf("Shoud Equal andy win, but %s", score)
		}
	})
}
