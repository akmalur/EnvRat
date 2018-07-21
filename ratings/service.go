package ratings

type RatingService struct {
	repo RatingsRepository
}

func NewService(repo RatingsRepository) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

func (svc *RatingService) SaveData(input *InputData) {
	svc.repo.Save(input)
}

func (svc *RatingService) GetRating(query *InputQuery) *Rating {
	regionData := svc.repo.LoadRegionData(query.Region)
	personalData := svc.repo.LoadPersonalData(query.Name)

	totalHomes := len(regionData)
	betterHomes := betterInsulated(regionData, personalData)

	rawRating := (betterHomes * 100) / totalHomes

	return NewRating(query, rawRating)
}

func betterInsulated(homes []*InputData, current *InputData) int {
	result := 0
	for _, home := range homes {
		if home.RValue.GreaterThan(*current.RValue) {
			result++
		}
	}
	return result
}
