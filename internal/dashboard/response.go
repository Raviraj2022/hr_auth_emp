package dashboard

type DashboardResponse struct {
	TotalEmployees  int64 `json:"total_employees"`
	ActiveEmployees int64 `json:"active_employees"`
	Departments     int64 `json:"departments"`

	TodayPresent int64 `json:"today_present"`
	TodayAbsent  int64 `json:"today_absent"`

	PendingLeaves int64 `json:"pending_leaves"`

	PaidPayrolls    int64 `json:"paid_payrolls"`
	PendingPayrolls int64 `json:"pending_payrolls"`
}