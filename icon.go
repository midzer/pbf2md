package main

import "strings"

func getIcon(shop string) string {
	icons := []string{
		"alcohol",
		"art",
		"bag",
		"bakery",
		"beauty",
		"bed",
		"beverages",
		"bicycle",
		"bookmaker",
		"books",
		"butcher",
		"car_parts",
		"carpet",
		"car_repair",
		"car",
		"charity",
		"chemist",
		"clothes",
		"coffee",
		"computer",
		"confectionery",
		"convenience",
		"copyshop",
		"dairy",
		"deli",
		"department_store",
		"diy",
		"doityourself",
		"electronics",
		"fabric",
		"florist",
		"furniture",
		"garden_centre",
		"gift",
		"greengrocer",
		"hairdresser",
		"hifi",
		"houseware",
		"ice_cream",
		"interior_decoration",
		"jewellery",
		"jewelry",
		"kiosk",
		"laundry",
		"marketplace",
		"massage",
		"medical_supply",
		"mobile_phone",
		"motorcycle",
		"musical_instrument",
		"music",
		"newsagent",
		"news",
		"optician",
		"outdoor",
		"paint",
		"perfumery",
		"pet",
		"photo",
		"seafood",
		"second_hand",
		"shoes",
		"sports",
		"stationery",
		"supermarket",
		"tea",
		"ticket",
		"tobacco",
		"toys",
		"trade",
		"travel_agency",
		"tyres",
		"variety_store",
		"video_games",
		"video",
	}
	shopSplit := strings.Split(shop, ";")
	_, found := find(icons, shopSplit[0])
	if found {
		return shopSplit[0]
	}
	return "other"
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
