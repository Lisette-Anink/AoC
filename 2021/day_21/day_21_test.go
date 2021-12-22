package day_21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func useTestData() (int, int) {
	// input := utils.ImportFileLines("test_input_21")
	// lines := utils.SplitInLines(input)
	return 4, 8
}

func useInputData() (int, int) {
	// input := utils.ImportFromAoC("2021", "21")
	// lines := utils.SplitInLines(input)
	return 10, 9
}

func Test_deterministicDie_roll3(t *testing.T) {
	type fields struct {
		current int
		total   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want2  *deterministicDie
	}{
		{"test 0", fields{0, 0}, 6, &deterministicDie{3, 3}},
		{"test 80", fields{80, 0}, 246, &deterministicDie{83, 3}},
		{"test 98", fields{98, 0}, 200, &deterministicDie{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deterministicDie{
				current: tt.fields.current,
				total:   tt.fields.total,
			}
			got := d.roll3()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, d)
		})
	}
}

func Test_player_move(t *testing.T) {
	type fields struct {
		current int
		score   int
	}
	type args struct {
		steps int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *player
	}{
		{"case 10", fields{1, 0}, args{10}, &player{1, 1}},
		{"case 15", fields{8, 0}, args{15}, &player{3, 3}},
		{"case 24", fields{10, 10}, args{24}, &player{4, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				current: tt.fields.current,
				score:   tt.fields.score,
			}
			p.move(tt.args.steps)
			assert.Equal(t, tt.want, p)
		})
	}
}

func Test_playGame(t *testing.T) {
	type args struct {
		p1 int
		p2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 4,8", args{4, 8}, 739785},
		{"Part1", args{10, 9}, 918081},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := playGame(tt.args.p1, tt.args.p2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_rollsDirac(t *testing.T) {
	tests := []struct {
		name string
		want map[int]int
	}{
		{"run", map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rollsDirac()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_playerD_move(t *testing.T) {
	type fields struct {
		players map[player]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[player]int
	}{
		{"test", fields{map[player]int{player{4, 0}: 1}}, map[player]int{{1, 1}: 6, {2, 2}: 3, {3, 3}: 1, {7, 7}: 1, {8, 8}: 3, {9, 9}: 6, {10, 10}: 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &playerD{
				players: tt.fields.players,
			}
			p.move()
			assert.Equal(t, tt.want, p.players)
			p.move()
			assert.Equal(t, tt.want, p.players)

		})
	}
}

func Test_playDirac(t *testing.T) {
	type args struct {
		p1 int
		p2 int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test 4,8", args{4, 8}, 444356092776315, 341960390180808},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := playDirac(tt.args.p1, tt.args.p2)
			if got != tt.want {
				t.Errorf("playDirac() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("playDirac() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
