package main

import (
	"fmt"
	"strings"
)

// Function to count X-MAS patterns
func countXMAS(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Iterate through every possible center
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			// Check if the center is 'A'
			if grid[r][c] != 'A' {
				continue
			}

			// Check diagonals for the X-MAS pattern
			// Top-left to bottom-right
			if c > 0 && r > 0 && c < cols-1 && r < rows-1 {
				if (grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S') ||
					(grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M') {
					if (grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S') ||
						(grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M') {
						count++
					}
				}
			}
		}
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
` // Add full input here as a multi-line string

	// Prepare the grid
	grid := strings.Split(strings.TrimSpace(puzzle), "\n")

	// Count X-MAS patterns
	result := countXMAS(grid)

	fmt.Printf("Total occurrences of X-MAS: %d\n", result)
}
