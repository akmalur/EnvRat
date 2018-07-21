package ratings

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Rating struct {
	Name   string
	Region *Location
	Rating int
}

func NewRating(query *InputQuery, perc int) *Rating {
	return &Rating{
		Name:   query.Name,
		Region: query.Region,
		Rating: getRating(perc),
	}
}

func (r *Rating) String() string {
	return fmt.Sprintf("\"%s\" \"%s\" %d", r.Name, r.Region.String(), r.Rating)
}

func getRating(perc int) int {
	if perc >= 90 {
		return 1
	} else if perc >= 80 {
		return 2
	} else if perc >= 70 {
		return 3
	} else if perc >= 60 {
		return 4
	} else if perc >= 50 {
		return 5
	} else if perc >= 40 {
		return 6
	} else if perc >= 30 {
		return 7
	} else if perc >= 20 {
		return 8
	} else if perc >= 10 {
		return 9
	} else {
		return 10
	}
}

type InputData struct {
	Name   string
	Region *Location
	RValue *decimal.Decimal
}

func NewInputData(name string, region *Location, rValue *decimal.Decimal) *InputData {
	return &InputData{
		Name:   name,
		Region: region,
		RValue: rValue,
	}
}

type InputQuery struct {
	Name   string
	Region *Location
}

func NewInputQuery(name string, region *Location) *InputQuery {
	return &InputQuery{
		Name:   name,
		Region: region,
	}
}

type Location struct {
	country string
	state   string
	city    string
}

func NewLocation(country string, state string, city string) *Location {
	return &Location{
		country: country,
		state:   state,
		city:    city,
	}
}

func (l *Location) String() string {
	result := []string{}
	result = append(result, l.country)
	if l.HasState() {
		result = append(result, l.state)
	}
	if l.HasCity() {
		result = append(result, l.city)
	}
	return strings.Join(result, "/")
}

func (l *Location) Country() string {
	return l.country
}

func (l *Location) HasState() bool {
	return l.state != ""
}

func (l *Location) State() string {
	return fmt.Sprintf("%s/%s", l.country, l.state)
}

func (l *Location) HasCity() bool {
	return l.city != ""
}

func (l *Location) City() string {
	return fmt.Sprintf("%s/%s/%s", l.country, l.state, l.city)
}
