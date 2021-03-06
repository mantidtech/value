// Code generated by valuegenerator. DO NOT EDIT

package value

var testStringGetTestData = []TestStringGetTest{
	{
		name: "basic",
		a:    ptrToString("foo"),
		want: string("foo"),
	},
}

var testStringSetTestData = []TestStringSetTest{
	{
		name:    "basic",
		s:       "foo",
		a:       ptrToString("foo"),
		wantErr: false,
	},
	{
		name:    "empty",
		s:       "",
		a:       ptrToString(""),
		wantErr: false,
	},
}

var testStringStringTestData = []TestStringStringTest{
	{
		name: "basic",
		a:    ptrToString("foo"),
		want: "foo",
	},
}

var testBindStringTestData = []TestBindStringTest{
	{
		name: "basic",
		x:    stringPtr("foo"),
		want: ptrToString("foo"),
	},
}
