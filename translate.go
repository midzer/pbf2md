package main

import "strings"

func translateShop(shop string) string {
	shopSplit := strings.Split(shop, ";")
	var translatedShop string
	switch shopSplit[0] {
	// Lebensmittel, Getränke
	case "alcohol":
		translatedShop = "alcohol"
	case "bakery":
		translatedShop = "bakery"
	case "beverages":
		translatedShop = "beverages"
	case "brewing_supplies", "brewery":
		translatedShop = "brewery"
	case "butcher":
		translatedShop = "butcher"
	case "cheese":
		translatedShop = "cheese"
	case "chocolate":
		translatedShop = "chocolate"
	case "coffee", "coffee_roasting", "coffeemaker":
		translatedShop = "coffee"
	case "confectionery", "sweets", "cookies":
		translatedShop = "confectionery"
	case "convenience", "food", "grocery":
		translatedShop = "convenience"
	case "deli":
		translatedShop = "deli"
	case "dairy":
		translatedShop = "dairy"
	case "farm":
		translatedShop = "farm"
	case "frozen_food":
		translatedShop = "frozen food"
	case "greengrocer", "vegetables":
		translatedShop = "greengrocer"
	case "health_food", "organic":
		translatedShop = "health food"
	case "ice_cream":
		translatedShop = "ice cream"
	case "pasta":
		translatedShop = "pasta"
	case "pastry", "cake":
		translatedShop = "pastry"
	case "seafood", "fishmonger":
		translatedShop = "seafood"
	case "spices":
		translatedShop = "spices"
	case "tea":
		translatedShop = "tea"
	case "wine", "winery":
		translatedShop = "wine"
	case "water":
		translatedShop = "water"
	// Warenhaus, Kaufhaus, Einkaufszentrum
	case "department_store":
		translatedShop = "department store"
	case "general", "country_store":
		translatedShop = "general"
	case "kiosk":
		translatedShop = "kiosk"
	case "mall":
		translatedShop = "mall"
	case "supermarket":
		translatedShop = "supermarket"
	case "wholesale":
		translatedShop = "wholesale"
	// Kleidung, Schuhe, Accessoires
	case "baby_goods":
		translatedShop = "baby goods"
	case "bag", "suitcases":
		translatedShop = "bag"
	case "boutique", "fashion", "fashion_accessories":
		translatedShop = "boutique"
	case "clothes":
		translatedShop = "clothes"
	case "fabric", "fabrics":
		translatedShop = "fabric"
	case "jewelry", "jewellery", "gemstone", "gold", "gold_buyer":
		translatedShop = "jewelry"
	case "leather":
		translatedShop = "leather"
	case "sewing", "sewing_machines":
		translatedShop = "sewing"
	case "shoes", "shoe_repair", "shoe_repairlocksmith":
		translatedShop = "shoes"
	case "tailor":
		translatedShop = "tailor"
	case "watches", "clocks", "watchmaker":
		translatedShop = "watches"
	case "wool":
		translatedShop = "wool"
	// Discounter, Wohltätigkeit
	case "charity", "second_hand":
		translatedShop = "charity"
	case "variety_store":
		translatedShop = "variety store"
	// Gesundheit und Schönheitspflege
	case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
		translatedShop = "beauty"
	case "chemist":
		translatedShop = "chemist"
	case "erotic":
		translatedShop = "erotic"
	case "hairdresser":
		translatedShop = "hairdresser"
	case "hairdresser_supply", "combs":
		translatedShop = "hairdresser supply"
	case "hearing_aids":
		translatedShop = "hearing aids"
	case "herbalist":
		translatedShop = "herbalist"
	case "massage":
		translatedShop = "massage"
	case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
		translatedShop = "medical supply"
	case "nutrition_supplements":
		translatedShop = "nutrition supplements"
	case "optician", "eyeglasses":
		translatedShop = "optician"
	case "perfumery":
		translatedShop = "perfumery"
	case "tattoo":
		translatedShop = "tattoo"
	// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
	case "agrarian", "agricultural":
		translatedShop = "agrarian"
	case "appliance":
		translatedShop = "appliance"
	case "bathroom_furnishing", "bathroom", "bath":
		translatedShop = "bathroom"
	case "doityourself":
		translatedShop = "doityourself"
	case "electrical", "electric":
		translatedShop = "electrical"
	case "energy", "battery":
		translatedShop = "energy"
	case "fireplace", "furnace", "oven":
		translatedShop = "fireplace"
	case "florist":
		translatedShop = "florist"
	case "garden_centre":
		translatedShop = "garden centre"
	case "garden_furniture":
		translatedShop = "garden furniture"
	case "gas":
		translatedShop = "gas"
	case "glaziery":
		translatedShop = "glaziery"
	case "hardware":
		translatedShop = "hardware"
	case "houseware", "household", "house":
		translatedShop = "houseware"
	case "locksmith", "keys":
		translatedShop = "locksmith"
	case "paint", "paintings", "colors", "painter":
		translatedShop = "paint"
	case "security":
		translatedShop = "security"
	case "trade":
		translatedShop = "trade"
	// Möbel und Innenausstattung
	case "antiques":
		translatedShop = "antiques"
	case "bed":
		translatedShop = "bed"
	case "candles":
		translatedShop = "candles"
	case "carpet":
		translatedShop = "carpet"
	case "curtain":
		translatedShop = "curtain"
	case "doors":
		translatedShop = "doors"
	case "flooring", "parquet":
		translatedShop = "flooring"
	case "furniture":
		translatedShop = "furniture"
	case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
		translatedShop = "interior decoration"
	case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
		translatedShop = "kitchen"
	case "lamps", "lighting":
		translatedShop = "lamps"
	case "tiles", "tile", "tiling":
		translatedShop = "tiles"
	case "window_blind", "shutter", "shutters":
		translatedShop = "window blind"
	// Elektronik
	case "computer":
		translatedShop = "computer"
	case "robot":
		translatedShop = "robot"
	case "electronics", "electronics_repair", "electro":
		translatedShop = "electronics"
	case "hifi":
		translatedShop = "hifi"
	case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
		translatedShop = "mobile phone"
	case "radiotechnics":
		translatedShop = "radiotechnics"
	case "vacuum_cleaner":
		translatedShop = "vacuum cleaner"
	// Outdoor und Sport, Fahrzeuge
	case "atv":
		translatedShop = "atv"
	case "bicycle", "bike_repair":
		translatedShop = "bicycle"
	case "boat", "yachts":
		translatedShop = "boat"
	case "car":
		translatedShop = "car"
	case "car_repair":
		translatedShop = "car repair"
	case "car_parts":
		translatedShop = "car parts"
	case "caravan", "caravaning":
		translatedShop = "caravan"
	case "fuel":
		translatedShop = "fuel"
	case "fishing", "fishing_gear":
		translatedShop = "fishing"
	case "free_flying":
		translatedShop = "free flying"
	case "golf":
		translatedShop = "golf"
	case "hunting":
		translatedShop = "hunting"
	case "jetski":
		translatedShop = "jetski"
	case "military_surplus", "military":
		translatedShop = "military"
	case "motorcycle", "motorcycle_repair":
		translatedShop = "motorcycle"
	case "outdoor":
		translatedShop = "outdoor"
	case "scuba_diving":
		translatedShop = "diving"
	case "ski":
		translatedShop = "ski"
	case "snowmobile":
		translatedShop = "snowmobile"
	case "sports", "water_sports", "hobby":
		translatedShop = "sports"
	case "swimming_pool":
		translatedShop = "swimming pool"
	case "trailer", "car_trailer":
		translatedShop = "trailer"
	case "tyres":
		translatedShop = "tyres"
	// Kunst, Musik, Hobbys
	case "art", "arts", "artwork":
		translatedShop = "art"
	case "collector", "coins", "comics", "stamps":
		translatedShop = "collector"
	case "craft":
		translatedShop = "craft"
	case "frame", "picture_frames":
		translatedShop = "frame"
	case "games":
		translatedShop = "games"
	case "model", "modelrailway":
		translatedShop = "model"
	case "music":
		translatedShop = "music"
	case "musical_instrument", "woodwind_repair":
		translatedShop = "musical instrument"
	case "photo", "photo_studio", "photographic_studio", "photographer":
		translatedShop = "photo"
	case "camera":
		translatedShop = "camera"
	case "trophy":
		translatedShop = "trophy"
	case "video":
		translatedShop = "video"
	case "video_games":
		translatedShop = "video games"
	// Schreibwaren, Geschenke, Bücher und Zeitungen
	case "anime", "japan":
		translatedShop = "anime"
	case "books", "book_restoration":
		translatedShop = "books"
	case "gift":
		translatedShop = "gift"
	case "lottery":
		translatedShop = "lottery"
	case "newsagent":
		translatedShop = "newsagent"
	case "stationary":
		translatedShop = "stationary"
	case "ticket", "tickets":
		translatedShop = "ticket"
	// Andere
	case "bookmaker":
		translatedShop = "bookmaker"
	case "cannabis", "growshop":
		translatedShop = "cannabis"
	case "copyshop", "printing", "print_shop", "printer_ink", "ink_cartridges", "printer", "printers", "paper", "printshop", "printery":
		translatedShop = "copyshop"
	case "e-cigarette", "vape":
		translatedShop = "e-cigarette"
	case "funeral_directors":
		translatedShop = "funeral directors"
	case "laundry", "rotary_iron", "ironing", "dry_cleaning":
		translatedShop = "laundry"
	case "party":
		translatedShop = "party"
	case "pawnbroker", "money_lender":
		translatedShop = "pawnbroker"
	case "pet":
		translatedShop = "pet"
	case "pet_grooming", "dog_beauty", "dog_hairdresser":
		translatedShop = "pet grooming"
	case "pest_control":
		translatedShop = "pest control"
	case "pyrotechnics":
		translatedShop = "pyrotechnics"
	case "religion":
		translatedShop = "religion"
	case "storage_rental", "rental":
		translatedShop = "storage rental"
	case "tobacco", "smokers":
		translatedShop = "tobacco"
	case "toys":
		translatedShop = "toys"
	case "travel_agency":
		translatedShop = "travel agency"
	case "vacant":
		translatedShop = "vacant"
	case "weapons", "guns", "arms", "knives", "knife":
		translatedShop = "weapons"
	case "outpost":
		translatedShop = "outpost"
	// Benutzerdefiniert
	case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
		translatedShop = "apiary"
	case "auction_house", "auctioneer", "auction":
		translatedShop = "auction house"
	case "car_accessories", "child_safety_seats":
		translatedShop = "car accessories"
	case "car_service":
		translatedShop = "car service"
	case "caretaker", "building_cleaner":
		translatedShop = "caretaker"
	case "carpenter", "cabinet_maker", "carpentry":
		translatedShop = "carpenter"
	case "casino", "gambling":
		translatedShop = "casino"
	case "catalogue":
		translatedShop = "catalogue"
	case "catering", "catering_supplies":
		translatedShop = "catering"
	case "equestrian", "horse_equipment", "horse":
		translatedShop = "equestrian"
	case "esoteric":
		translatedShop = "esoteric"
	case "estate_agent":
		translatedShop = "estate agent"
	case "event_service":
		translatedShop = "event service"
	case "fanshop":
		translatedShop = "fanshop"
	case "fitness_equipment":
		translatedShop = "fitness equipment"
	case "flour":
		translatedShop = "flour"
	case "glass":
		translatedShop = "glass"
	case "garden_service":
		translatedShop = "garden service"
	case "garden_machinery", "gardening_tools", "lawn_mower":
		translatedShop = "garden machinery"
	case "grill", "bbq":
		translatedShop = "grill"
	case "grinding":
		translatedShop = "grinding"
	case "groundskeeping":
		translatedShop = "groundskeeping"
	case "hat", "hats":
		translatedShop = "hat"
	case "health":
		translatedShop = "health"
	case "heating_system", "heater", "heating":
		translatedShop = "heating system"
	case "hookah", "shisha":
		translatedShop = "hookah"
	case "hypnotism":
		translatedShop = "hypnotism"
	case "internet_service_provider":
		translatedShop = "internet service provider"
	case "joiner":
		translatedShop = "joiner"
	case "kids_furnishing":
		translatedShop = "kids furnishing"
	case "furs":
		translatedShop = "furs"
	case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
		translatedShop = "lettering"
	case "locksmithery", "metalworker", "metalwork", "metalworking":
		translatedShop = "locksmithery"
	case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
		translatedShop = "machinery"
	case "office_supplies", "office", "stationery":
		translatedShop = "office supplies"
	case "pet_food", "fodder":
		translatedShop = "pet food"
	case "plumber", "plumbing_business":
		translatedShop = "plumber"
	case "piercing":
		translatedShop = "piercing"
	case "pottery":
		translatedShop = "pottery"
	case "ship_chandler":
		translatedShop = "ship chandler"
	case "software":
		translatedShop = "software"
	case "solarium", "sunstudio":
		translatedShop = "solarium"
	case "stones", "tombstone", "tombstones", "gravestones":
		translatedShop = "stones"
	case "tanning":
		translatedShop = "tanning"
	case "tools", "screws", "tool_hire":
		translatedShop = "tools"
	case "wedding_gown", "wedding":
		translatedShop = "wedding gown"
	case "whirlpool", "pool":
		translatedShop = "whirlpool"
	case "wood", "timber", "sawmill":
		translatedShop = "wood"
	case "worldshop", "one_world", "afro", "oneworld":
		translatedShop = "worldshop"
	default:
		translatedShop = "shop"
		//fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}
