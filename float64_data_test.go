// Code generated by valuegenerator. DO NOT EDIT

package value

var testFloat64GetTestData = []TestFloat64GetTest{
	{
		name: "basic",
		a:    ptrToFloat64(7),
		want: float64(7),
	},
}

var testFloat64SetTestData = []TestFloat64SetTest{
	{
		name:    "basic",
		s:       "7",
		a:       ptrToFloat64(7),
		wantErr: false,
	},
	{
		name:    "not a number",
		s:       "dooker",
		a:       nil,
		wantErr: true,
	},
}

var testFloat64StringTestData = []TestFloat64StringTest{
	{
		name: "basic",
		a:    ptrToFloat64(7),
		want: "7",
	},
}

var testBindFloat64TestData = []TestBindFloat64Test{
	{
		name: "basic",
		x:    float64Ptr(7),
		want: ptrToFloat64(7),
	},
}
