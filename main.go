package main

import "fmt"

func main() {
	var day uint16
	var month uint16
	fmt.Scanln(&day)
	fmt.Scanln(&month)

	switch month {
	case 1:
		if day >= 21 {
			fmt.Println("acuario")
		} else {
			fmt.Println("capricornio")
		}
		break
	case 2:
		if day >= 19 {
			fmt.Println("piscis")
		} else {
			fmt.Println("acuario")
		}
		break
	case 3:
		if day >= 21 {
			fmt.Println("aries")
		} else {
			fmt.Println("piscis")
		}
		break
	case 4:
		if day >= 21 {
			fmt.Println("tauro")
		} else {
			fmt.Println("aries")
		}
		break
	case 5:
		if day >= 21 {
			fmt.Println("geminis")
		} else {
			fmt.Println("tauro")
		}
		break
	case 6:
		if day >= 22 {
			fmt.Println("cancer")
		} else {
			fmt.Println("geminis")
		}
		break
	case 7:
		if day >= 23 {
			fmt.Println("leo")
		} else {
			fmt.Println("cancer")
		}
		break
	case 8:
		if day >= 23 {
			fmt.Println("virgo")
		} else {
			fmt.Println("leo")
		}
		break
	case 9:
		if day >= 23 {
			fmt.Println("libra")
		} else {
			fmt.Println("virgo")
		}
		break
	case 10:
		if day >= 23 {
			fmt.Println("escorpio")
		} else {
			fmt.Println("libra")
		}
		break
	case 11:
		if day >= 23 {
			fmt.Println("sagitario")
		} else {
			fmt.Println("escorpio")
		}
		break
	case 12:
		if day >= 22 {
			fmt.Println("capricornio")
		} else {
			fmt.Println("sagitario")
		}
		break
	}
}
