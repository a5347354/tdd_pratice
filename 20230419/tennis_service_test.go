package tennis_service

import (
	"fmt"
	"math"
	"testing"
)

func TestTennis_Service(t *testing.T) {
	t.Run("Should eqal love all ", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(0)
		if service.GetScore() != "love all" {
			t.Errorf("Should eqal love all, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal fifteen love", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(0)
		if service.GetScore() != "fifteen love" {
			t.Errorf("Should eqal fifteen love, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal thirty love", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(2)
		service.SetPlayer2Score(0)
		if service.GetScore() != "thirty love" {
			t.Errorf("Should eqal thirty love, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal forty love", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(0)
		if service.GetScore() != "forty love" {
			t.Errorf("Should eqal forty love, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal love fifteen", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(1)
		if service.GetScore() != "love fifteen" {
			t.Errorf("Should eqal love fifteen, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal love thirty", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(2)
		if service.GetScore() != "love thirty" {
			t.Errorf("Should eqal love thirty, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal love forty", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(0)
		service.SetPlayer2Score(3)
		if service.GetScore() != "love forty" {
			t.Errorf("Should eqal love forty, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal fifteen all", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(1)
		service.SetPlayer2Score(1)
		if service.GetScore() != "fifteen all" {
			t.Errorf("Should eqal fifteen all, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal thirty all", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(2)
		service.SetPlayer2Score(2)
		if service.GetScore() != "thirty all" {
			t.Errorf("Should eqal thirty all, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal deuce", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(3)
		if service.GetScore() != "deuce" {
			t.Errorf("Should eqal deuce, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal joey adv", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(4)
		service.SetPlayer2Score(3)
		if service.GetScore() != "joey adv" {
			t.Errorf("Should eqal joey adv, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal andy adv", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(4)
		if service.GetScore() != "andy adv" {
			t.Errorf("Should eqal andy adv, but eqal %s", service.GetScore())
		}
	})

	t.Run("Should eqal andy win", func(t *testing.T) {
		service := NewTennisService()
		service.SetPlayer1Score(3)
		service.SetPlayer2Score(5)
		if service.GetScore() != "andy win" {
			t.Errorf("Should eqal andy win, but eqal %s", service.GetScore())
		}
	})
}

type TennisService struct {
	Player1Score int
	Player2Score int
}

func (s *TennisService) SetPlayer1Score(score int) {
	s.Player1Score = score
}

func (s *TennisService) SetPlayer2Score(score int) {
	s.Player2Score = score
}

func (s *TennisService) GetScore() string {
	scoreMap := map[int]string{
		0: "love",
		1: "fifteen",
		2: "thirty",
		3: "forty",
	}
	if s.Player1Score >= 3 || s.Player2Score >= 3 {
		if s.Player1Score == s.Player2Score {
			return "deuce"
		}
		if math.Abs(float64(s.Player1Score-s.Player2Score)) == 1 {
			if s.Player1Score > s.Player2Score {
				return "joey adv"
			}
			return "andy adv"
		}
		if math.Abs(float64(s.Player1Score-s.Player2Score)) == 2 {
			if s.Player1Score > s.Player2Score {
				return "joey win"
			}
			return "andy win"
		}
	}
	if s.Player1Score == s.Player2Score {
		return fmt.Sprintf("%s all", scoreMap[s.Player1Score])
	}
	return fmt.Sprintf("%s %s", scoreMap[s.Player1Score], scoreMap[s.Player2Score])
}

func NewTennisService() TennisService {
	return TennisService{}
}
