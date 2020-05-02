package main

func swapInt(param1 *int, param2 *int) {
	tmp := *param2
	*param2 = *param1
	*param1 = tmp
}

func swapFloat(param1 *float64, param2 *float64) {
	tmp := *param2
	*param2 = *param1
	*param1 = tmp
}

func swapFloat32(param1 *float32, param2 *float32) {
	tmp := *param2
	*param2 = *param1
	*param1 = tmp
}
