package dashboard

type Repository interface {
	GetDashboard() (*DashboardResponse, error)
}