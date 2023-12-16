package mock

import (
	"miniBankCLI/pkg"
	"miniBankCLI/pkg/models"
)

func FillCities() {
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Dushanbe",
		Region: "RRP",
	})
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Hissor",
		Region: "RRP",
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Vahdat",
		Region: "RRP",
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Varzob",
		Region: "RRP",
	})
}
