package main

import (
	"fmt"
)

func main() {
	input := "lrgrvgvttzmtmtgglmgmccpclppvdvtvvllvggvrggbwwlzlmzzbppnvpnvppcjjzhjhthnhjnhhhndhnnnsbnnhzzvhhplplzlrzzgpzpwzpwwsvsjvjfvvphpspwswrswscwscwsscffspsbbjjcjwjrwwtgwwgswswwzbzddqnnpqnpnqppwzwszsnsjjpddhvvcbbhhpzzlpzlzppfpvvmcmvvflfttrltrlldlglbgblltqtffrtrwrzwwzmzwmwwlzzhttwzzwnnmrrcdrdjrjqjvqvvjzzgccrllhmhzzfnfwwtzwzwpwhhdjhhmzzbbvggzdzccbzbbpcpqccjbcbppsttdjdnjnppjjnmmszmzgzddtctvctcvttgtbbzqqggnmmdllvdvmvzzhfffzvfvtfvtvwwcnwnvwwbccggjcjqcqcbcrrppdqppdzpzqppttjhjdjqjppzgzjjpllwrrbttrvvzzbhzzqppndppwqppnrpnnttfwttsrrgprggmtmhmzhzczwzmwzwrwqwrrrdqrrvssnlngnppfqqgbgjjcttbgtbtmtctmcmcmgmsgsffhghqhbbvtbbtltmltlnlpnngcnggbngbnnzgzccgcpgcpcjppnnzjzdjdggzjzljjhnncgcjcscfctcvttvqtqmqjjsqjqpqfqhqmmlvvmppfrfjjngnnfllrlhhppcjcbjcctgcgtcgcvgvffqfcfpcpdpffrbrvbvnnphpqpfqqtnttmtgtlgtgzttnvvpwvwvcwcfwcwmccwlclqlflpflplwpllndlltlqtlqqmqnqmnqnvqvrrtddqndnrdnnpzprrqnnggvqvhvpvptvvvzwzrwwscsqqmcmttbgtgpptzptzzvszvzdvvtsscbbrpptssltssztszttlvlqljlgljlhhwvhwvvqhvqhqrhqqcnqccnbcbppbffzqfqsfspsqsjjrhjjchcmhmnhmmzjmjmfjmmsbsvvgcggtdgghchrrpnnrttnthtdtmmhmdmppmgpgllrwlrwlwvvlmlglppzttsvsbsnbncnjnffddzcddbzzbzgbghhhtltwtggljjggsdswwpmmfhfsfvfrrgmrgrfggvzzbnbttwqqdcdppqcqpcpqpjqpjpbbgjbgjjfwfwpfpgpzgzmzgzdzzpwzwqqjqfqllgrgjjfvvqnvncntngnhgnhgnnzvvbsbmbqmqwmqwwhbwhhsccvhcclncnqccnvnzvvdgvgnvnttmbbhccwgwttlwtwqttqcqmcqcdcmmjpmmjsjhhprrnnqddjwdjjvvhvgvssthhnfhnnntfthhtggthhbrbrjbbjfbjjrgrsrjrqqqfwflfclflnnnnvggfqgqzzbbvttfcfvcvsswvssnzndndvnvqqznnrjnnsmmptmppncpchcctwtbbgbqqjqtqsqfsfvfvnvmvzzpgzppdzdvdqdjdnjnttvvjbbzrzqrqwrqrbqrqsqpspjssnqnpqqnjndjjzmmvbbrqrccrffhwhggbttpnpphwhhmrrndrnddzqzzfbfwbwnwtwjjwjmjsjcjgcjjfcftcffvpvwwbffgzgnnlfffnddtdbdlbbcjbjmmfpfzfbbwbdwwfmfpmmfjfffvzvdvvhrvrcvcscjjpfjjnfnzzrtrpphtppzrppwhhphthltlllttghgwwvlwlflhldlzzmbzzjppnwppvlplqqbtbwwccswccqzzjhjbbhbnhnshnsslmmlqqjfjrjjmvvhpjqhzqffhsdsbwpjvgpvmbfqltrmpnwfcptpfmtjcpbzfldbhcmzchshrlbjgggrfjcqhzqqvbzsczmbgqmzqmltlrtlbnsfvmlhbbcqbbltjpdrpznrglshvgdnqwlhthghvtbffddcjwgdzfswzbppjtdhstcqqmvzmjrvfjbhmrznwqczdjjclnhbmtdvvzwttwnrlfqwpglpcppdwdcvfqpqfnmbvzvmqlmnlgnrsqdjvtsftgnlrtzsrcqhltmhzhpmzqqfqrjwhqfnqdtnshwgfhcpjrlplnqczdlntnhsczrgfhflsfbmftsbptflqbpwblrfnfzvqtpblftmscpzgdhhsbdbjhqclnptwtmhbbfglmvwnbqgvqhmmswwjpfwqjbvznmcpdzcvbzjmfqnwstvvtdnlvnpznnblfqzjjrjgnsbtmmbjzsvmgwddtnzcvhvtdrmjgtcrjzznrssscrzcfbfpgpnpppsqcqpccnbdjnwrbvhrcwgqncjrzbdhzqpfhqbnvbfrzmlfbfvtpggrtdswnvlsvpjsmfchhpbbszbnqqfrmhpqzdjhmhmnnmplbtrpgphvvqdfbcfnrfrbfbtshlmlfltjnbmggqntvhdnlvtcvlhmlrlfzfrqmlwqzrdghvdvtsqvmpdjrjclmlmgjqwzzldnzvfmwmrrnfghsvpcwjdtlnrhpjczwpgfbhpnmcbpthsndfflbjhnlwdbbmlttfqcmswvppslptgzbvfgppvpnhjccrpgrpwtngmmccjghhcwddmnglschnpjwqtrtsvggnpzvsqshfvcnhptphtlmqmpznfzwvbnhwpsfwvpflsdjcjgfzjprbbfzgdbmrjgwrgfdphghrhnpvfncrdzcwtthmqtdwlhjsdthqpzhbjpgggndtrmwvcsqhzrzwbhtqsqthvqncprvnpsrlpvlvcjrcflhbdhrfthlfnqbzbmvlvhmbjnbbjhpjwlfflfhpfwcwnnsljthvzwprqjmgpldlzjnjtjfjrgnrpzpvzfcsrprbjhwnmccwhppjrlnndjdjzqwpcwnvqwgmnwbrjqqvbplvsncnmdfrbhrrhghfllhrghzmlnltgdsqlgbvnlchgcbqlpqptdwmsjpqrprlhqmstzjfnzgbgvlfshwpcrgzcqmmfwvhwlsdvplmdgrtfrjwpfvhnjqdbwsfcqhchstlzfpdljgvcqsfcnqccnpmvsqbmwjtzwhpglhbjwzmvgqwjhvwfhnlbtsgljzmlldcpjwdcfppmnmphdmhpmdqwwtjtrdhlrjlvzgpbcgvwcmtclgpqwhtpbdtdbdscfzbrzmgjlbppcnvphphfnvzdzzlvfsvsgbgqcnlqwmtcrpwzcvnmnvtmcdsstvqpqzdpvtdsbvtwhdvgzqmzvwlspgbwmlnsrqdqnjwrllncflqsrzdqtjqvpnpjlqfwqtlqfqwlltszcwtpmjtldjgvmvptpmzqhwmlvjgnntpvcslmhlhdbjtjjnvsbnzwtdclwbzrvlqzjljtbdjvwgbwcltvnbhfvtgqrbmzbbfvldhmdvfvtlqglnblfmmpjqmzlnfjltsqdrgmlhbhngrrmhnjndggsdcfmtssmmtmzvhzrmwjsqjcvbsgqgtvdmvqlvlrvglrtlshfdmfrmljjggwjbcsztsjmjftcbbjwrmgqvssrvtgzcgthtlgsjspfmdgwptjdrbswqlpfsbtjlnhllmjpbfhgpfcprpdnqqvqdmcbqhbcqtstvnjdzwzwvhhwmcvcfbdwczpwpdhvnstjnbblbprzsccmwrzgfhmrpvzfztvsrtncdhzhptpfqtnqwvqtwdpvcqztgjgrcbdnvqftphtfbtqdhrffdrdmwsbpvhshzvjbvsrljnzddmmfgcnfdssvzdbsfwmfjsdnslbrqsqfwfqbqszjwvgcjbhrfjcnlfhzvhcbbbpmhhvjdtgrqlcchqtvnhlrgtssllvgcdjrlzlzfbrrrvwvvcgfjdlpscsqljmmwmvwnvrgdmgcbvmwmgprbfrbgptlfjbhrmczwrzwbdhdvtgvldnzfgcngdfhbgqsfzlrbwbvdflrrsrcwthjzvgmdtndgtsjtswfbdqvcjtsdvrvqpmmdlghsdbzplgpfnstplpjdvttgzmnhssftqcqjvdvvdrmltbrpsjvqwbljrqrtqldzbwzznsdstvmdzbrvvtgrrphmbrzwnjbmqvfhljcdlbzqtcbjsfqdqcr"

	marker_size := 14
	marker := make([]int32, marker_size)

	for i, c := range input {
		marker[i%marker_size] = c
		// fmt.Println(last_four)
		if uniqueSlice(marker) {
			fmt.Println("unique! ", i+1)
			return
		}
	}
}

func uniqueSlice(s []int32) bool {
	count := make(map[int32]int)
	for _, v := range s {
		if v == 0 {
			return false
		}
		count[v] += 1
		if count[v] > 1 {
			return false
		}
	}
	return true
}
