package main

import (
	b "aoc/utils"
	"github.com/Goldziher/go-utils/sliceutils"
	"github.com/alecthomas/participle/v2"
	"log"
	"slices"
)

func Parse(contents string) *SeedingPlan {

	// parser, parse_error := participle.Build[SEED_PLAN]()
	var basicParser = participle.MustBuild[SeedingPlan](
		participle.Lexer(basicLexer),
		participle.UseLookahead(2),
	)

	ast, err := basicParser.ParseString("", contents)
	if err != nil {
		log.Println("Parse String error", err)
		return nil
	}
	return ast
}

func (s *SeedingPlan) GetSeedsList(c *Context) {
	for _, seed := range s.Commands[0].GivenSeedsCommand.SeedsList {
		c.seeds = append(c.seeds, *seed.SeedIds)
	}
}

func (s *SeedingPlan) GetSeedsListFromRange(c *Context) {
	seedsList := s.Commands[0].GivenSeedsCommand.SeedsList
	for i := 0; i < len(seedsList); i += 2 {
		seedRange := AddressRange{start: *seedsList[i].SeedIds, end: *seedsList[i].SeedIds + *seedsList[i+1].SeedIds - 1}
		c.seedsRange = append(c.seedsRange, seedRange)
	}
}

func (section *MappingSection) GetMappedIds(sourceId int) int {
	for _, mappingRecord := range section.MappingLines {
		var (
			rangeLen    = *mappingRecord.RangeLength
			sourceStart = *mappingRecord.SourceStart
			destStart   = *mappingRecord.DestStart
			sourceEnd   = sourceStart + rangeLen
		)

		if sourceId >= sourceStart && sourceId <= sourceEnd {
			return sourceId - sourceStart + destStart
		}
	}
	return sourceId
}
func (section *MappingSection) SplitInputRange(inputRange AddressRange, outputs *[]AddressRange) {
	for mappingRecordIndex, mappingRecord := range section.MappingLines {
		var (
			sourceRange = AddressRange{start: *mappingRecord.SourceStart, end: *mappingRecord.SourceStart + *mappingRecord.RangeLength - 1}
		)

		splitRanges := splitRange(inputRange, sourceRange)

		log.Printf(".......For range %+v recorded split as %+v from rule %d\n", inputRange, splitRanges, mappingRecordIndex)

		*outputs = append(*outputs, splitRanges...)
	}
}
func (section *MappingSection) TranslateRanges(splitRange []AddressRange) []AddressRange {
	var returnValue = make([]AddressRange, 0)
	for _, inputRange := range splitRange {
		// log.Println("Translating", inputRange)
		var foundMappingForInput = false
		for _, mappingRecord := range section.MappingLines {

			sourceMapping := AddressRange{start: *mappingRecord.SourceStart, end: *mappingRecord.RangeLength + *mappingRecord.SourceStart - 1}
			destinationMapping := AddressRange{start: *mappingRecord.DestStart, end: *mappingRecord.RangeLength + *mappingRecord.DestStart - 1}

			if inputRange.start >= sourceMapping.start && inputRange.end <= sourceMapping.end {
				// log.Println("Found matching mapping, source:", sourceMapping, "Destination:", destinationMapping)
				diffStart := inputRange.start - sourceMapping.start
				diffEnd := sourceMapping.end - inputRange.end

				returnValue = append(returnValue, AddressRange{start: destinationMapping.start + diffStart, end: destinationMapping.end - diffEnd})
				foundMappingForInput = true
				break
			}

		}
		if foundMappingForInput {
			continue
		} else {
			returnValue = append(returnValue, inputRange)
		}
	}
	return returnValue
}

/*
*  Returns splits iff there is a INTERESECTING range. Else nil
 */
func splitRange(source AddressRange, destination AddressRange) []AddressRange {
	var rangeStops = []int{source.start, source.end, destination.start, destination.end}
	slices.Sort(rangeStops)

	rangeStops = b.Filter(rangeStops, (func(aStop int) bool {
		if aStop < source.start {
			return false
		}
		if aStop > source.end {
			return false
		}
		return true
	}))

	rangeStops = sliceutils.Unique(rangeStops)

	var doesRangeIntersect = source.intersects(destination)
	var splitRanges = make([]AddressRange, 0)

	// log.Println(rangeStops, doesRangeIntersect)
	if !doesRangeIntersect {
		return splitRanges
	}

	var start int = -1
	for stopIndex, stop := range rangeStops {
		if start == -1 {
			start = stop
			continue
		}
		correct_stop := stop - 1
		if stopIndex == len(rangeStops)-1 {
			correct_stop = stop
		}
		splitRanges = append(splitRanges, AddressRange{start: start, end: correct_stop})
		start = stop
	}
	return splitRanges
}

func (s *AddressRange) intersects(another AddressRange) bool {
	if s.start < another.start &&
		s.end < another.start {
		return false
	}
	if s.start > another.end &&
		s.end > another.end {
		return false
	}

	return true
}

func (s *SeedingPlan) Evaluate(c *Context) {
	s.GetSeedsList(c)

	for _, mappingSection := range b.Map(s.Commands[1:], (func(c *Command) *MappingSection { return c.MappingSection })) {
		for seedIndex, seed := range c.seeds {
			c.seeds[seedIndex] = mappingSection.GetMappedIds(seed)
		}
	}

	log.Println(slices.Min(c.seeds))
}

func (s *SeedingPlan) EvaluatePart2(c *Context) {
	s.GetSeedsListFromRange(c)

	// log.Println(c.seedsRange)

	for _, mappingSection := range b.Map(s.Commands[1:2], (func(c *Command) *MappingSection { return c.MappingSection })) {
		log.Println("ENTERING STAGE", mappingSection.MappingHeader, c.seedsRange)
		var allTranslatedRanges = make([]AddressRange, 0)
		for _, seed := range c.seedsRange {
			var splitRanges = make([]AddressRange, 0)
			mappingSection.SplitInputRange(seed, &splitRanges)

			if len(splitRanges) == 0 {
				splitRanges = append(splitRanges, seed)
			}
			log.Printf("Seed %+v Split into to %+v", seed, splitRanges)
			var translatedRanges = mappingSection.TranslateRanges(splitRanges)
			allTranslatedRanges = append(allTranslatedRanges, translatedRanges...)
			log.Printf("Split %+v translated to %+v", splitRanges, translatedRanges)
		}
		c.seedsRange = allTranslatedRanges
		log.Println("EXITING STAGE", mappingSection.MappingHeader, c.seedsRange)
	}

	var lowestLocation = slices.Min(b.Map(c.seedsRange, func(a AddressRange) int { return a.start }))

	log.Println("Solution is : ", lowestLocation)
}

type Context struct {
	seeds          []int
	seedsRange     []AddressRange
	mappingRecords []MappingRecord
}

type AddressRange struct {
	start int
	end   int
}
