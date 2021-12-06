package assignment

import (
	"strings"
)

type Day06 struct{}

type d06Fish struct {
	timer int
}

func (d *Day06) readFish(input string) []*d06Fish {
	in := strings.Split(strings.TrimSpace(input), ",")
	nums := MakeInts(in)

	fish := make([]*d06Fish, 0, 20_000_000)
	for _, n := range nums {
		fish = append(fish, &d06Fish{
			timer: n,
		})
	}
	return fish
}

func (d *Day06) PassDay(fish []*d06Fish) []*d06Fish {
	var nwFish []*d06Fish
	for i := range fish {
		f := fish[i]

		f.timer--
		if f.timer >= 0 {
			continue
		}

		f.timer = 6

		nwFish = append(nwFish, &d06Fish{
			timer: 8,
		})
	}

	return nwFish
}

func (d *Day06) PassDays(fish []*d06Fish, days int) int64 {
	for i := 0; i < days; i++ {
		fish = append(fish, d.PassDay(fish)...)
	}
	return int64(len(fish))
}

func (d *Day06) SolveI(input string) int64 {
	f := d.readFish(input)

	return d.PassDays(f, 80)
}

type d06FishTimer struct {
	days [9]int64
}

func (f *d06FishTimer) PassDay() {
	newFish := f.days[0]
	for i := 1; i <= 7; i++ {
		f.days[i-1], f.days[i] = f.days[i], f.days[i+1]
	}
	f.days[6] += newFish
	f.days[8] = newFish
}

func (f *d06FishTimer) PassDays(days int) int64 {
	for i := 0; i < days; i++ {
		f.PassDay()
	}

	var sum int64
	for _, fish := range f.days {
		sum += fish
	}
	return sum
}

func (d *Day06) SolveII(input string) int64 {
	in := d.readFish(input)

	days := [9]int64{}
	for _, f := range in {
		days[f.timer]++
	}

	f := d06FishTimer{
		days: days,
	}

	return f.PassDays(256)
}
