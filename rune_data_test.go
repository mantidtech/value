// Code generated by valuegenerator. DO NOT EDIT

package value

var testRuneGetTestData = []TestRuneGetTest{
	{
		name: "basic",
		a:    ptrToRune('a'),
		want: 'a',
	},
}

var testRuneSetTestData = []TestRuneSetTest{
	{
		name: "basic",
		s:    "a",
		a:    ptrToRune('a'),
	},
	{
		name: "multi-byte",
		s:    "🤞",
		a:    ptrToRune(0x1f91e),
	},
	{
		name:    "multiple",
		s:       "I see a bad rune rising",
		wantErr: true,
	},
	{
		name:    "empty",
		s:       "",
		wantErr: true,
	},
}

var testRuneStringTestData = []TestRuneStringTest{
	{
		name: "basic",
		a:    ptrToRune('a'),
		want: "a",
	},
}

var testBindRuneTestData = []TestBindRuneTest{
	{
		name: "basic",
		x:    runePtr('a'),
		want: ptrToRune('a'),
	},
}