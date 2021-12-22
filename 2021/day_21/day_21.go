package day_21

import "fmt"

type deterministicDie struct {
	current int
	total   int
}

type player struct {
	current int
	score   int
}

func playGame(p1, p2 int) int {
	player1 := &player{current: p1}
	player2 := &player{current: p2}
	die := deterministicDie{0, 0}
	for {
		val := die.roll3()
		player1.move(val)
		if player1.score >= 1000 {
			break
		}
		val2 := die.roll3()
		player2.move(val2)
		if player2.score >= 1000 {
			break
		}
	}
	fmt.Println(player1)
	fmt.Println(player2)
	lowest := min(player1.score, player2.score)
	score := lowest * die.total
	return score
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func (p *player) move(steps int) {
	new := (p.current + steps) % 10
	if new == 0 {
		new = 10
	}
	p.score += new
	p.current = new
}

func (d *deterministicDie) roll3() int {
	thrown := []int{}
	for a := 0; a < 3; a++ {
		d.current = (d.current + 1) % 100
		if d.current == 0 {
			d.current = 100
		}
		thrown = append(thrown, d.current)
	}
	d.total += 3

	return total(thrown)
}

// p1 - 1,2, 3 3*2 = 6
// p2 - 4,5,6 = 3*5 = 15
// p1- 7,8,9 = 3*8 = 24

func total(a []int) (t int) {
	for _, v := range a {
		t += v
	}
	return
}

// 1 turn = 3 rolls is 27 universes
// 3,3x4,
// player has max of 10 positions
// player has max of 21 scores
// players{pos:4,score:0}:1
// 27-> players{7,3}:1, {8,4}:3, {9,5}:6, {10,6}:1, {1,1}:1, 2,2:1, 3,3:1, 4,4:1, 5,5:1, 6,6:1, 7,7:1

type playerD struct {
	players map[player]int
}

func rollsDirac() map[int]int {
	valTimes := map[int]int{}
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				total := a + b + c
				valTimes[total]++
			}
		}
	}
	return valTimes
}

func (p *playerD) wins() (bool, int) {
	wins := 0
	count := 0
	for pl, a := range p.players {
		if pl.score > 21 {
			count++
			wins += a
		}
	}
	if count >= len(p.players) {
		return true, wins
	}
	return false, wins
}
func (p *playerD) move() {
	rolls := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	// rolls := map[int]int{1: 1, 2: 1, 3: 1}
	newPlayers := map[player]int{}
	for pl, amount := range p.players {
		// position := pl.current
		// score := pl.score
		for val, am := range rolls {
			cur, sc := calMove(pl, val)
			n := player{cur, sc}
			// fmt.Printf("%+v a: %d\n", n, am*amount)
			newPlayers[n] += am * amount
		}
	}
	p.players = newPlayers
	// fmt.Printf("%+v\n", p.players)
}

func calMove(p player, steps int) (int, int) {
	new := (p.current + steps) % 10
	if new == 0 {
		new = 10
	}
	return new, (p.score + new)
}

func playDirac(p1, p2 int) (int, int) {
	player1 := &playerD{players: map[player]int{{p1, 0}: 1}}
	player2 := &playerD{players: map[player]int{{p2, 0}: 1}}
	for {
		player1.move()
		if win, _ := player1.wins(); win {
			break
		}
		player2.move()

		// fmt.Printf("%+v\n", player2.players)
		if win, _ := player2.wins(); win {
			break
		}
	}
	// fmt.Printf("%+v, %d\n", player2.wins())
	// score := 0
	_, score1 := player1.wins()
	_, score2 := player2.wins()
	return score1, score2
}
