// +build all puzzle palindrome test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// PalindromePhaseTestCase struct
type PalindromePhaseTestCase struct {
	Data     string
	HasPhase bool
}

// PalindromeTestCase struct
type PalindromeTestCase struct {
	Data     string
	Expected string
}

// TestPalindrome tests palindrome
func TestPalindrome(t *testing.T) {
	negativeInt := -1111111111
	testPalindromeNumber(t, 12345, false)
	testPalindromeNumber(t, 9245224529, false)
	testPalindromeNumber(t, 543212345, true)
	testPalindromeNumber(t, 33333333333333, true)
	testPalindromeNumber(t, uint64(negativeInt), false)
	testPalindromeNumber(t, 0, true)
	testPalindromeString(t, "input", false)
	testPalindromeString(t, "abba", true)
	testPalindromeString(t, "ZZZZZZZZZ", true)
	testPalindromeString(t, "A", true)
}

// TestPalindromePhase tests if a string is a palindrome
// of digits and letters (ignoring spaces and symbols)
func TestPalindromePhase(t *testing.T) {
	tests := []PalindromePhaseTestCase{
		{"A man, a plan, a canal, Panama!", true},
		{"A nut for a jar of tuna", true},
		{"A Santa lived as a devil at NASA", true},
		{"A Toyota's a Toyota", true},
		{"Able was I ere I saw Elba", true},
		{"Amor, Roma", true},
		{"An igloo! Cool, Gina!", true},
		{"Dammit, I'm mad!", true},
		{"deliver 'n' reviled", true},
		{"dioramas 'n' samaroid", true},
		{"Do geese see God?", true},
		{"Doc, note: I dissent. A fast never prevents a fatness. I diet on cod", true},
		{"Dog, as a devil deified, lived as a god.", true},
		{"Eva, can I stab bats in a cave?", true},
		{"Go hang a salami, I'm a lasagna hog", true},
		{"Ha! Was it a car or a cat I saw? Ah?", true},
		{"Live on time, emit no evil", true},
		{"Ma is as selfless as I am", true},
		{"Madam in Eden, I'm Adam", true},
		{"Mr. Owl ate my metal worm", true},
		{"Never odd or even", true},
		{"no 'x' in Nixon", true},
		{"noon", true},
		{"Not nil: Clinton", true},
		{"“Nurses run.”says sick Cissy as nurses run.", true},
		{"On a clover, if alive erupts a vast pure evil, a fire volcano", true},
		{"put it up", true},
		{"race car", true},
		{"Rats live on no evil star", true},
		{"rewarder 'n' redrawer", true},
		{"Rise to vote, sir", true},
		{"Sator Arepo Tenet Opera Rotas", true},
		{"Sex at noon taxes.", true},
		{"stack cats", true},
		{"step on no pets", true},
		{"stressed desserts", true},
		{"Suez 'n' Zeus", true},
		{"swap 'n' paws", true},
		{"T. Eliot, top bard, notes putrid tang emanating, is sad; I'd assign it a name: gnat dirt upset on drab pot toilet.", true},
		{"taco cat", true},
		{"tattarrattat", true},
		{"Tenth gin knight net", true},
		{"You have no name, Manon Eva Huoy!", true},
		{"帘卷晚晴天，天晴晚卷帘。", true},
		{"枯眼望遙山隔水，往來曾見幾心知。壺空怕酌一杯酒，筆下難成和韻詩。途路阻人離別久，訊音無雁寄回遲。孤燈夜守長寥寂，夫憶妻兮父憶兒。兒憶父兮妻憶夫，寂寥長守夜燈孤。遲回寄雁無音訊，久別離人阻路途。詩韻和成難下筆，酒杯一酌怕空壺。知心幾見曾來往，水隔山遙望眼枯。", true},
		{"白影横窗月，月窗横影白。", true},
		{"篱菊粉墙西，西墙粉菊篱。", true},
		{"花枝弄影照窗纱，影照窗纱映日斜；斜日映纱窗照影，纱窗照影弄枝花。", true},
		{"苦思相见翻无语，语无翻见相思苦。", true},
		{"落雪飞芳树，幽红雨淡霞。薄月迷香雾，流风舞艳花。花艳舞风流，雾香迷月薄。霞淡雨红幽，树芳飞雪落。", true},
		{"しんぶんし", true},          // 新聞報紙
		{"にんてんどうがうどんてんに。", true}, // 任天堂在烏龍麵店
		{"わたしまけましたわ。", true},     // 我輸了唷
		{"日曜日", true},            // 星期日
		{"達人の人達", true},          // 達人的人們
		{"This is not a palindrome", false},
		{"このこはねこのこ", false}, // 這孩子是貓的孩子
		{"信言不美，美言不信", false},
		{"日往則月來，月往則日來。", false},
		{"非人磨墨墨磨人", false},
		{"", false},
	}

	for index, test := range tests {
		var pal = PalindromeString{input: test.Data}
		var val = pal.IsPalindromePhase()
		var msg = fmt.Sprintf("expecting palindrome phase '%v' ? %v", test.Data, test.HasPhase)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.HasPhase, val, msg)
	}
}

// TestPalindromeSubstring tests GetPalindromicSubstring function
// See: https://leetcode.com/problems/longest-palindromic-substring/
func TestPalindromeSubstring(t *testing.T) {
	tests := []PalindromeTestCase{
		{"afdjfjdfdjfj", "jfjdfdjfj"},
		{"abc112123xx1xx3211cba", "123xx1xx321"},
		{"abc", "a"},
		{"bbq", "bb"},
		{"vvvv", "vvvv"},
		{"bbb", "bbb"},
		{"", ""},
	}

	for index, test := range tests {
		var val = GetPalindromicSubstring(test.Data)
		var msg = fmt.Sprintf("expecting palindrome '%v' in '%v'", test.Expected, test.Data)
		var pal = PalindromeString{input: test.Data}
		var sub = PalindromeString{input: val}
		var foo = test.Data == val
		var ms1 = fmt.Sprintf("expecting '%v' is palindrome ? %v", pal, foo)
		var ms2 = fmt.Sprintf("expecting '%v' is palindrome", sub)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
		assert.Equal(t, pal.IsPalindrome(), foo, ms1)
		assert.True(t, sub.IsPalindrome(), ms2)
	}
}

// testPalindrome tests palindrome for any input
func testPalindrome(t *testing.T, input Palindrome, expected bool) {
	result := input.IsPalindrome()
	t.Logf("%v:\tpalindrome ? %5v (expected %5v)\n", input.GetData(), result, expected)
	assert.Equal(t, expected, result, fmt.Sprintf("%v : palindrome ? %v\n", input, expected))
}

func testPalindromeNumber(t *testing.T, data uint64, expected bool) {
	testPalindrome(t, &PalindromeNumber{input: data}, expected)
}

func testPalindromeString(t *testing.T, data string, expected bool) {
	testPalindrome(t, &PalindromeString{input: data}, expected)
}
