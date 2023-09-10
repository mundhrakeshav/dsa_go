package stacknqueue
type FirstCircularTour_FuelDist struct {
	Fuel, Distance uint;
}

func FirstCircularTour(input []FirstCircularTour_FuelDist) int {
	start, extraFuel, reqFuel := 0, 0, 0;

	for i := 0; i < len(input); i++ {
		extraFuel += int(input[i].Fuel) - int(input[i].Distance)
		if extraFuel < 0 {
			start = i + 1;
			reqFuel += extraFuel;
			extraFuel = 0;
		}
	}
	if reqFuel + extraFuel >= 0 {
		return start;
	} else {
		return -1
	}

}