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
		translatedShop = "panadería"
	case "beverages":
		translatedShop = "bebidas"
	case "brewing_supplies", "brewery":
		translatedShop = "cervecería"
	case "butcher":
		translatedShop = "carnicero"
	case "cheese":
		translatedShop = "queso"
	case "chocolate":
		translatedShop = "chocolate"
	case "coffee", "coffee_roasting", "coffeemaker":
		translatedShop = "café"
	case "confectionery", "sweets", "cookies":
		translatedShop = "confitería"
	case "convenience", "food", "grocery":
		translatedShop = "comodidad"
	case "deli":
		translatedShop = "charcutería"
	case "dairy":
		translatedShop = "lácteos"
	case "farm":
		translatedShop = "granja"
	case "frozen_food":
		translatedShop = "alimentos congelados"
	case "greengrocer", "vegetables":
		translatedShop = "frutería"
	case "health_food", "organic":
		translatedShop = "alimentación sana"
	case "ice_cream":
		translatedShop = "helado"
	case "pasta":
		translatedShop = "pasta"
	case "pastry", "cake":
		translatedShop = "pastelería"
	case "seafood", "fishmonger":
		translatedShop = "marisco"
	case "spices":
		translatedShop = "especias"
	case "tea":
		translatedShop = "té"
	case "wine", "winery":
		translatedShop = "vino"
	case "water":
		translatedShop = "agua"
	// Warenhaus, Kaufhaus, Einkaufszentrum
	case "department_store":
		translatedShop = "grandes almacenes"
	case "general", "country_store":
		translatedShop = "tienda rural"
	case "kiosk":
		translatedShop = "quiosco"
	case "mall":
		translatedShop = "centro comercial"
	case "supermarket":
		translatedShop = "supermercado"
	case "wholesale":
		translatedShop = "mayorista"
	// Kleidung, Schuhe, Accessoires
	case "baby_goods":
		translatedShop = "artículos para bebés"
	case "bag", "suitcases":
		translatedShop = "bolsas y maletas"
	case "boutique", "fashion", "fashion_accessories":
		translatedShop = "tienda"
	case "clothes":
		translatedShop = "ropa"
	case "fabric", "fabrics":
		translatedShop = "tela"
	case "jewelry", "jewellery", "gemstone", "gold", "gold_buyer":
		translatedShop = "joyería"
	case "leather":
		translatedShop = "cuero"
	case "sewing", "sewing_machines":
		translatedShop = "coser"
	case "shoes", "shoe_repair", "shoe_repairlocksmith":
		translatedShop = "zapatos"
	case "tailor":
		translatedShop = "sastre"
	case "watches", "clocks", "watchmaker":
		translatedShop = "relojes"
	case "wool":
		translatedShop = "lana"
	// Discounter, Wohltätigkeit
	case "charity", "second_hand":
		translatedShop = "caridad"
	case "variety_store":
		translatedShop = "tienda de variedades"
	// Gesundheit und Schönheitspflege
	case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
		translatedShop = "cosméticos"
	case "chemist":
		translatedShop = "farmacia"
	case "erotic":
		translatedShop = "erótico"
	case "hairdresser":
		translatedShop = "peluquería"
	case "hairdresser_supply", "combs":
		translatedShop = "suministros de peluquería"
	case "hearing_aids":
		translatedShop = "audífonos"
	case "herbalist":
		translatedShop = "herbolario"
	case "massage":
		translatedShop = "masaje"
	case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
		translatedShop = "suministros médicos"
	case "nutrition_supplements":
		translatedShop = "suplementos nutricionales"
	case "optician", "eyeglasses":
		translatedShop = "óptico"
	case "perfumery":
		translatedShop = "perfumería"
	case "tattoo":
		translatedShop = "tatuaje"
	// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
	case "agrarian", "agricultural":
		translatedShop = "agraria"
	case "appliance":
		translatedShop = "aparato"
	case "bathroom_furnishing", "bathroom", "bath":
		translatedShop = "cuarto de baño"
	case "doityourself":
		translatedShop = "hágalo usted mismo"
	case "electrical", "electric":
		translatedShop = "eléctrico"
	case "energy", "battery":
		translatedShop = "energía"
	case "fireplace", "furnace", "oven":
		translatedShop = "horno y estufa"
	case "florist":
		translatedShop = "floristería"
	case "garden_centre":
		translatedShop = "centro de jardinería"
	case "garden_furniture":
		translatedShop = "muebles de jardín"
	case "gas":
		translatedShop = "gas"
	case "glaziery":
		translatedShop = "vidriería"
	case "hardware":
		translatedShop = "hardware"
	case "houseware", "household", "house":
		translatedShop = "menaje del hogar"
	case "locksmith", "keys":
		translatedShop = "cerrajero"
	case "paint", "paintings", "colors", "painter":
		translatedShop = "pintura"
	case "security":
		translatedShop = "seguridad"
	case "trade":
		translatedShop = "comercio"
	// Möbel und Innenausstattung
	case "antiques":
		translatedShop = "antigüedades"
	case "bed":
		translatedShop = "cama"
	case "candles":
		translatedShop = "velas"
	case "carpet":
		translatedShop = "alfombra"
	case "curtain":
		translatedShop = "cortina"
	case "doors":
		translatedShop = "puertas"
	case "flooring", "parquet":
		translatedShop = "suelos"
	case "furniture":
		translatedShop = "muebles"
	case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
		translatedShop = "decoración interior"
	case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
		translatedShop = "cocina"
	case "lamps", "lighting":
		translatedShop = "lámparas"
	case "tiles", "tile", "tiling":
		translatedShop = "baldosas"
	case "window_blind", "shutter", "shutters":
		translatedShop = "persianas"
	// Elektronik
	case "computer":
		translatedShop = "ordenador"
	case "robot":
		translatedShop = "robot"
	case "electronics", "electronics_repair", "electro":
		translatedShop = "electrónica"
	case "hifi":
		translatedShop = "hifi"
	case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
		translatedShop = "teléfono móvil"
	case "radiotechnics":
		translatedShop = "radiotecnia"
	case "vacuum_cleaner":
		translatedShop = "aspiradora"
	// Outdoor und Sport, Fahrzeuge
	case "atv":
		translatedShop = "atv"
	case "bicycle", "bike_repair":
		translatedShop = "bicicleta"
	case "boat", "yachts":
		translatedShop = "barco"
	case "car":
		translatedShop = "coche"
	case "car_repair":
		translatedShop = "reparación de automóviles"
	case "car_parts":
		translatedShop = "piezas de automóviles"
	case "caravan", "caravaning":
		translatedShop = "caravana"
	case "fuel":
		translatedShop = "combustible"
	case "fishing", "fishing_gear":
		translatedShop = "pesca"
	case "free_flying":
		translatedShop = "vuelo libre"
	case "golf":
		translatedShop = "golf"
	case "hunting":
		translatedShop = "caza"
	case "jetski":
		translatedShop = "moto acuática"
	case "military_surplus", "military":
		translatedShop = "militar"
	case "motorcycle", "motorcycle_repair":
		translatedShop = "motocicleta"
	case "outdoor":
		translatedShop = "exterior"
	case "scuba_diving":
		translatedShop = "buceo"
	case "ski":
		translatedShop = "esquiar"
	case "snowmobile":
		translatedShop = "moto de nieve"
	case "sports", "water_sports", "hobby":
		translatedShop = "deportes"
	case "swimming_pool":
		translatedShop = "piscina"
	case "trailer", "car_trailer":
		translatedShop = "remolque"
	case "tyres":
		translatedShop = "neumáticos"
	// Kunst, Musik, Hobbys
	case "art", "arts", "artwork":
		translatedShop = "arte"
	case "collector", "coins", "comics", "stamps":
		translatedShop = "colector"
	case "craft":
		translatedShop = "artesanía"
	case "frame", "picture_frames":
		translatedShop = "marco"
	case "games":
		translatedShop = "juegos"
	case "model", "modelrailway":
		translatedShop = "modelo"
	case "music":
		translatedShop = "música"
	case "musical_instrument", "woodwind_repair":
		translatedShop = "instrumento musical"
	case "photo", "photo_studio", "photographic_studio", "photographer":
		translatedShop = "foto"
	case "camera":
		translatedShop = "cámara"
	case "trophy":
		translatedShop = "trofeo"
	case "video":
		translatedShop = "vídeo"
	case "video_games":
		translatedShop = "videojuegos"
	// Schreibwaren, Geschenke, Bücher und Zeitungen
	case "anime", "japan":
		translatedShop = "anime"
	case "books", "book_restoration":
		translatedShop = "libros"
	case "gift":
		translatedShop = "regalo"
	case "lottery":
		translatedShop = "lotería"
	case "newsagent":
		translatedShop = "quiosco"
	case "stationary":
		translatedShop = "estacionario"
	case "ticket", "tickets":
		translatedShop = "entradas"
	// Andere
	case "bookmaker":
		translatedShop = "corredor de apuestas"
	case "cannabis", "growshop":
		translatedShop = "cannabis"
	case "copyshop", "printing", "print_shop", "printer_ink", "ink_cartridges", "printer", "printers", "paper", "printshop", "printery":
		translatedShop = "copyshop"
	case "e-cigarette", "vape":
		translatedShop = "cigarrillo electrónico"
	case "funeral_directors":
		translatedShop = "directores de funerarias"
	case "laundry", "rotary_iron", "ironing", "dry_cleaning":
		translatedShop = "lavandería"
	case "party":
		translatedShop = "fiesta"
	case "pawnbroker", "money_lender":
		translatedShop = "prestamista"
	case "pet":
		translatedShop = "mascotas"
	case "pet_grooming", "dog_beauty", "dog_hairdresser":
		translatedShop = "peluquería canina"
	case "pest_control":
		translatedShop = "control de plagas"
	case "pyrotechnics":
		translatedShop = "pirotecnia"
	case "religion":
		translatedShop = "religión"
	case "storage_rental", "rental":
		translatedShop = "alquiler"
	case "tobacco", "smokers":
		translatedShop = "tabaco"
	case "toys":
		translatedShop = "juguetes"
	case "travel_agency":
		translatedShop = "agencia de viajes"
	case "vacant":
		translatedShop = "vacante"
	case "weapons", "guns", "arms", "knives", "knife":
		translatedShop = "armas"
	case "outpost":
		translatedShop = "puesto de avanzada"
	// Benutzerdefiniert
	case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
		translatedShop = "colmenar"
	case "auction_house", "auctioneer", "auction":
		translatedShop = "casa de subastas"
	case "car_accessories", "child_safety_seats":
		translatedShop = "accesorios para coches"
	case "car_service":
		translatedShop = "servicio de coches"
	case "caretaker", "building_cleaner":
		translatedShop = "conserje"
	case "carpenter", "cabinet_maker", "carpentry":
		translatedShop = "carpintero"
	case "casino", "gambling":
		translatedShop = "casino"
	case "catalogue":
		translatedShop = "catálogo"
	case "catering", "catering_supplies":
		translatedShop = "catering"
	case "equestrian", "horse_equipment", "horse":
		translatedShop = "ecuestre"
	case "esoteric":
		translatedShop = "esotérico"
	case "estate_agent":
		translatedShop = "agente inmobiliario"
	case "event_service":
		translatedShop = "servicio de eventos"
	case "fanshop":
		translatedShop = "fanshop"
	case "fitness_equipment":
		translatedShop = "equipos de fitness"
	case "flour":
		translatedShop = "harina"
	case "glass":
		translatedShop = "vidrio"
	case "garden_service":
		translatedShop = "servicio de jardinería"
	case "garden_machinery", "gardening_tools", "lawn_mower":
		translatedShop = "maquinaria de jardinería"
	case "grill", "bbq":
		translatedShop = "parrilla"
	case "grinding":
		translatedShop = "molienda"
	case "groundskeeping":
		translatedShop = "jardinería"
	case "hat", "hats":
		translatedShop = "sombreros"
	case "health":
		translatedShop = "salud"
	case "heating_system", "heater", "heating":
		translatedShop = "sistema de calefacción"
	case "hookah", "shisha":
		translatedShop = "shisha"
	case "hypnotism":
		translatedShop = "hipnotismo"
	case "internet_service_provider":
		translatedShop = "proveedor de servicios de internet"
	case "joiner":
		translatedShop = "carpintero"
	case "kids_furnishing":
		translatedShop = "mobiliario infantil"
	case "furs":
		translatedShop = "pieles"
	case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
		translatedShop = "rotulación"
	case "locksmithery", "metalworker", "metalwork", "metalworking":
		translatedShop = "cerrajería"
	case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
		translatedShop = "maquinaria"
	case "office_supplies", "office", "stationery":
		translatedShop = "material de oficina"
	case "pet_food", "fodder":
		translatedShop = "alimentos para mascotas"
	case "plumber", "plumbing_business":
		translatedShop = "fontanero"
	case "piercing":
		translatedShop = "perforación"
	case "pottery":
		translatedShop = "cerámica"
	case "ship_chandler":
		translatedShop = "proveedor de buques"
	case "software":
		translatedShop = "software"
	case "solarium", "sunstudio":
		translatedShop = "solarium"
	case "stones", "tombstone", "tombstones", "gravestones":
		translatedShop = "piedras"
	case "tanning":
		translatedShop = "bronceado"
	case "tools", "screws", "tool_hire":
		translatedShop = "herramientas"
	case "wedding_gown", "wedding":
		translatedShop = "vestido de novia"
	case "whirlpool", "pool":
		translatedShop = "piscina"
	case "wood", "timber", "sawmill":
		translatedShop = "madera"
	case "worldshop", "one_world", "afro", "oneworld":
		translatedShop = "worldshop"
	default:
		translatedShop = "general"
		//fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}
