package main

import "strings"

func translateShop(shop string) string {
	shopSplit := strings.Split(shop, ";")
	var translatedShop string
	switch shopSplit[0] {
	// Lebensmittel, Getränke
	case "alcohol":
		translatedShop = "alcool"
	case "bakery":
		translatedShop = "boulangerie"
	case "beverages":
		translatedShop = "boissons"
	case "brewing_supplies", "brewery":
		translatedShop = "brasserie"
	case "butcher":
		translatedShop = "boucherie"
	case "cheese":
		translatedShop = "fromage"
	case "chocolate":
		translatedShop = "chocolat"
	case "coffee", "coffee_roasting", "coffeemaker":
		translatedShop = "café"
	case "confectionery", "sweets", "cookies":
		translatedShop = "confiserie"
	case "convenience", "food", "grocery":
		translatedShop = "commodité"
	case "deli":
		translatedShop = "charcuterie"
	case "dairy":
		translatedShop = "produits laitiers"
	case "farm":
		translatedShop = "ferme"
	case "frozen_food":
		translatedShop = "aliments surgelés"
	case "greengrocer", "vegetables":
		translatedShop = "légumes"
	case "health_food", "organic":
		translatedShop = "alimentation saine"
	case "ice_cream":
		translatedShop = "crème glacée"
	case "pasta":
		translatedShop = "pâtes"
	case "pastry", "cake":
		translatedShop = "pâtisserie"
	case "seafood", "fishmonger":
		translatedShop = "fruits de mer"
	case "spices":
		translatedShop = "épices"
	case "tea":
		translatedShop = "thé"
	case "wine", "winery":
		translatedShop = "vin"
	case "water":
		translatedShop = "eau"
	// Warenhaus, Kaufhaus, Einkaufszentrum
	case "department_store":
		translatedShop = "grand magasin"
	case "general", "country_store":
		translatedShop = "magasin de campagne"
	case "kiosk":
		translatedShop = "kiosque"
	case "mall":
		translatedShop = "centre commercial"
	case "supermarket":
		translatedShop = "supermarché"
	case "wholesale":
		translatedShop = "vente en gros"
	// Kleidung, Schuhe, Accessoires
	case "baby_goods":
		translatedShop = "produits pour bébés"
	case "bag", "suitcases":
		translatedShop = "sac"
	case "boutique", "fashion", "fashion_accessories":
		translatedShop = "boutique"
	case "clothes":
		translatedShop = "vêtements"
	case "fabric", "fabrics":
		translatedShop = "tissu"
	case "jewelry", "jewellery", "gemstone", "gold", "gold_buyer":
		translatedShop = "bijoux"
	case "leather":
		translatedShop = "cuir"
	case "sewing", "sewing_machines":
		translatedShop = "couture"
	case "shoes", "shoe_repair", "shoe_repairlocksmith":
		translatedShop = "chaussures"
	case "tailor":
		translatedShop = "tailleur"
	case "watches", "clocks", "watchmaker":
		translatedShop = "montres"
	case "wool":
		translatedShop = "laine"
	// Discounter, Wohltätigkeit
	case "charity", "second_hand":
		translatedShop = "charité"
	case "variety_store":
		translatedShop = "magasin de variétés"
	// Gesundheit und Schönheitspflege
	case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
		translatedShop = "beauté"
	case "chemist":
		translatedShop = "chimiste"
	case "erotic":
		translatedShop = "érotique"
	case "hairdresser":
		translatedShop = "coiffeur"
	case "hairdresser_supply", "combs":
		translatedShop = "fournitures pour coiffeurs"
	case "hearing_aids":
		translatedShop = "les appareils auditifs"
	case "herbalist":
		translatedShop = "herboriste"
	case "massage":
		translatedShop = "massage"
	case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
		translatedShop = "approvisionnement médical"
	case "nutrition_supplements":
		translatedShop = "les compléments alimentaires"
	case "optician", "eyeglasses":
		translatedShop = "opticien"
	case "perfumery":
		translatedShop = "parfumerie"
	case "tattoo":
		translatedShop = "tatouage"
	// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
	case "agrarian", "agricultural":
		translatedShop = "agraire"
	case "appliance":
		translatedShop = "appareil ménager"
	case "bathroom_furnishing", "bathroom", "bath":
		translatedShop = "salle de bains"
	case "doityourself":
		translatedShop = "à faire soi-même"
	case "electrical", "electric":
		translatedShop = "électrique"
	case "energy", "battery":
		translatedShop = "énergie"
	case "fireplace", "furnace", "oven":
		translatedShop = "cheminée"
	case "florist":
		translatedShop = "fleuriste"
	case "garden_centre":
		translatedShop = "centre de jardinage"
	case "garden_furniture":
		translatedShop = "meubles de jardin"
	case "gas":
		translatedShop = "gaz"
	case "glaziery":
		translatedShop = "vitrerie"
	case "hardware":
		translatedShop = "matériel informatique"
	case "houseware", "household", "house":
		translatedShop = "articles ménagers"
	case "locksmith", "keys":
		translatedShop = "serrurier"
	case "paint", "paintings", "colors", "painter":
		translatedShop = "peinture"
	case "security":
		translatedShop = "sécurité"
	case "trade":
		translatedShop = "commerce"
	// Möbel und Innenausstattung
	case "antiques":
		translatedShop = "antiquités"
	case "bed":
		translatedShop = "lit"
	case "candles":
		translatedShop = "bougies"
	case "carpet":
		translatedShop = "tapis"
	case "curtain":
		translatedShop = "rideau"
	case "doors":
		translatedShop = "portes"
	case "flooring", "parquet":
		translatedShop = "revêtement de sol"
	case "furniture":
		translatedShop = "meubles"
	case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
		translatedShop = "décoration intérieure"
	case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
		translatedShop = "cuisine"
	case "lamps", "lighting":
		translatedShop = "lampes"
	case "tiles", "tile", "tiling":
		translatedShop = "tuiles"
	case "window_blind", "shutter", "shutters":
		translatedShop = "store de fenêtre"
	// Elektronik
	case "computer":
		translatedShop = "ordinateur"
	case "robot":
		translatedShop = "robot"
	case "electronics", "electronics_repair", "electro":
		translatedShop = "électronique"
	case "hifi":
		translatedShop = "hifi"
	case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
		translatedShop = "téléphone portable"
	case "radiotechnics":
		translatedShop = "radiotechnique"
	case "vacuum_cleaner":
		translatedShop = "aspirateur"
	// Outdoor und Sport, Fahrzeuge
	case "atv":
		translatedShop = "atv"
	case "bicycle", "bike_repair":
		translatedShop = "vélo"
	case "boat", "yachts":
		translatedShop = "bateau"
	case "car":
		translatedShop = "voiture"
	case "car_repair":
		translatedShop = "réparation de voitures"
	case "car_parts":
		translatedShop = "pièces de voitures"
	case "caravan", "caravaning":
		translatedShop = "caravane"
	case "fuel":
		translatedShop = "carburant"
	case "fishing", "fishing_gear":
		translatedShop = "pêche"
	case "free_flying":
		translatedShop = "vol libre"
	case "golf":
		translatedShop = "golf"
	case "hunting":
		translatedShop = "chasse"
	case "jetski":
		translatedShop = "jetski"
	case "military_surplus", "military":
		translatedShop = "militaire"
	case "motorcycle", "motorcycle_repair":
		translatedShop = "moto"
	case "outdoor":
		translatedShop = "extérieur"
	case "scuba_diving":
		translatedShop = "plongée"
	case "ski":
		translatedShop = "ski"
	case "snowmobile":
		translatedShop = "motoneige"
	case "sports", "water_sports", "hobby":
		translatedShop = "sports"
	case "swimming_pool":
		translatedShop = "piscine"
	case "trailer", "car_trailer":
		translatedShop = "remorque"
	case "tyres":
		translatedShop = "pneus"
	// Kunst, Musik, Hobbys
	case "art", "arts", "artwork":
		translatedShop = "art"
	case "collector", "coins", "comics", "stamps":
		translatedShop = "collecteur"
	case "craft":
		translatedShop = "artisanat"
	case "frame", "picture_frames":
		translatedShop = "cadre"
	case "games":
		translatedShop = "jeux"
	case "model", "modelrailway":
		translatedShop = "modèle"
	case "music":
		translatedShop = "musique"
	case "musical_instrument", "woodwind_repair":
		translatedShop = "instrument de musique"
	case "photo", "photo_studio", "photographic_studio", "photographer":
		translatedShop = "photo"
	case "camera":
		translatedShop = "caméra"
	case "trophy":
		translatedShop = "trophée"
	case "video":
		translatedShop = "vidéo"
	case "video_games":
		translatedShop = "jeux vidéo"
	// Schreibwaren, Geschenke, Bücher und Zeitungen
	case "anime", "japan":
		translatedShop = "anime"
	case "books", "book_restoration":
		translatedShop = "livres"
	case "gift":
		translatedShop = "cadeau"
	case "lottery":
		translatedShop = "loterie"
	case "newsagent":
		translatedShop = "marchand de journaux"
	case "stationary":
		translatedShop = "stationnaire"
	case "ticket", "tickets":
		translatedShop = "billet"
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
		translatedShop = "directeurs de funérailles"
	case "laundry", "rotary_iron", "ironing", "dry_cleaning":
		translatedShop = "blanchisserie"
	case "party":
		translatedShop = "fête"
	case "pawnbroker", "money_lender":
		translatedShop = "prêteur sur gages"
	case "pet":
		translatedShop = "animal de compagnie"
	case "pet_grooming", "dog_beauty", "dog_hairdresser":
		translatedShop = "toilettage des animaux"
	case "pest_control":
		translatedShop = "contrôle des nuisibles"
	case "pyrotechnics":
		translatedShop = "pyrotechnie"
	case "religion":
		translatedShop = "religion"
	case "storage_rental", "rental":
		translatedShop = "location de stockage"
	case "tobacco", "smokers":
		translatedShop = "tabac"
	case "toys":
		translatedShop = "jouets"
	case "travel_agency":
		translatedShop = "agence de voyage"
	case "vacant":
		translatedShop = "vacant"
	case "weapons", "guns", "arms", "knives", "knife":
		translatedShop = "armes"
	case "outpost":
		translatedShop = "avant-poste"
	// Benutzerdefiniert
	case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
		translatedShop = "rucher"
	case "auction_house", "auctioneer", "auction":
		translatedShop = "maison d'enchères"
	case "car_accessories", "child_safety_seats":
		translatedShop = "accessoires automobiles"
	case "car_service":
		translatedShop = "service automobile"
	case "caretaker", "building_cleaner":
		translatedShop = "concierge"
	case "carpenter", "cabinet_maker", "carpentry":
		translatedShop = "charpentier"
	case "casino", "gambling":
		translatedShop = "casino"
	case "catalogue":
		translatedShop = "catalogue"
	case "catering", "catering_supplies":
		translatedShop = "restauration"
	case "equestrian", "horse_equipment", "horse":
		translatedShop = "équestre"
	case "esoteric":
		translatedShop = "ésotérique"
	case "estate_agent":
		translatedShop = "agent immobilier"
	case "event_service":
		translatedShop = "service événementiel"
	case "fanshop":
		translatedShop = "boutique de fans"
	case "fitness_equipment":
		translatedShop = "matériel de fitness"
	case "flour":
		translatedShop = "farine"
	case "glass":
		translatedShop = "verre"
	case "garden_service":
		translatedShop = "service de jardinage"
	case "garden_machinery", "gardening_tools", "lawn_mower":
		translatedShop = "machines de jardinage"
	case "grill", "bbq":
		translatedShop = "gril"
	case "grinding":
		translatedShop = "broyage"
	case "groundskeeping":
		translatedShop = "entretien des terrains"
	case "hat", "hats":
		translatedShop = "chapeau"
	case "health":
		translatedShop = "santé"
	case "heating_system", "heater", "heating":
		translatedShop = "système de chauffage"
	case "hookah", "shisha":
		translatedShop = "narguilé"
	case "hypnotism":
		translatedShop = "hypnotisme"
	case "internet_service_provider":
		translatedShop = "fournisseur de services internet"
	case "joiner":
		translatedShop = "menuisier"
	case "kids_furnishing":
		translatedShop = "mobilier pour enfants"
	case "furs":
		translatedShop = "fourrures"
	case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
		translatedShop = "lettrage"
	case "locksmithery", "metalworker", "metalwork", "metalworking":
		translatedShop = "serrurerie"
	case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
		translatedShop = "machines"
	case "office_supplies", "office", "stationery":
		translatedShop = "fournitures de bureau"
	case "pet_food", "fodder":
		translatedShop = "aliments pour animaux"
	case "plumber", "plumbing_business":
		translatedShop = "plombier"
	case "piercing":
		translatedShop = "piercing"
	case "pottery":
		translatedShop = "poterie"
	case "ship_chandler":
		translatedShop = "manutentionnaire de navires"
	case "software":
		translatedShop = "logiciel"
	case "solarium", "sunstudio":
		translatedShop = "solarium"
	case "stones", "tombstone", "tombstones", "gravestones":
		translatedShop = "pierres"
	case "tanning":
		translatedShop = "bronzage"
	case "tools", "screws", "tool_hire":
		translatedShop = "outils"
	case "wedding_gown", "wedding":
		translatedShop = "robe de mariée"
	case "whirlpool", "pool":
		translatedShop = "tourbillon"
	case "wood", "timber", "sawmill":
		translatedShop = "bois"
	case "worldshop", "one_world", "afro", "oneworld":
		translatedShop = "boutique en ligne"
	default:
		translatedShop = "shop"
		//fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}
