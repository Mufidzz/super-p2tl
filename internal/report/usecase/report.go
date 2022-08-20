package usecase

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"time"
)

func (uc *Usecase) CreateSingleFindingReport(input presentation.CreateFindingReportsRequest) (lastInsertedID int, err error) {
	return uc.repositories.CreateFindingReports([]presentation.CreateFindingReportsRequest{input})
}

func (uc *Usecase) FinishTOSOCheck(userToSoID int) error {
	return uc.repositories.UpdateUserTOSO(presentation.UpdateUserTOSORequest{
		ID:       userToSoID,
		ToSoID:   0,
		UserID:   0,
		FinishAt: time.Now(),
	})
}

func (uc *Usecase) CreateSinglePenormalanReports(input presentation.CreatePenormalanReportsRequest) (insertedID []int, err error) {
	return uc.repositories.CreatePenormalanReports([]presentation.CreatePenormalanReportsRequest{input})
}

func (uc *Usecase) GetPerformanceKwhReport() (res []int64, err error) {

	dbData, err := uc.repositories.GetPerformanceKwhReport()
	if err != nil {
		return nil, err
	}

	month := []string{"", "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	monthOnlyRes := map[string]int64{
		"Jan": 0,
		"Feb": 0,
		"Mar": 0,
		"Apr": 0,
		"May": 0,
		"Jun": 0,
		"Jul": 0,
		"Aug": 0,
		"Sep": 0,
		"Oct": 0,
		"Nov": 0,
		"Dec": 0,
	}

	for _, v := range dbData {
		_time, err := time.Parse("2006-01-02T00:00:00+07:00", v.CreatedAt)

		if err != nil {
			return nil, response.InternalError{
				Type:         "UC",
				Name:         "Report",
				FunctionName: "GetPerformanceKwhReport",
				Description:  "failed parse time",
				Trace:        err,
			}.Error()
		}

		monthOnlyRes[month[int(_time.Month())]] = v.PemakaianKwh
	}

	for i := 0; i < len(month); i++ {
		res = append(res, monthOnlyRes[month[i]])
	}

	return res, nil
}

func (uc *Usecase) GetListTemuanMangkrak(pagination presentation.Pagination) (res []presentation.GetListTemuanMangkrakResponse, err error) {
	return uc.repositories.GetListTemuanMangkrak(pagination)
}

func (uc *Usecase) GetListTemuanMangkrakCount() (res int, err error) {
	return uc.repositories.GetListTemuanMangkrakCount()
}
