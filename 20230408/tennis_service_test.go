package tennis

import (
	"fmt"
	"math"
	"testing"
)

type TennisService struct {
	Player1Score int
	Player2Score int
}

func (s *TennisService) SetPlayer2Score(score int) {
	s.Player2Score = score
}

func (s *TennisService) SetPlayer1Score(score int) {
	s.Player1Score = score
}

func (s *TennisService) GetScore() string {
	scoreMap := map[int]string{
		0: "love",
		1: "fifteen",
		2: "thirty",
		3: "forty",
	}
	if s.Player1Score == s.Player2Score {
		if s.Player1Score >= 3 && s.Player2Score >= 3 {
			return "deuce"
		}
		return fmt.Sprintf("%s all", scoreMap[s.Player1Score])
	}
	if s.Player1Score >= 3 || s.Player2Score >= 3 {
		abs := math.Abs(float64(s.Player1Score - s.Player2Score))
		if abs == 1 {
			if s.Player1Score > s.Player2Score {
				return "joey adv"
			}
			return "andy adv"
		}
		if abs == 2 {
			if s.Player1Score > s.Player2Score {
				return "joey win"
			}
			return "andy win"
		}
	}
	return fmt.Sprintf("%s %s", scoreMap[s.Player1Score], scoreMap[s.Player2Score])
}

func TestTennisService(t *testing.T) {
	t.Run("Should Equal Love All", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(0)
		if service.GetScore() != "love all" {
			t.Errorf("Should Equal love all, but get %v", service.GetScore())
		}
	})
	t.Run("Should Equal fifteen love", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(0)
		if service.GetScore() != "fifteen love" {
			t.Errorf("Should Equal fifteen love, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal thirty love", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(2)
		service.SetPlayer2Score(0)
		if service.GetScore() != "thirty love" {
			t.Errorf("Should Equal thirty love, but get %v", service.GetScore())
		}
	})
	t.Run("Should Equal forty love", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(0)
		if service.GetScore() != "forty love" {
			t.Errorf("Should Equal forty love, but get %v", service.GetScore())
		}
	})
	t.Run("Should Equal love fifteen", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(1)
		if service.GetScore() != "love fifteen" {
			t.Errorf("Should Equal love fifteen, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal love thirty", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(2)
		if service.GetScore() != "love thirty" {
			t.Errorf("Should Equal love thirty, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal love forty", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(3)
		if service.GetScore() != "love forty" {
			t.Errorf("Should Equal love forty, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal fifteen all", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(1)
		if service.GetScore() != "fifteen all" {
			t.Errorf("Should Equal fifteen all, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal thirty all", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(2)
		service.SetPlayer2Score(2)
		if service.GetScore() != "thirty all" {
			t.Errorf("Should Equal thirty all, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal deuce", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(3)
		if service.GetScore() != "deuce" {
			t.Errorf("Should Equal deuce, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal joey adv", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(4)
		service.SetPlayer2Score(3)
		if service.GetScore() != "joey adv" {
			t.Errorf("Should Equal joey adv, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal andy adv", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(4)
		if service.GetScore() != "andy adv" {
			t.Errorf("Should Equal andy adv, but get %v", service.GetScore())
		}
	})

	t.Run("Should Equal andy win", func(t *testing.T) {
		service := TennisService{}
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(5)
		if service.GetScore() != "andy win" {
			t.Errorf("Should Equal andy win, but get %v", service.GetScore())
		}
	})
}
