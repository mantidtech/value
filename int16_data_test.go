// Code generated by valuegenerator. DO NOT EDIT

package value

var testInt16GetTestData = []TestInt16GetTest{
	{
		name: "basic",
		a:    ptrToInt16(7),
		want: int16(7),
	},
}

var testInt16SetTestData = []TestInt16SetTest{
	{
		name:    "basic",
		s:       "7",
		a:       ptrToInt16(7),
		wantErr: false,
	},
	{
		name:    "not a number",
		s:       "dooker",
		a:       nil,
		wantErr: true,
	},
}

var testInt16StringTestData = []TestInt16StringTest{
	{
		name: "basic",
		a:    ptrToInt16(7),
		want: "7",
	},
}

var testBindInt16TestData = []TestBindInt16Test{
	{
		name: "basic",
		x:    int16Ptr(7),
		want: ptrToInt16(7),
	},
}
