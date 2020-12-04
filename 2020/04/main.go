package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var colorRegex = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

var passportIDRegex = regexp.MustCompile(`^[0-9]{9}$`)

type passport map[string]string

var mandatory = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func (p passport) isValidPtI() bool {
	for _, key := range mandatory {
		if p[key] == "" {
			return false
		}
	}
	return true
}

func (p passport) isValidPtII() bool {
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

	if !colorRegex.MatchString(p["hcl"]) {
		return false
	}

	switch p["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	if !passportIDRegex.MatchString(p["pid"]) {
		return false
	}

	return true
}

func retrievePassports(file string) []passport {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n\n")

	const separator = ":"

	var passports []passport
	for i := range split {
		data := strings.TrimSpace(strings.ReplaceAll(split[i], "\n", " "))
		parts := strings.Split(data, " ")
		p := passport{}
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

func main() {
	wd, _ := os.Getwd()
	passports := retrievePassports(filepath.Join(wd, "04/input.txt"))

	partOneValid := 0
	partTwoValid := 0
	for i := range passports {
		if passports[i].isValidPtI() {
			partOneValid++
		}
		if passports[i].isValidPtII() {
			partTwoValid++
		}
	}
	fmt.Printf("Part I: There are %d valid passwords\n", partOneValid)
	fmt.Printf("Part II: There are %d valid passwords\n", partTwoValid)
}
