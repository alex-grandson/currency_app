package usecase

import (
	"devops_course_app/internal/entity"
	"net/http"
)

type (
	InfoReq interface {
		InitRequest(dateFormatted string) (*http.Request, error)
		SendRequest(r *http.Request) (*http.Response, error)
		DecodeResponse(response *http.Response) (*entity.ValCurs, error)
		FindCurrencyRate(currency string, currencyRates *entity.ValCurs) (string, error)
	}

	InfoContract interface {
		GetCurrencyRate(currency string, date string) (string, error)
	}
)
