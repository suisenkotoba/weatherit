package weatherforecast

type Domain struct {
	ID          int    
	Name        string 
	Description string 
}

type Repository interface {
	GetCurrentForecast(lat, long float64) Domain
	GetTargetDTForecast(lat, long float64, dt1, dt2 int64, mode string) Domain
}
