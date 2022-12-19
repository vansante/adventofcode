package assignment

import (
	"fmt"

	"github.com/vansante/adventofcode/2022/util"
)

type Day19 struct{}

type d19Resources struct {
	ore      uint16
	clay     uint16
	obsidian uint16
	geode    uint16
}

func (r d19Resources) has(other d19Resources) bool {
	if r.ore < other.ore {
		return false
	}
	if r.clay < other.clay {
		return false
	}
	if r.obsidian < other.obsidian {
		return false
	}
	return true
}

func (r d19Resources) subtract(other d19Resources) d19Resources {
	r.ore -= other.ore
	r.clay -= other.clay
	r.obsidian -= other.obsidian
	r.geode -= other.geode
	return r
}

type d19Blueprint struct {
	id          uint8
	oreBot      d19Resources
	clayBot     d19Resources
	obsidianBot d19Resources
	geodeBot    d19Resources
}

func (d *Day19) getBlueprints(input string) []d19Blueprint {
	lines := util.SplitLines(input)

	blueprints := make([]d19Blueprint, len(lines))
	for i, line := range lines {
		n, err := fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. "+
				"Each obsidian robot costs %d ore and %d clay. "+
				"Each geode robot costs %d ore and %d obsidian.",
			&blueprints[i].id,
			&blueprints[i].oreBot.ore,
			&blueprints[i].clayBot.ore,
			&blueprints[i].obsidianBot.ore,
			&blueprints[i].obsidianBot.clay,
			&blueprints[i].geodeBot.ore,
			&blueprints[i].geodeBot.obsidian,
		)
		util.CheckErr(err)
		if n != 7 {
			panic("invalid scan")
		}
	}
	return blueprints
}

type d19BotCounts struct {
	oreBots      uint16
	clayBots     uint16
	obsidianBots uint16
	geodeBots    uint16
}

func (b d19BotCounts) clone() d19BotCounts {
	return d19BotCounts{
		oreBots:      b.oreBots,
		clayBots:     b.clayBots,
		obsidianBots: b.obsidianBots,
		geodeBots:    b.geodeBots,
	}
}

type d19State struct {
	minute    uint8
	resources d19Resources
	bots      d19BotCounts
}

type d19BotCollection struct {
	states    map[d19State]uint16
	blueprint d19Blueprint
}

func (b *d19BotCollection) collect(minutes uint8, bots d19BotCounts, resources d19Resources) uint16 {
	if minutes <= 0 {
		return 0
	}

	state := d19State{
		minute:    minutes,
		resources: resources,
		bots:      bots,
	}
	if result, ok := b.states[state]; ok {
		return result
	}

	// Collect resources
	resources.ore += bots.oreBots
	resources.clay += bots.clayBots
	resources.obsidian += bots.obsidianBots
	resources.geode += bots.geodeBots

	var maxGeodes uint16

	// Build bots
	if resources.has(b.blueprint.geodeBot) {
		bots := bots.clone()
		bots.geodeBots++
		geodesMined := uint16(minutes - 1)
		maxGeodes = util.Max(
			maxGeodes,
			geodesMined+b.collect(minutes-1, bots, resources.subtract(b.blueprint.geodeBot)),
		)

		// Always build one when we can
		b.states[state] = maxGeodes
		return maxGeodes
	}

	if bots.obsidianBots < b.blueprint.geodeBot.obsidian && resources.has(b.blueprint.obsidianBot) {
		bots := bots.clone()
		bots.obsidianBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources.subtract(b.blueprint.obsidianBot)),
		)
	}

	if bots.clayBots < b.blueprint.obsidianBot.clay && resources.has(b.blueprint.clayBot) {
		bots := bots.clone()
		bots.clayBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources.subtract(b.blueprint.clayBot)),
		)
	}

	maxOre := util.Max(b.blueprint.geodeBot.ore, b.blueprint.obsidianBot.ore, b.blueprint.clayBot.ore, b.blueprint.oreBot.ore)
	if bots.oreBots < maxOre && resources.has(b.blueprint.oreBot) {
		bots := bots.clone()
		bots.oreBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources.subtract(b.blueprint.oreBot)),
		)
	}

	// Continue without building bot:
	maxGeodes = util.Max(
		maxGeodes,
		b.collect(minutes-1, bots, resources),
	)

	b.states[state] = maxGeodes
	return maxGeodes
}

func (d *Day19) SolveI(input string) any {
	bps := d.getBlueprints(input)

	fmt.Println(bps)

	var sum int
	for _, bp := range bps {
		bc := d19BotCollection{
			states:    make(map[d19State]uint16, 1_000_000),
			blueprint: bp,
		}
		max := bc.collect(24, d19BotCounts{oreBots: 1}, d19Resources{})
		fmt.Println(bp.id, max)

		sum += int(bp.id) * int(max)
	}

	return sum
}

func (d *Day19) SolveII(input string) any {
	return "Not Implemented Yet"
}
