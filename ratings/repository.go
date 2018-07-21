package ratings

type InMemoryRepository struct {
	countries map[string][]*InputData
	states    map[string][]*InputData
	cities    map[string][]*InputData
	personal  map[string]*InputData
}

func NewInMemoryRepository() RatingsRepository {
	return &InMemoryRepository{
		countries: make(map[string][]*InputData),
		states:    make(map[string][]*InputData),
		cities:    make(map[string][]*InputData),
		personal:  make(map[string]*InputData),
	}
}

type RatingsRepository interface {
	Save(input *InputData)
	LoadRegionData(region *Location) []*InputData
	LoadPersonalData(name string) *InputData
}

func (repo *InMemoryRepository) Save(input *InputData) {
	if data, ok := repo.countries[input.Region.Country()]; ok {
		repo.countries[input.Region.Country()] = append(data, input)
	} else {
		repo.countries[input.Region.Country()] = []*InputData{input}
	}

	if data, ok := repo.states[input.Region.State()]; ok {
		repo.states[input.Region.State()] = append(data, input)
	} else {
		repo.states[input.Region.State()] = []*InputData{input}
	}

	if data, ok := repo.cities[input.Region.City()]; ok {
		repo.cities[input.Region.City()] = append(data, input)
	} else {
		repo.cities[input.Region.City()] = []*InputData{input}
	}

	repo.personal[input.Name] = input
}

func (repo *InMemoryRepository) LoadRegionData(region *Location) []*InputData {
	if region.HasCity() {
		return repo.cities[region.City()]
	} else if region.HasState() {
		return repo.states[region.State()]
	} else {
		return repo.countries[region.Country()]
	}
}

func (repo *InMemoryRepository) LoadPersonalData(name string) *InputData {
	return repo.personal[name]
}
