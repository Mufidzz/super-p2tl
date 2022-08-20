package dbutils

func GenerateParamSequence(currenSequence, numOfSeq int) (res []interface{}) {
	for i := 0; i < numOfSeq; i++ {
		res = append(res, currenSequence+i)
	}

	return res
}
