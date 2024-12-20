package main

import (
	"fmt"
	"strings"
)

// Function to count "XMAS" occurrences in a single row or column
func countInLine(line string) int {
	count := 0
	target := "XMAS"
	revTarget := reverse(target)

	for i := 0; i <= len(line)-len(target); i++ {
		if line[i:i+len(target)] == target || line[i:i+len(target)] == revTarget {
			count++
		}
	}
	return count
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to count "XMAS" in all diagonal directions
func countInDiagonals(grid []string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Traverse diagonals from top-left to bottom-right
	for d := 0; d < rows+cols-1; d++ {
		var diagonal string
		for i := 0; i < rows; i++ {
			j := d - i
			if j >= 0 && j < cols {
				diagonal += string(grid[i][j])
			}
		}
		count += countInLine(diagonal)
	}

	// Traverse diagonals from top-right to bottom-left
	for d := 0; d < rows+cols-1; d++ {
		var diagonal string
		for i := 0; i < rows; i++ {
			j := cols - 1 - (d - i)
			if j >= 0 && j < cols {
				diagonal += string(grid[i][j])
			}
		}
		count += countInLine(diagonal)
	}

	return count
}

func main() {
	// Puzzle input
	puzzle := `
MSAMXXMASXSXMMSASXMASMXMASXSXSSMSSSMMMSSSSMSAMXASAMXXAXMSMMMAXMAMSAMXAMXMXXMXMASMXSAXMXMAMXSXSXXMSSXMSSMMSSMMMMXSXAXMXMAMMASMSXSXSSXMSSMSSMM
MSASASAAXSMAMMSAMXSXXMAMXSASAAAASMSASAASAAAMMMSXXMMMSMMXAAASMMMMSAMXSAXASMMSSMMXAXXMSSXXMSMMASMXMASXAAXAAAAAAMAASXSAMAXSMSMMASAXAXAAAAAMAAAS
ASAMASMMMASXMAMAMXXAXSMSAMAMMMMMMAMMMMXXMMMMSASXXMAXAASXSSMSASAMXMSAMXSASAMAASAMXMAAAMXMASAMAMXMMASMMMSMMSSMMMMMSAAASMSMAAMSAMAMXMXMMXSMXXMM
MMAMAMAAAMAMMASMMXMMMXAAMMXMMXXXMXSXASXSMSXXMASAMSSSMSMAAMASXMASAAMASXSAMXMSSMMXSMMMXSAXAXAMSSMAMASXMASXMAMMSMMASAMXMMAMMMMMXSXMSMXMASMXSASA
XSAMAXSMSXAMSASASXMAAXSSMXXXAXMXSAAXXSAAASASMAMAXAXAMXMSMMMMMSAMMMSXMMMAMXAXAAAAMXMSAXMSSMSMASASMAMXMASXMAXAAAMMSAXXSSXMXSXMXXAMXMMMSMAASMMS
MSXMMXMXXMMMMASAMMSMXXAAXXSASMMAMAMMXMMMMMAMMASAMMSAMSXMASXXXSAXSXSMSAMASXMSSMMSXAXMASMXMAXMMXSAMSSSMXSASAMXSSMXSASAMXMAXMASMSSMASAXAMMMSMAM
XMAMSAMMSAMXMMMMMAASMMSAMXXMMAMASMXSAMASXMXMXASASXAAASASAMMMASXMSXMASXSASAMAMXAMMSMMXMXAMAMAMSMMSAMAAMSAMXSAAXMASAMXMXMASMXMASASASXSMAXAXMAS
SMAMXASAMAMXMMXAMSSSXAMASXXXXXSASAAXASASXMAXMXXXMMXMXSAMXSXAXSMMMAMAMAMMSXMMSMASAAAMXSSMMXSXMAAXMASMMMMAMAMMMMMXSXXAAASAMXAXXSAMXSAMXSMMXSAS
XSXSSMMXSAMXMMMMMMAMMXSAMMMMSXMAMMMSAMASASMSSSXAXMAMXMMMMMXSAMAAXAMASXSAMXXSAMXMMSMMAXASAAMASXSMSMMMAAXAMSMSAMSASXSMSMSAMSMMMMMMSMMMAMASXMAS
MAMXAAAAXMMAMAAMAMSMSXMASAAAAAMXMAASXMMMMMAAAXSXASAXMXMAXSAMXSSMSSSXSXMXSAXXMSAMXXXMXSAMMMSAMAAAAMASXMXMMMAMAMSAMAXAMASMMAMASAMXAAAMAMAMAMAM
SAAMSMMSSXXAXSMSAXXASXMXSXMMSAMXMMMSXSXSAMMMSMMSMMAXSASASMAMAXAAAAAAMMSXMXSXASMSAMAMAMAMXAMAXXXMAMASAMASAMAMMAMAMSMMMAMASXSASASAMSMSMSSSSMAM
SXMAAXXAAMSXXAASMMMMMAXMSAMXXAMXXMXMXXAXAXXMAAMXXMSMSASXXMSMMSAMMSMAMAMAMAMMXMXMAXAMASAMMMXSMXMASMMSASASASXSMMMMAAAXMASMSAMXXAMMXAXAMXMAMSAS
SAASAXSMXMAMMMMMMAXASXMASMMMSAMMSMAMSMSMSMAMSXXASMMAMMMMSMMAMSXXAXXXMMSAMASASXSXMMXMAMASAMAAMMMAAAASXMASMMXXXXAASXXMXAAXMMMSMMSAMXSMSSMAMSXS
SAMAAMSXSMMSMAAXMMSMMMAMSXAAXSMSASASXAAAMSSMMMSXSAMXMAMASASAMXMMMSAMXAMASXXASMSAMSSMSSXMAMSSSXMXSMMSXXXSAMXAXMMMMASXSMMXXSAMXAAAMAAXAXMAXXAM
MAXXSXSAMMXAXMSSMMMXMMSMMMSSXXXSAXXXMXMXMAXXAXMASMMSSSMASMMAMAAMAXAMMSSMMXMMMAMAMAAAAXAMMMXAMMMMXXAMXMMSAMMMMAMSMAMAAAMMMMAMMXSMMMXMASMMSMAM
XMMMXMMAMSSMSXAXMAASMAXAXAAXMMMMSMASXXSAMAMSMSMAMXSXAAMAMMXMSSMSMSAMAAXAAXMASMSAMSMMMXMMAXMAMAAMMMXXAAAXXMAMMAXXMASMMMMAMMAMSMMAAXSAMXXAAMXM
AMAXAMSMMMAASMMSXMXMMAMMMSXMAAAAMMMMAASASAAAAAMXXXMMMMMSMSAMXMAAXSAMMSSMMSSMMAXXMAMASAXMMMSAXMSSMAAXSXMSASXSMMXSAMXMXSXSMMAMXASXMSXXMAMXXSXM
XXAMXMAAASMMMAMAMSMSMSSXAXMASXMSMMXMXMXASASMSMSSXSMSXSAXXXASAMXMMMAMAMXAXAMXMASXSMSASXSAAXSMSMXAMMSXMAXXMAMAAAMAMXASXSAXAMMSMXMAXXMSMSSSSMAX
MXSXMMSSMSAMXSMMMAASAXMMXSAMAAMAXSAMXXMAMAXAXAAXAXAAAMXSMSMMASAMXSXMASXSMSMXMMSAAAMASASMMMMAAXMXMMMAMXMASMSSMSSSMSXSAMAMMXMAMXXMMSMSAAAMXSMM
SAMAAAAAXSXMAXAXSMSMXMASMSXSXXMAXSASMSXMMSMAMMMMMMMMSMMXMAMXAMASAAXSSSMMAMAAMXMMMMMAMXMAAASXMMXMMXMXMSMXAAAAAXAMXXAMAMAMXXSASMMXAMAMAMSAAAMA
MASXMMXSMMMMMSXMASXMMAAXASAMXSXSASAMMXXAAAASXMMXMXMAXAAASASMASXMXSXXAXAMAMSSSSSXMAMAMAXSMXXAASMASAXMXAXAXXMXSMSMXMSMASXSAXMASAAASMXMXXXMSSMS
MAMMXSAMASXAMXASMMMASMSMMMXMMMAMMMAMAMSMSSXXAXSAMMSMMSSMSASXXSMSAMMMMMSSXXAMMAXMMMXAXXXXMMSMMAASMXXAXMSMSXSAMMXSXXXSMSAMMSMAMMMSMAMXXMSXMAXA
MMSMAMXMXMMMMMAMAXMAXMAMXSXSAMMMXSAMAMXAMXASAMSASAMXAXAMXXMASMAMMMMAXXXMMMASMSMMXMASMSXSAAXASMSMSMMXAXAXMXMASXAAMXMMMMXMSXMXMMXXXXSXASMASMMM
MSAMMSMSSSXSAAAXXMMAMSAMASASAMMSXMXSAMMAMMMMMXSAMXAMXSAMXMXXXMSMXXMMSMSMXSSMAXAXAMAMAXAXMMSXMXAMAMASASMXSAXMMMMSMASAMXAXXAMASMMSMMSXMXSXMXAX
MXASAAAXAAASXXMXSAMXAMASXMMMXSSMMSASAXSSMSXASMMXMMXMMSAMMXSMMMXMAXSAAAAAAMAMMMMSASXMSMSMMAXXXMXSAXMAMXMASMMMXAMAXXSXSXMASAMASAAAAAMAMMMMMSSM
SSMMXXSMMMMMAMSASXSMMSASXMAMMMXAASMMSMAAXASAXASXSMSXASASAMXAMSAXAMMSMSMSXXAAMXXMAXAAMMMAMXSSSSMSMXMAMAMMMXAXSXSASMSAMXXAMAXASMMMSMXAMAAAXAAS
XAMXMMMAXXXMXMMASAXAAMXMASMSASMMMSXAXMSMMMAXSAMXSAXMXSAMMMSXMXMXASAXXAAAXSSSMXXXMMMMMASAMXAAMSASASXSSSSMSMMXAAMASXSAXAMMSSMMSMSXMMSSSSSSMMSM
SAMXSASXSXSMMXMAMAMMMSASXMASMSAAAMMSSMXAMXAAMAMAMXMMMXXMAMXSAAMSXMASXMSMMMMAMXMMMASXSMSASXMXMMXMAMAAXAAAAASMXMMXMASXMMSMAMAXMAMAXMAMXMAMXXAM
MMSMMMSASAMXAXMASXSXSAMXAMMMXSXMMXAXAAXMAXMXMAMASAMMASASXXAAXMMASMAMAXXXAASAMASASASAXXXXMAXSMSMMMMMMMSMMMAMAAXMSAMXMXAAMMSSMMAXMMMASMMSMSMMS
MAAXAXMAMXMASXMXMMXMAXXSAMXSMSMMMMXXMMMSSSMXXASAXXXAMXAXAMSSMSXASMXSXMXASMXASMSASAMXMMSMSXMMAAXAXXXMAXXXMASMMSAMMMAXMASMMMASXAASASASMAMASAMX
MSSMMSXXAAMXXXMAAMMSSXMMXXMSAXMAASMSMXXAAAASMMMMSMSSSMXMMMMAAXMASXMAMXMMMXMAMXMMMMXMSXAASXSMSMSMSSMMSMMXMAXMASASXSMSXAXAASAMXXSAAMASMXMASAMM
XAMAXMMXSXSASMMSSMAAMXMAASAMMMMSXMAAXAMMSMMMAXMXAAXAAXMAMASMMMMXMASXSXXSAXXAMXSXASAAMMMMMASAMASAAAMMAAMSSMMAMSXMAMSSXMMSMMASMMMMAMXMXAMAMASX
MAXSASXAMAMMSSMXAMMSSXMAMMAMXXAXAMSMSXXAMAXSXMASMSMSMMMXSASASAXASXMAMAMMASAAMAXSAMMSMMAAXAMAMAMMSSMSMSMAAAXSXSAMMMASASAXXSAMXSASMSAMSMSMSAMX
MMMMAXMMMAMASAXSAMAMMMMXSSSMMMMSMMAAMAMASMMMMMXXAMAMAMSAMMSASMSASAMAMAMXAMMSMAMSMSMXASMSMASXMAMMAMAXAAXSSMMMASMMXMASXMASAMSSMSASASAMSXAMMASX
SASMSMAAMXMXSAMSAMXXAAAMMAMMMAMAASMMMAAXMAMAAXMMXMASAMMASAMMMAXAMAMXSMSMSSXMMXXMAAASAMAAMAXXSXXMAMSMSMXMMMAMXMXMXMASAMAMAMMXMMXMMSAMMMSXSMMM
MASAXSSMXSMAMMMMXAXSSXXAMAMXSSSSMMAMSMSSSXSSSSMSASXSASMXMMSAMSMMSASASXAXAMAXMMMMSMXMSMSMSMSASAXSXMMMMMMSXSASXXAMSMMSAMXMXMMMSXSAXMAMAAXMMSAS
MSMMMMMSAAMXXAAMAMXXMAXSSMSMMMAMXMAMAMAAXAMXAXAMMSXSXMASXXMAAMAMSAMASXMMMMMMMAXAXXSAMXXAMAMAMXMMAMXAASAAASASAMASXAXMXXSXSMSAMAMXMSSMMMSAMMXM
SAAXXAAMSSMMSSSSSMXMMMMMAMXAMMSMASXSSSSMMSMMSMSMXMAXASAMMASAMMAMMAMAMAMAXAAMXXMASAMXMXMAMSMAMMMSAMXSSSMMMMXMMSMXXMMASMSAMAAXMSMAMAXAMAMAMAAM
ASMMSMXMAMXAAAAAMXAMAXASXMMSXAAMXSAAAAMXAXASXAAXMASXXMASMMMSMSMXSAMXXAMXMSMSASMMSAMXSXSXMXSAXAMSXMXMAMAMSMSSXAMMMSMMMAMSMAMSMAXAMMSMMAMXMSAS
MMMXXXXAASMSMMMMMMXSMSASAAAMMSMSAMMMMMMMSXMMMSMSMAMXMXXMAAAMMAAXMMXXXASAMMASMXAASXMXSAMAMAMMMSMMMSMMAMAAAAXASASAASAMMMAAXAMAMMSMMXAXMSSMXMAA
XMMXMASMXSAXMXMXSAMAAMXSXMXSAXAMMXXASXSAXAMXAAMAMXMAMMSMSMMXSMMMAMAXMMMMXSASXMMMMAMAMAMAMSSMAAAAAXXMASMSMMMMSXMAMSAMAMSSSXSXXXAXASMSMMAMAMXM
AMXMSAAXAMXMSXSASAXXXXXXMSMMMMAMSSMXSAMASAMSSMSXSASASAXAASMMSAXSSMSMXAAXAMXMAXSXSAMMSMSMSXAMSSSMSSXAMAMXAMSXMASXMSAMXMXAXASAMSAMXSXAMSAMAMMX
SAAMMSSMMSASXSMASMMSSMMSAMXAMSSMMAMAMXMASAMAXXAMSASASMMSMXSAMXMMAAAASXMMMSAMXMAASXSXAAAMMMAMXMMAXMAXSAMXSAXMAMXMASXMXXMMMMMAMMMMMMMAXMXMAMXM
MXMMAMXAXSXSXAMAMAXAAAAASMMAXAXXXAMXSXSASXMASXSMMAMAMXAAAXMMMAAXMMMMMAAMASXSMAMXMMMXMSMSAMMSMSMMMSAXSASAMMMSAMAAMMMMXMMMSMMMMMAAASXMMSMMAMAA
SAMMSMXAMXMMXAMASMMSXMXSMMMSMASMSMSASMMASMSAMXMAMSMMMMASAMXSMXSXSMXXXSMMASXMAXMXMAAAXMAMXSAAMMASAMXMSAMXSXAXAMSSXAASAAAASAAAAXMXMMAAAAMMSSXS
SAMMAASXMSAXSXXXAMAXXMMXASAAMAMMAXMASAMAMMXXMASAMAAAXMAMAXAXSAXMASXMAXAMXMAXXASXSMSSSMAMAMXXSSMMXSXMMMMMSMMSSXMXMSSXMMMXSAMSASXSSSSMMXXMXAMX
SAMMMMMAASMMXASMXMASAMXXAMSMMSXSXSMMMXMXSSSMSASMXSSMAMASMMSSMXXAMXMAMXAXSSSMMMAAXMAXXMAMMXSXAAXAXSMMAAAMXAAAMXAMMMAMMASAMXMMXXAAAMXXSMSXMASX
SMMXSASMMMAMMMMXAMAXAMXMSMMXAXAXAMXSAXXSMAMAMXMAAXAMXSXSXAMAXSMSXXXASMSMAAMXASMAMMMSSSXMSASMSMMMSMASXSSSSMMSSMMMAXMASASASMXXMMMMMMMAAAAXSAMM
SAMXSAXAASAMAMASXMASXSSXAASMMMXMASAMXSMMMAMSMAMMMXAMASASMSSMMSAMASMAMAAMMMMSASAAMAMMAMAAMAMAAAMMMMMMAAAMAASXMAXSMSXAMASAMAAXXAAAAAXXMXMMMASA
SAMXMAMSMSSXXMMAMXXMAAASMSMASAMSMMASMMXMAMMXMXSASXMMMMMXAXMAAMAMAMMSMSMMMAASXSXMMAXSAMXMMMMSSMMMAMXMXXSXSMMASMMMASXMMAMAMSMSMMMXXSSXMXSASAMX
SSMXMAMAAXAXSMSMSMMMMMMXMXMAMMMAAXXMAXMAASMSXMXXMASMSMXMSMXMMSAMXSAAAAXMXMXMAMXSSSXSASMSASAAAAMMAXXXMXMAMXMXMAMMAMAXMAXXMAMSASMSMAMXMAXMXMAS
SAXMMMSXSMMMXAAAAAASXSASMAMMMASXSMSSSMASMSAMASXMSMMAAMAMAMXXMXASMMMSSMXSAMXMSMAXAMASAMAAAMMSSMMSSSMASAMXMSSMSMMMASMMSASAMXXSAMAAMAMMMASMXSSM
SMMSAMXMMAXAMSMSMSMSASAMMXSASMSAXAXXXAXXAMMSAMMAAMMSMSXXAXXMSSMMAMMMAXAMAXMMMAMMMMMMAMMMMMAMXXMAMAAMMASMAAAAAAASXMXMXMAXMMASMMXMSMXSMASMAMAM
SXASMSAMSSMSAMAMAXAMXMAMAXSXSAMMMXMXMSMMMMAMASXMASAXXAASMSMAAXMASAASXMMSAMSASAMAAAXSAMASXMASMMAMMMMXMAMMMMMSMSMSAMXMAMAAXSASMSMXMMAXMXMMAXMM
XMAMXMAMAXAXMMAMXSAXXAMMMXMASASMSAAXAAXAXMXSMMAMMMXSXMAMAAMMMMXAASAMXAMMMMSASMMXMXXSASAMASXSAMMMASXMSSXXAXXXMMASAMAMAMXMMMAXAASAAMMSSMMSSSXS
SSMMSSSMMSXMASXMASAMSSMAXASASAMAMXXXSSSMXAMXXMMSAAAXAXMMSAMXSMMMMMSSMSAXAAMMMAASXMMSMMXSXMASMMASMSAMAMASMSMAMMMMSSMSMSAXMMSMSMSMSAAMAAXAASMS
AAAXAAXMAXAAXXAMXMAMAAXXMAMXMMMAMMMMAAAMMSASAMAMMSMMMMXAMASAMXAMXAAXXAMMMMXAMXAAAAMXMAXXAMAMASAMXSAMAMXMAAMAMAMMASMAASAMXAMXMAXAMMXSSMMMSMXX
SXMMMMMMMSSSXSXXXMAMMSMXAAMAAAXAXSAAXMAMAAMSAMASXMXAMXMAMAMASMSSMMSSSMXAAMSMSMSMSXMAMMMSMMASXMMSMMXMAMAXXMSSSMSMAXSMXMMMMAMAMMMMMAXMASASAMMS
MASAMXAXXAAXXAMXSMMMAAAXMASMXMSMSAMMMXMMSSMSAMXXAAMMSAXXMSSMMAMXMAMXMXAXMXAAXXXAXMMXSAAAASASMSMAASMSSSSMMXMXAAAMSMXXMASXXAMSMSAMMSMMASXSASAX
SAMASXMMMMSMSXMASAAMSSSMSMXMXXAMAMSXSAXAAAAXXAMSMMMXMASXXAAAMSSSSSSSSXSSSSMSXMMSMMSAMMMSMMASAAXMXMAAXSAASMXSMMMAMASXMXAAMMMXASXSAAMMXSMSMMXS
MXSAMMSAAXMAMXMAXAMXXAXAXXAMXMASAXAAMAMMSMMMSXMAMASXMASXMMSXMMAAXAAAXAXAAXMMAMMAASXSMAXXAMXMMMXSAMMMMSSMMMASXMXASMMAMXMMAAAMXMMMMSSMXSAMAMAS
MMMMSASXSMMAMMMMSSXSSMMMMMSMSSMMMMXMMAXMAXSAXAMMSASAMSMMAXAAMMMMMMMMMXMMSMASAMSMSMAMMXMXAMMMMSAMMXXMXMAXAMASASMMSXXAMAXXSMMSMMXSAMXMAMSMAMAS
AAAMMASXXASAMMAMXMAXAXAXSAAAAMSAXMMMSXSMXMMMSSMMMMXAMAAXSMMSMMASAMAXAMXXXMASXMAXAMXMSSSSSMAAAMAXSAMSMMMSMMASMMAMSMMMMXMMXMXXAMASAMXMMMAXSXMS
SMSSMSXMMAMAXSXSAMSMXMMMAXMXMMSMMSAMAMXMAMXMAMAXAMSSSMSMAMAAMMASASXMXAXSXMASASAXMAMXSAXAAXXMXXXMMASAAAAAAMAXMSSMMAASMMXSAMXSSMASMSMMMMMSMAAX
XAMXMMAMSSXMAAASMXAAASXSSSSXXXXAASXSSSMSSSMMASMMSAAAMAMXAMSSSMASAMXXASAMAMAMMMMSXAASMMMSMMSMSMAMSMMXSMSSSMSSXAXMMMMMAAAMASMAXMASAAAAAASAMAMS
MAMAMSMMAXAMXMXMASMSMSAXMAMMMSMMMSAXXAMXAAAMAMAAMMSMSMSSXXXAMMAMASAXMMASAMMMSAASMSMMAXAXMXAAASXMAXMAMMMAMAMAMMSMMSMSMMMSAMMMMMAMMMSMSXSAXSSM
AXSAMAMMMMXMMMMSMMXAAMMMMAMAAXAASMMMSMMMSMMMASMXMAMASAAMXMMMMMMSAMMMXSAMASXAXMXSAXAMXMMSSSMMMSXMMMMMMXMAMSMXMXAMXAAAAMXMMMAMXMSXXMAXXMMMMXAX
SASMSXSAAMSMMAASXSSMSMSXSXSMSSSMSAAMMAMMMMXSXSXXMMSAMMMMAMMXSAMMASXMAMMSXMMXSXMMMMMASMMAAAASXMAMAASXMASXMAMMXMAMSMSMMSAMSSSXMAMAAASAMXASXSSM
XXSXMMSMSMSAXMMXAAAAAAMXXMSAMAAAMXMMSAMMAMASAMAXSMMASXMXMSAASAMSAMAMASAAXSMXAXMAXSMAAAMMSMMMAAXMAMSAXXSASXSAAMXMSXMAAMXXAXMAMAMSMMMXSSSSXAXX
MMSAMASMMAMAMSMMSMMSMSMAXAMAMSMMMAMMSXMMAXAMAMMMSASAMXSAMMMXSAMMSSXMXSMMMAMMMXXAAXMAXXXAXMASMMMAMSMXMAMASAMMXSAMXAMMMSMSMSSXMAXAAASAMMMMMMMM
AAMXMASAMMMSMMAAXAXXAMAMXSMMMMXAMAXAXSSSSMSXMMMXSAMASAMMSXMMMXMAMMMSAMXXSAMAAMASXMSSMMMSSSMMMASXMAAAMXMSMAMAMSASMSMSXAAAXAAXSSSSSMSAMXAMAAAX
SSMAMAMMXSAXAMMMSMMMSMMXAXAXXXMMMSMMMMMAXAMXSAMMMMSXMASAMXSSXMMMSAAMAXXASMSMMSAMAMAAAAXAAXAXSASXSMSMSXMAMXMAXSAMAAAMAMXMMMMMAMXXMASAMXMXSSSX
AAXSMXMXXMASXMSXSXMAXAMMMSSMMAASAXASXAAAMXMAXASMAAAAMAMMSAAMSMAAMMMSAMSXMXAXMMASXMMSMMMMSMMMMASXMXXAMMXAMMSXMXMMSMSMXMAAAAXXXXAMXASXMASAMAMX
SSMMAMSMSMAMAMSAMXSASAMSMAMXASMSASAMXMMXMAMMSMXMMXSAMASXMMMMASMSSMAMAXMASXMSXSAMXSXXXAXXAAXXXMSAMAMASMSXMAXSXAMXXAXMASXSMSSSMMSSMXMASASAMAMM
MMAMXMAASMSXSMMAMAMXSXMXMASMXXAXAMAMXSAXSAMXAMSMMMXMXMXAXSXSAMSAXMXMXAXXMAAAXMASAASMSSSMMSMSAMXAMMSAMMAXASMMSXMAMMMMXAMXMXAAXAAAXSMMSMMMMMSS
MMMMSSMSMAXSMASXMXSASAASMXSAMMMMMMXXXAAXMMSSMXAAAAMAAMXMMMAMASAMXASXMMSSSMMMSXAXXAMXAASAAMASMASAMXMASMSMMXAMSMSXSASXSMSASASMMMSSMSAMMAXMSAAA
MAAAAXXMMAMAXXMXAAMASXMAXMXXMAXMSAMXSMMMMAAMMSSSMSSSMXAAAXSMMMMXXMASAAAAMMMXXMMSMXAMMSSSMMAMMMAAXAXXMAXXXSAMXASXSASXAASMSAAXMAMAASMXMSMAXMAS
SSSMMSSMMSSMMSMXMXMAMASXSSSMSSSXMASAXXSXMSSMAAMXAAAAAXSXMSMMSAXSASASMMMMMASASXAAASMSMAMAMMXMASMMMSMAMMMMMSXMASMAMMMMMMMAMMMMMMSMMMXMSMAMAXSM
XAAXXAMAAXAXAAXAMSMMSAMMAAAAAAMXSAMXSAMXAAAMMSXXMMXMMMMASAMAMASXXMASXSXXMAXMXMMXMAAAMAMAXXMMXMAMAXMASAAAMXMAMAMXMASAMMMMMMXAAMXXXMAMAASXMSAS
MSMMMASMMSMMSSSXSAAXMASAMSMMMXMAMXMAMXAMMMSMMAXXASXAMSSMMAMXMMSAMXXMASMMMSSMAMSSMXMMSSSXSAXXASAMXXSASMSSSMXMMSMXSASMSAAMAXSXMSASMSSSSSMAMAMX
AAXMAMMXMAXAXXMXAXXMMXMXMXMXMAMAMAMSSSXSAAXMMSSMAMXSXAAXSAMXSMMMSSMMXMSXAAAXASAAXSXXMAAASMMMXSXSAAMXSAMAAAASXXAAMMSXSXSSMXSASMAMAXMAMXMXMMXX
XSAMXMXSMMSMMMXXMAMMMXXAMAMASASASAMAAXAMMMMAAAAMAMMMMSSMSASXSAAAXMAMXXAMMSXMMMXSMMASMMMMMXSXMMMSXXXAMXMSMSMSASMMSXMASXMAMAMMMAMMSMMSAMXMXXXX
MXXMASASAMMMAMMMXSAASMSASASMSXSMSASMSMMMSMSMMSMSASXAAXAMXXMASMMMSSSMSSMSAMMSXAAMAMAMAXXXXAXAAAAXAXMXSAAXAXXMXMASXASAMASXMASAMXMAXXXMMMAMXMMM
MMSMXMASXMSMXSASAMSMAASASASASAMASAMAXXXXAAAXXMASASXMSSMSAAMXMASXMAAAMMXMAMASMMXSAMASXMSXMASXMMSMMXMASMSMSMXXASMMSXMAMXMAMXMXXSMXXSXMASAMXASX
AAAAAMXMMXSAASMMXMASMAMAMXMXMMMAMMMSMMMAMSMMAMXMAMAMXMASXXSASAMAMSMMMXMMAMMMAAAXASASAMXAMASASMXAXASAMAAAAXMXXMSAMXMSSMMSMMMSMMMMMSASASAXSAMX
MMSSXSMMSAMMMMMSMXMSXAMSMSMSMXMXSXAAAAXXXMMSSSMMXSAMXMXMMMSXMASAMAMSMAXSAMXSMMSSMMASMMASMASMMAMMXMMAXSMSMMMSMMMMSAMAAXAAMAAAAAAAXSAMAMXXMASX
SXMMASAXMMSMXAXXASMMXSSXAXAAMASMMMSSSMSXAXAAAAXSASASXMMAAAXMAMXMXMXAMAMSASXSXAMAXMMMASAXXXMMMSMSSMMSMXAMXSASAAAASAMXSMSSXMSSXSSSMMMMMMSXMMSA
XAASAMSXXXMXSMSMXMXAAXMMMMSMSASAXXAAXAMMMMMMMSMSASAMASMXMXSAMXXMASXMXSASAMAMMMSMMMXSAMMSMMSMXMAMAAAMAXXMXMAXSMMMSAMXXAXAAXAMAMAMAMXMSAMASAMM
SXMMAMMMMSXAMSAXASMXMMXAAXXAMMSXMMMXMXMAXAXMMXASXMAMAMMMSAMXMXMMAXAXXMAMAMAMSAAMASAMSSMAMMAMMMAMXMMSSMXAXMSMXMSAXMSSMMMSMMASXMAMASAXMASXMASX
AMSXSMAAAAMXSAMSXSAAMMMSMSMAMASASXXXAASASXSXAMAMAMXMAXAMXXMASAAMXSSMSMAMXSSXMSAXAMASAXXASXASXSMMXSAMAMMXSAMXAXMMMMSXASAMXMMMMSMSAMXSXAXASAMA
SASAXXSMSXAXAXMAAXXSAAAAXMMMMMSAMAMMSMMMAMMMXMAMAMXXMSMSMMSASMSMMXMAMSASAMXAMMXMASMMMMMMMMAMXAMXAMXSAMMMSAMMMSASXMXXAMXAAXAAXAAMMSMMMMSMMMSS
MAMAMAXAMXSASXMSSMMXAMSSSMSMMMMMMXMAAXMMMXAMAXSMSMSMAAMAAAMASAXXAAMAMMXSAMSMMAAXMMXAAAASXSASXSAMXSAMASAAXAMXAXAXMASMSSMSMSMSSSSMAAAXAMAMAXAA
MXMSMSMSAAXSMAXAAAXSSXAMMAAAAAMMSXMSSXMASMSSSXMAXASMSMSMMMSASASASMXMXSXMXMAMMMXAAASXMSAXXMAMAMMXMXMSMSMSSSMXMMSSMXMAMAMXXXXAXAMMXXXXMMMMXSSM
MXMXAXAMMMSASMMSSMMMMMXSMSMXMXXASAMXAMSSMAAAMAMXMXMAXAXXXXSAXAAMXMASAAAAMSXMAAMMMMSAMMMMMMSMMMMSMMXAMXAAAAAXXAAAAXMMSSMXXSMSSMXMSSSSMXSAMXAX
MXMMMMXMMXMMAMXAXAMXXMMMAMAMSSMMSAMMMMXAMXMAMMMXAXMXMXMMSMMMMMMXMSXMASXMAMMSMXSXAAMMXAAMMAAAASAAAXSMSMMMSMXSMMSSMMSAAXMXMMAAXXXMAAAXAAMAMXMX
SXMAMSMXXMMMSAMXSXMXAMXSMMAMAAAMMAMMXMSXMASMXAMXSAMASAXMAXAMXAXAXXAMAMXAXXAAXAXXMMMXXMMMMSMSMSXSMMXXAAAAXAAAMAMAAASMSMSAMMMMSSSMMSMMMMMAAAMM
AMMAMMXXASMAMXSXMASMMAAXMSXMSSMMSSMSAMXAMASAMXAMXMSASMSSMXXSSSSMASAMSSMSMMSSMMMAMXXSXXAXMMMMASAXMSSSXSMSSMMXMASMMMXAAMMAMXMASMAAAAXXXAXSSMSA
MSSMSAAMSMMMSMMASAMASAMXMMMMAAAXAMASXSXSMMSMAMAXAXMXSAXAMSMXMMAMMAMMMAAMXXAMXXAMXSMSASMMMAAMAMAXMAAMMMXMXMSXSASXSMMMMSSSMMMSSSSMMMXAMMXMAMAX
XMAMXMASXAAAAAXXMASXMMMAAXSSSMXSAMXMASAMAMXXMASMMXSAMXMAMAAAMSXMMMXAXAMXXMSMASXMMXAMAMAASMMSAMXMMMSMASAXMXMAMAMAAXMAXXAAAMSMMXXAXXMXXMAMAMMS
SMMMSXMAXXMMXMXSSMMASASXSMAAAXASAMXAASASMMSXMAMAMXMMSXSAMMSSXMASXXSXXSAMXSAMXXAAMXMMAMXMMAASAMXSMMMXAMAXXSMMMXMSMSXMXMXMMMMAMXSXMXMASMMSMSAS
MAAAXMAMAMSXSMXMAASMMAAAMMMMMMASAMXMAMAMAXXAMASAMMSXXAXAXXMAXSAMMXMASMAAXSMSXSSMMMASXSASMXMSAMAAAAAMMXSASAAXXMAMAAXAMSAMXMXAMMMSAXMXXAAXAMAM
SAMSSXXXMAMAAXAMMMMXMSMAMAAMXSASXXXXXXSSMMMXMASAMASAMXMMMSAMXXSXXAMXMSMSMMXAAXMASXMMASASAMASXMASMMSSMAMXAXMMMSAMXMXXAXMXASMXSAXXSXSMSMMMMMSM
MXAAMMMXMAMXMSXSXMASAMXXMXMSASAMXSMMMMMAMMSAMXSAMSXMXAXMAXXXMMMMSSSSXMXAAMMMMMSXMAXMAMASAMXMXMAXAAAXMXMSXXSXMASAXSMSSXMMMXAASXMXMXMAAXMXSAMX
MMMASXSSSSSSXAASXSAMASXSSMAMAXAMAMMAAAXXMASASASXMAMXSMSMSSXXMASASAAMMSMSXMAAAAXMMXMMXSASXMSSMXAXMMSMMSXMAXXAMXMMSXAAXAMASMMMMXMASAMXMSAXMASA
XASAMAMAAXAAMMSMAMXSAMAAAMAMMXMMXXASMMSAMXXMMMMXSSMXAAXXMXMXMSSSMMMMAXMMMMSSMSSSMSAXAMAMAAAAAMSSXAMAAXAMXMSSMMAXMMMMSXSAMAAMXSMMSASAAMMMMXMX
SXMMMAMMMMMMXSAMXMMMXMMSMMMSMXSASAAXAXMASAXSASXMAMASXAMXAAAXMAMMMMMMSSXSAXAXXAAAASXMMSSSMMMMMMMAMSSMMSSMXMAMASXXAAAXXMMMSSMSAAMAXAMMXMXASAMX
MMSASXSAMXMMXMXXXSXMXSAMXSAXAAMAMMXMMMSAMXMMASMMAXAAXAAMSMSMMASAAMAMXMASXMMSAMMMMMXSXAAAAMXMAXXAXXAAXAMXXMXXMSMSSSSSMXAMAXXMASMSSMMMMMSXXASX
AAXAMASMSMSMMMSSMMAMASAMXMMMMMSASXXMXAMXSSSMAMXSMMSSSMSMMMAASMSMSMSSMMAMAMAXMXSXXXSAMSSXMMXXAMSXMSXMMSSMMMSXXMAXAAXAXSSSMSMSXXAMXXAXAAMXSAMA
MMSAMXMXAAAXXAAAASXMASAMXXAAAASXSASAMXSXMAAMMMXAXAMAMXAAXSSXMMSXMAXXAMXSAMXXXASXMMXMAXMASMSMXXMAMAASAMAAAAXXAMXMMMMSMMXAXSAMXMAMXSSSXSAAMAMX
XXAXXXMSMSMSMMSSMMAMXXXAMXSMMMSMMAMSAMXAMXMMMSSMMMMMAXXMMXXXMASAMMMSMMXSMSMSMMMAXAAXMXMASAAXMASXMSSMASMMMSSXSMMXXXAAAXSXMMSMXSAMXMXSAMXXSAMM
XMASMAMMXMXXXAXXMSXSMSXMXAAMMAMAMXMXXMXAMAXAXAAASASXSSSSXMAXSXMMMAAAAAAMXMAAMXSMMSMMSASAMXMMSAMXXXAMXMXMSAMAXASXSMSSSMXASAMXMSXSMSAMXMSASASM
MXAXMAMMMSXMMMMMXXAAXMASMSSXMASXMMSXXMASXMSMSSSMMASAAAAXMXSXMAXAMXMMSMXXAMSMSMMSAAMAXXXASXMAMMSSXXMMSMAMMAMMXAMMAAMMXMSSMMSSMMXMAMAXSAMXSAMX
SXMSSMXSAAAAASASMMMMAMXMAAAASXMXMASXMASXAXAAMMMAMSMMMMMMAAXASMSMSXSAXXXXMXAMXMAASXMASMSMMAMXSXAXASAMXMASMXMXMSMMMXMMAMXMASAMAXSMASMMXXSAMXMX
SASAMXAMMXXMAMASMAAXSSSMSMSMMASXMASMAXMSMMMMMASAMAMSMMMSMMSAMAAXAAMMMXSAMSAMAMXMXXMAMXAMXXSSMMXSMAAMXSMSAASXMAAXXMASASMSMMMSMMAXAXXAMMMXSMSS
SXMASMMXXXSXSMAMMSSMMAMAAXMASAMAAXSMMMMXXAAASXXMMAXMAAAAAXMXMMMMMSMSAMXAMAASMSMSSXMXSSMMSMMAAMAMXAAXMSAMMMSASMXMSSMMAMMAXAAAXAMMSXSAAASXMMAS
SMSAMAXMSMMAMMXMAXXASAMMMMSAMMSMMMMXAAMSMXSMSAMXMMSMMMXXSMMASXSAXAAMSSMSMMXMAAAXMXSAAAXAAAMSXMXSMSSXAXXXXXSMMSAAAMXMSMSSMMSXSSXAMMSXSMSASMAM
SAMXSSMMSAMXMASMMSSXMAMXAMMASASMSMAXMSMAMMMAMMMAASAMXXSAMXSXSAASXMXMMXAXMAXMSMMMMAMMSMMSSMMMSMMXMAMMSMSMSMXSAXMMMSAMXASAAAAXSXMXSAMMXASAMMSS
MAMXAXMASMMMAMMAAXMMSAMSSMMAMXMAXMMSMAXXSAMAMAMSSMASAAMMMXMMMMMMMMMXMMSMAXXMXXMXMASAMXXMAASAMXAAMASMAXAAAXMMSSMSXSASXMMSMMSSMAMAMMMAMMMXMAMX
SAMMMXMAMAMMAMSMMSAAMAXMAAMSSSMSMMAAXMASMMSSMMXMAMAMMSMAMSASXSXAAAXMMAXXASMMAXAASXMASMSMSMMASMSXSAXXMMMSMMSAAMXMASMMAMXXXXXXMAMASXMXSXMSMSMM
XAXASXMASAMSAXAAMMMMSMMMMXMAAXAMAMSMSMMAMXAAXXASXMXSXAXAXSAMAMXMSMSASAMSXSASAMSASAMXXAXXMASAMXMAMMSMSAAAAAMMMSAMMMAMAMAMMASMXXSAXAMAMAAMAAMA
SMSMSASASAMSXSMSMXAAXAAXXSAMXMAMSMMAAXXMMMMSMSXSAMXMXMMMXMSMSMAMAMSXMAMAAMMMXXXAMMSSMMMMXMAMXAMAMSAASMSMMMSAXSXSMSSMASASXASASXMMSMMASAMMMMSS
AXAAXAMXSXMSAXMMMMMSSSMSAMAXMMXMAAMSMMXXAAMXAMXMAMASAXASAMXMAXMMAMMXSXMMMMASXSMSMAAXAAASAMMMMMXAMMMMMAAXXXMMXMXSXAXXXMAAMAMXAXSAAASASXXSAAXX
MSMSMXMXSMMMXMAMSASAMXAAASAMXSASMSMMXMSMSMSMSMSSMSASASASASXSMSXSXMXAMMMSXSAMAMAMMMMSSMXXASXMASMMSXMAMMMXSMMSXMASMMMMMMSMMASAMXMMSXMASXAXMXMX
XAXAAXMASAXAMMXAAMMSSMMSMMMMXSASAMXMAMAMMAMAXXMAMSXSXMASMMAAAAXAXSMSXSAAXMASAMAMASMMMSMSMMMMXSAAAASXSMXAAAAXAMASXXAAMAXXMASAXAXMMMMXMMMXMAAX
SSSMSXSAXXMMMSAMMXMAXAXMASXSMMAXXAAMSSSXSAMAMXSXMMAMXMXMXMSMMMSMSMAMAMMSSMXAMSASASASAMXAXAMSASMMMMMMAAMSSMMSAMXSMMMSMASXSMSMSMSMAAAAXAXAXXSA
MAMAMMMAXSAMASXXAAMXXSASMXAAXMAMMSXAAAMASASMMASMMMAMMMMSAMXMAMAXAMAMXMAXAXMAXAXXXSAMXSSMXMSMASASXXAMMMMXAXAAASAMASAAMXMASAMAMXAMMMXMMMMAASMX
MAMAMAMAMXAMASXMSXSMMMASXMSMMXXMMAAMMSMASMMAMAXAASASAASMMMXSAXSSSMSSSMSSSMSAMXMXMMMMXMXXAMAMXMASMSSSXXSXMMXXXMAMAMSSSSMMMAMAMMSSMSAMSSSMMXAX
SXMAXMMASXMMMXMAAAXAAMAMAXXAAAMSXMMSAXMASMSASMSSMSASMSMMXSAXMSMAAAXAMAMAMAXMAAMAAAAMSMSAMMAMXSAXMAXMAMSXSASMSSSMAMAMAAASMXMXSXAAAMASAAAXXMXM
SMSMSASAMAMAMSXXMSMMSSMSSMMSXXAAAAAMXXMXSASXMAXMASXXMMXXSMAXSAMMMMMAMXMAMMMMMSSMSSXSAAXASXSMMMASMMXSMMMXMASAAAMAXMASMXMMAAXAMMMMMSSMMSMASMMM
MAAAAAMMMASASAAMXMAXAAXAMXAXMASXSMMSMSMXMAMAMXMMMMMSXMASMMMMMASMMSMSMMSAMMASAMAXAAMMMXMAMAAMSMMMMMASAXSAMAMMMMSSMSMMXSMXXMMASAXXAXAAMAMAXXAX
MXMSMMMSSXMXXMXAAXXMSSMSSMMSSMXAMMMAAAAMMXMMSMAMMAAXAAXAMASXSMMAMSAAAASAMSAMASXMMSSMASAMMSMMXAXSAMMSAMMAMXSAMXAAAXXSAMSAAMSAXMSMMXMXSASMSSSS
XSXXXXAXMMSSSXSSSSMMXMAMAXMAASMMMMXMSMSAMXMAAXMMSMSSSMSMSASAAXMAMMMMMMSAMMASAMAAXAMXMMMSAMAXSMMSASAMAMSXMAXAMXSMMXXMAMXSSMMMSXAASAXMSXSAAXAA
MAAMAMXXAXAAAAAXMAMMAMMMAXMSMMXMASAMXMAXXAMXMMSAAXXAMAAXMASMMMSSSMSSSMXAMSMMMXSXMMSMSMXMASMMMAASXMMSMMAMMSSMMMXAMAMSSMAMXMASAMXMMASAMXMMMMMM
AMXSMMXSMMMSMMMMMAMXAMAMASXMAAMMASASAMXMSMSSXAMXSSMMMMMMSASASAXAAMAAMASAMXMAMMMASXMAAAASMMXAXSMSMMAMAXMXAAXAAASAMXAAXXXMASXMXSSSSXMXXMMAXASX
SXAAXMAAAAMAAAXAMSSSXSXSSMMSMMSMASAMXMAMXAAXMMMXXMAXAXSXMASXMAMMMMSSMMMXMASASMSAMMMMMSMMSASMMXAXAMAXMMXMMSXXMXSMASMSXMMMMMSMXAAXAXAXMMSAMXSS
MMXMAMSSSMSSSMXSXXAMAAXSAMAXAAXMAMXMASXMMMMXSAAMSSMMMXMAMXMASXMSXAAXAXMASMSASAMXMASXMMAAMMMMSMSMSSMSAMXAAMMSMMMXAXXMASASAAMXMMXMMMXSXAMXMMXX
MSMSMAAXAXMXMMSMMMAMXMXSAMSMMMMMAMXXAMAXXMAASAAXXAXMMASAMASAMAASMMSSMAXASAMXMAMXSAMXSSMMMSAAAAAAAAXSAMXMXSAXAAAMMSMXMMASMXMAXXASXSMAMXMSAMXM
SAAAAMMSMMSAMXAAASAMAXAXXMAMASASMSSMSSMMSMMXSMSXSAMXSASASMSAMMMMAMMAMAMSMMMXSXMMMXSMXASAAMMSSSMSMSMMAMAMMMASMMMMSAAAMSASASMSMMASAMAMSMAMXMSM
SMSMSMAAXAMASXSSMSASXMMSMSMSASXSAAMAMAXAAASAXAAXXMAMMASAMXSAMXAASMXAMASAAAMAMAAASAMXSAMMMSAAAAAXMAXSASASAMXMMAAXSMSMXAAXASAMXMAMAMSXXMASAXAX
XMAXXMMSMMSXMXMAXSXMXXXAXSXMXSXMMMMASXMASAMXMXMMAXSAMXMXMASAMSSSXASXSXSMSMMXSMMSMXSAMXMSAAMMSMMMXAASXSXSXSMMSMMMMXAXAMXMXMAMXMASAMAXXMASXSAS
` // Add your full input here as a multi-line string

	// Prepare the grid
	lines := strings.Split(strings.TrimSpace(puzzle), "\n")

	// Count occurrences in horizontal and vertical directions
	total := 0
	for _, line := range lines {
		total += countInLine(line) // Count in rows
	}

	for col := 0; col < len(lines[0]); col++ {
		var column string
		for row := 0; row < len(lines); row++ {
			column += string(lines[row][col])
		}
		total += countInLine(column) // Count in columns
	}

	// Count occurrences in diagonal directions
	total += countInDiagonals(lines)

	fmt.Printf("Total occurrences of 'XMAS': %d\n", total)
}
