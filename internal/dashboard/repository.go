package dashboard

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetDashboard() (*DashboardResponse, error) {

	var response DashboardResponse

	// Total Employees
	if err := r.db.Table("employees").
		Count(&response.TotalEmployees).Error; err != nil {
		return nil, err
	}

	// Active Employees
	if err := r.db.Table("employees").
		Where("status = ?", "Active").
		Count(&response.ActiveEmployees).Error; err != nil {
		return nil, err
	}

	// Departments
	if err := r.db.Table("departments").
		Count(&response.Departments).Error; err != nil {
		return nil, err
	}

	// Today's Present
	if err := r.db.Table("attendances").
		Where("DATE(check_in) = CURRENT_DATE").
		Count(&response.TodayPresent).Error; err != nil {
		return nil, err
	}

	// Pending Leaves
	if err := r.db.Table("leaves").
		Where("status = ?", "Pending").
		Count(&response.PendingLeaves).Error; err != nil {
		return nil, err
	}

	// Paid Payrolls
	if err := r.db.Table("payrolls").
		Where("status = ?", "Paid").
		Count(&response.PaidPayrolls).Error; err != nil {
		return nil, err
	}

	// Pending Payrolls
	if err := r.db.Table("payrolls").
		Where("status = ?", "Pending").
		Count(&response.PendingPayrolls).Error; err != nil {
		return nil, err
	}

	// Today Absent
	response.TodayAbsent = response.ActiveEmployees - response.TodayPresent
	if response.TodayAbsent < 0 {
		response.TodayAbsent = 0
	}

	return &response, nil
}