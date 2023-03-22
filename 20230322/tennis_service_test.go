package tennis

import (
	"fmt"
	"math"
	"testing"
)

type TennisService struct {
}

func (t TennisService) GetScore(player1 int, player2 int) string {
	scoreMap := map[int]string{
		0: "love",
		1: "fifteen",
		2: "thirty",
		3: "forty",
	}
	if player1 == player2 {
		if player1 >= 3 && player2 >= 3 {
			return "deuce"
		}
		return fmt.Sprintf("%s all", scoreMap[player1])
	}
	if player1 >= 3 && player2 >= 3 {
		suffix := "win"
		player := "andy"
		if player1 > player2 {
			player = "joey"
		}
		if math.Abs(float64(player2)-float64(player1)) < 2 {
			suffix = "adv"
		}
		return fmt.Sprintf("%s %s", player, suffix)
	}
	return fmt.Sprintf("%s %s", scoreMap[player1], scoreMap[player2])
}

func Test_Tennis_Service_ShouldEqual_LoveAll(t *testing.T) {
	service := TennisService{}
	if service.GetScore(0, 0) != "love all" {
		t.Errorf("Not Equal love all")
	}
}

func Test_Tennis_Service_ShouldEqual_fifteenAll(t *testing.T) {
	service := TennisService{}
	if service.GetScore(1, 0) != "fifteen love" {
		t.Errorf("Not Equal fifteen love")
	}
}

func Test_Tennis_Service_ShouldEqual_thirty_love(t *testing.T) {
	service := TennisService{}
	if service.GetScore(2, 0) != "thirty love" {
		t.Errorf("Not Equal fifteen love")
	}
}

func Test_Tennis_Service_ShouldEqual_forty_love(t *testing.T) {
	service := TennisService{}
	if service.GetScore(3, 0) != "forty love" {
		t.Errorf("Not Equal forty love")
	}
}

func Test_Tennis_Service_ShouldEqual_love_fifteen(t *testing.T) {
	service := TennisService{}
	if service.GetScore(0, 1) != "love fifteen" {
		t.Errorf("Not Equal love fifteen")
	}
}

func Test_Tennis_Service_ShouldEqual_love_thirty(t *testing.T) {
	service := TennisService{}
	if service.GetScore(0, 2) != "love thirty" {
		t.Errorf("Not Equal love thirty")
	}
}

func Test_Tennis_Service_ShouldEqual_love_forty(t *testing.T) {
	service := TennisService{}
	if service.GetScore(0, 3) != "love forty" {
		t.Errorf("Not Equal love forty")
	}
}

func Test_Tennis_Service_ShouldEqual_fifteen_all(t *testing.T) {
	service := TennisService{}
	if service.GetScore(1, 1) != "fifteen all" {
		t.Errorf("Not Equal fifteen all")
	}
}

func Test_Tennis_Service_ShouldEqual_thirty_all(t *testing.T) {
	service := TennisService{}
	if service.GetScore(2, 2) != "thirty all" {
		t.Errorf("Not Equal thirty all")
	}
}

func Test_Tennis_Service_ShouldEqual_deuce(t *testing.T) {
	service := TennisService{}
	if service.GetScore(3, 3) != "deuce" {
		t.Errorf("Not Equal deuce")
	}
}

func Test_Tennis_Service_ShouldEqual_Joey_Adv(t *testing.T) {
	service := TennisService{}
	if service.GetScore(4, 3) != "joey adv" {
		t.Errorf("Not Equal joey adv")
	}
}

func Test_Tennis_Service_ShouldEqual_adny_Adv(t *testing.T) {
	service := TennisService{}
	if service.GetScore(3, 4) != "andy adv" {
		t.Errorf("Not Equal andy adv")
	}
}

func Test_Tennis_Service_ShouldEqual_andy_win(t *testing.T) {
	service := TennisService{}
	if service.GetScore(3, 5) != "andy win" {
		t.Errorf("Not Equal andy win")
	}
}
