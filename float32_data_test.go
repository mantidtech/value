// Code generated by valuegenerator. DO NOT EDIT

package value

var testFloat32GetTestData = []TestFloat32GetTest{
	{
		name: "basic",
		a:    ptrToFloat32(7),
		want: float32(7),
	},
}

var testFloat32SetTestData = []TestFloat32SetTest{
	{
		name:    "basic",
		s:       "7",
		a:       ptrToFloat32(7),
		wantErr: false,
	},
	{
		name:    "not a number",
		s:       "dooker",
		a:       nil,
		wantErr: true,
	},
}

var testFloat32StringTestData = []TestFloat32StringTest{
	{
		name: "basic",
		a:    ptrToFloat32(7),
		want: "7",
	},
}

var testBindFloat32TestData = []TestBindFloat32Test{
	{
		name: "basic",
		x:    float32Ptr(7),
		want: ptrToFloat32(7),
	},
}
