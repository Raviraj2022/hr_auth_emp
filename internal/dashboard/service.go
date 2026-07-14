package dashboard

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDashboard() (*DashboardResponse, error) {
	return s.repo.GetDashboard()
}