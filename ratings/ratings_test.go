package ratings_test

import (
	"testing"

	"github.com/AkmalUr/test1/ratings"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func InitSvc() *ratings.RatingService {
	repo := ratings.NewInMemoryRepository()
	svc := ratings.NewService(repo)

	data := []*ratings.InputData{
		ratings.NewInputData("John Doe", ratings.NewLocation("Canada", "Ontario", "Toronto"), NewFromString("1.5")),
		ratings.NewInputData("Samantha Smith", ratings.NewLocation("Canada", "Ontario", "London"), NewFromString("3.7")),
		ratings.NewInputData("Adam Xin", ratings.NewLocation("Canada", "British Columbia", "Vancouver"), NewFromString("2.110")),
		ratings.NewInputData("Monica Taylor", ratings.NewLocation("Canada", "Ontario", "Toronto"), NewFromString("2.110")),
		ratings.NewInputData("Alicia Yazzie", ratings.NewLocation("US", "Arizona", "Phoenix"), NewFromString("5.532")),
		ratings.NewInputData("Mohammed Zadeh", ratings.NewLocation("Canada", "Ontario", "Toronto"), NewFromString("1.43")),
	}

	for _, d := range data {
		svc.SaveData(d)
	}
	return svc
}

func NewFromString(d string) *decimal.Decimal {
	r, _ := decimal.NewFromString(d)
	return &r
}

func TestCanadaJohn(t *testing.T) {
	svc := InitSvc()

	result := svc.GetRating(&ratings.InputQuery{
		Name:   "John Doe",
		Region: ratings.NewLocation("Canada", "", ""),
	})

	assert.Equal(t, "\"John Doe\" \"Canada\" 4", result.String())
}

func TestCanadaONJohn(t *testing.T) {
	svc := InitSvc()

	result := svc.GetRating(&ratings.InputQuery{
		Name:   "John Doe",
		Region: ratings.NewLocation("Canada", "Ontario", ""),
	})

	assert.Equal(t, "\"John Doe\" \"Canada/Ontario\" 5", result.String())
}

func TestUSAZ(t *testing.T) {
	svc := InitSvc()

	result := svc.GetRating(&ratings.InputQuery{
		Name:   "Alicia Yazzie",
		Region: ratings.NewLocation("US", "Arizona", ""),
	})

	assert.Equal(t, "\"Alicia Yazzie\" \"US/Arizona\" 10", result.String())
}
