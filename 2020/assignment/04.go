package assignment

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Day04 struct{}

var d04ColorRegex = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
var d04PassportIDRegex = regexp.MustCompile(`^[0-9]{9}$`)

type d04Passport map[string]string

var d04Mandatory = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func (p d04Passport) isValidPtI() bool {
	for _, key := range d04Mandatory {
		if p[key] == "" {
			return false
		}
	}
	return true
}

func (p d04Passport) isValidPtII() bool {
	if !p.isValidPtI() {
		return false
	}

	birth, err := strconv.ParseInt(p["byr"], 10, 32)
	if err != nil || birth < 1920 || birth > 2002 {
		return false
	}

	issue, err := strconv.ParseInt(p["iyr"], 10, 32)
	if err != nil || issue < 2010 || issue > 2020 {
		return false
	}

	expiration, err := strconv.ParseInt(p["eyr"], 10, 32)
	if err != nil || expiration < 2020 || expiration > 2030 {
		return false
	}

	var height int
	var unit string
	n, err := fmt.Sscanf(p["hgt"], "%d%s", &height, &unit)
	if err != nil || n != 2 {
		return false
	}
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		log.Panicf("[%s] invalid height unit: %d%s", p["hgt"], height, unit)
	}

	if !d04ColorRegex.MatchString(p["hcl"]) {
		return false
	}

	switch p["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	return d04PassportIDRegex.MatchString(p["pid"])
}

func retrievePassports(content string) []d04Passport {
	split := strings.Split(content, "\n\n")

	const separator = ":"

	var passports []d04Passport
	for i := range split {
		data := strings.TrimSpace(strings.ReplaceAll(split[i], "\n", " "))
		parts := strings.Split(data, " ")
		p := d04Passport{}
		for i := range parts {
			idx := strings.Index(parts[i], separator)
			if idx < 0 {
				log.Panicf("[%s] [%s]: no separator found", data, parts[i])
			}
			p[parts[i][:idx]] = parts[i][idx+1:]
		}
		passports = append(passports, p)
	}
	return passports
}

func (d *Day04) SolveI(input string) int64 {
	passports := retrievePassports(input)

	valid := 0
	for i := range passports {
		if passports[i].isValidPtI() {
			valid++
		}
	}
	return int64(valid)
}

func (d *Day04) SolveII(input string) int64 {
	passports := retrievePassports(input)

	valid := 0
	for i := range passports {
		if passports[i].isValidPtII() {
			valid++
		}
	}
	return int64(valid)
}
