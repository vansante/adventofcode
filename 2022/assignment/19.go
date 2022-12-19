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
	return d19Resources{
		ore:      r.ore - other.ore,
		clay:     r.clay - other.clay,
		obsidian: r.obsidian - other.obsidian,
	}
}

func (r d19Resources) add(bots d19Bots) d19Resources {
	return d19Resources{
		ore:      r.ore + bots.oreBots,
		clay:     r.clay + bots.clayBots,
		obsidian: r.obsidian + bots.obsidianBots,
	}
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

type d19Bots struct {
	oreBots      uint16
	clayBots     uint16
	obsidianBots uint16
}

func (b d19Bots) clone() d19Bots {
	return d19Bots{
		oreBots:      b.oreBots,
		clayBots:     b.clayBots,
		obsidianBots: b.obsidianBots,
	}
}

type d19State struct {
	minute    uint8
	resources d19Resources
	bots      d19Bots
}

type d19BotCollection struct {
	states    map[d19State]uint16
	blueprint d19Blueprint
}

func (b *d19BotCollection) collect(minutes uint8, bots d19Bots, resources d19Resources) uint16 {
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

	var maxGeodes uint16

	// Build bots
	if resources.has(b.blueprint.geodeBot) {
		resources := resources.subtract(b.blueprint.geodeBot).add(bots)
		geodesMined := uint16(minutes - 1)
		maxGeodes = util.Max(
			maxGeodes,
			geodesMined+b.collect(minutes-1, bots.clone(), resources),
		)

		// Always build one when we can
		b.states[state] = maxGeodes
		return maxGeodes
	}

	if bots.obsidianBots < b.blueprint.geodeBot.obsidian && resources.has(b.blueprint.obsidianBot) {
		resources := resources.subtract(b.blueprint.obsidianBot).add(bots)
		bots := bots.clone()
		bots.obsidianBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources),
		)
	}

	if bots.clayBots < b.blueprint.obsidianBot.clay && resources.has(b.blueprint.clayBot) {
		resources := resources.subtract(b.blueprint.clayBot).add(bots)
		bots := bots.clone()
		bots.clayBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources),
		)
	}

	maxOre := util.Max(b.blueprint.geodeBot.ore, b.blueprint.obsidianBot.ore, b.blueprint.clayBot.ore, b.blueprint.oreBot.ore)
	if bots.oreBots < maxOre && resources.has(b.blueprint.oreBot) {
		resources := resources.subtract(b.blueprint.oreBot).add(bots)
		bots := bots.clone()
		bots.oreBots++
		maxGeodes = util.Max(
			maxGeodes,
			b.collect(minutes-1, bots, resources),
		)
	}

	// Continue without building bot:
	maxGeodes = util.Max(
		maxGeodes,
		b.collect(minutes-1, bots, resources.add(bots)),
	)

	b.states[state] = maxGeodes
	return maxGeodes
}

func (d *Day19) makeBotCollections(bp d19Blueprint) d19BotCollection {
	return d19BotCollection{
		states:    make(map[d19State]uint16, 1_000_000),
		blueprint: bp,
	}
}

func (d *Day19) SolveI(input string) any {
	bps := d.getBlueprints(input)

	var sum int
	for _, bp := range bps {
		bc := d.makeBotCollections(bp)
		max := bc.collect(24, d19Bots{oreBots: 1}, d19Resources{})
		sum += int(bp.id) * int(max)
	}

	return sum
}

func (d *Day19) SolveII(input string) any {
	bps := d.getBlueprints(input)

	total := int64(1)
	for _, bp := range bps[:3] {
		bc := d.makeBotCollections(bp)
		max := bc.collect(32, d19Bots{oreBots: 1}, d19Resources{})
		total *= int64(max)
	}

	return total
}
