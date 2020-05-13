// Code generated by valuegenerator. DO NOT EDIT

package value

var testUintGetTestData = []TestUintGetTest{
	{
		name: "basic",
		a:    ptrToUint(7),
		want: uint(7),
	},
}

var testUintSetTestData = []TestUintSetTest{
	{
		name:    "basic",
		s:       "7",
		a:       ptrToUint(7),
		wantErr: false,
	},
	{
		name:    "not a number",
		s:       "dooker",
		a:       nil,
		wantErr: true,
	},
}

var testUintStringTestData = []TestUintStringTest{
	{
		name: "basic",
		a:    ptrToUint(7),
		want: "7",
	},
}

var testBindUintTestData = []TestBindUintTest{
	{
		name: "basic",
		x:    uintPtr(7),
		want: ptrToUint(7),
	},
}