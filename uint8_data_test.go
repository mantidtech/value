// Code generated by valuegenerator. DO NOT EDIT

package value

var testUint8GetTestData = []TestUint8GetTest{
	{
		name: "basic",
		a:    ptrToUint8(7),
		want: uint8(7),
	},
}

var testUint8SetTestData = []TestUint8SetTest{
	{
		name:    "basic",
		s:       "7",
		a:       ptrToUint8(7),
		wantErr: false,
	},
	{
		name:    "not a number",
		s:       "dooker",
		a:       nil,
		wantErr: true,
	},
}

var testUint8StringTestData = []TestUint8StringTest{
	{
		name: "basic",
		a:    ptrToUint8(7),
		want: "7",
	},
}

var testBindUint8TestData = []TestBindUint8Test{
	{
		name: "basic",
		x:    uint8Ptr(7),
		want: ptrToUint8(7),
	},
}
