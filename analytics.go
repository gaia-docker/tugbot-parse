package parse

type AnalyticsTestSet struct {
	Name     string
	Time     FloatNumber
	Total    int
	Failures int
	Errors   int
	Skipped  int
}

type AnalyticsTest struct {
	Name          string
	Time          FloatNumber
	Failure       string
	Status        string
	NumericStatus int
	TestSet       AnalyticsTestSet
}

func ToAnalyticsTests(junit []byte) ([]AnalyticsTest, error) {
	testSet, err := ToTestSet(junit)
	if err != nil {
		return nil, err
	}

	var ret []AnalyticsTest
	analyticsTestSet := getAnalyticsTestSet(testSet)
	for _, currTest := range testSet.Tests {
		ret = append(ret,
			AnalyticsTest{
				Name:          currTest.Name,
				Status:        currTest.Status,
				Time:          currTest.Time,
				Failure:       currTest.Failure,
				NumericStatus: toNumericStatus(currTest.Status),
				TestSet:       analyticsTestSet})
	}

	return ret, nil
}

func toNumericStatus(status string) int {
	ret := 0
	if status != Passed {
		ret = 1
	}

	return ret
}

func getAnalyticsTestSet(testSet *TestSet) AnalyticsTestSet {

	return AnalyticsTestSet{
		Name:     testSet.Name,
		Time:     testSet.Time,
		Total:    testSet.Total,
		Failures: testSet.Failures,
		Errors:   testSet.Errors,
		Skipped:  testSet.Skipped,
	}
}
