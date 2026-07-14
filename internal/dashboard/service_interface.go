package dashboard

type Service interface {
	GetDashboard() (*DashboardResponse, error)
}