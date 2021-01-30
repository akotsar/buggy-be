package seed

import "time"

// MakeSeed represents a seeded make.
type MakeSeed struct {
	Name        string
	Image       string
	Description string
	Models      []ModelSeed
}

// ModelSeed represents a seeded model.
type ModelSeed struct {
	Name        string
	Image       string
	Description string
	EngineVol   float64
	MaxSpeed    int
	Votes       int
	Comments    []CommentSeed
}

// CommentSeed represents a seeded comment.
type CommentSeed struct {
	Comment    string
	DatePosted time.Time
}

// GetMakes return seed makes data
func GetMakes() []MakeSeed {
	return []MakeSeed{
		{
			Name:        "Lamborghini",
			Image:       "Lamborghini-Logo.png",
			Description: "**Automobili Lamborghini S.p.A.** is an Italian brand and manufacturer of luxury sports cars and SUVs based in Sant'Agata Bolognese, Italy. The company is owned by the Volkswagen Group through its subsidiary Audi. In 2015, Lamborghini's 1,175 employees produced 3,248 vehicles.\n\nFerruccio Lamborghini, an Italian manufacturing magnate, founded Automobili Ferruccio Lamborghini S.p.A. in 1963 to compete with established marques, including Ferrari. The company gained wide acclaim in 1966 for the Miura sports coupé, which established rear mid-engine, rear wheel drive as the standard layout for high-performance cars of the era. Lamborghini grew rapidly during its first decade, but sales plunged in the wake of the 1973 worldwide financial downturn and the oil crisis. The firm's ownership changed three times after 1973, including a bankruptcy in 1978. American Chrysler Corporation took control of Lamborghini in 1987 and sold it to Malaysian investment group Mycom Setdco and Indonesian group V'Power Corporation in 1994. In 1998, Mycom Setdco and V'Power sold Lamborghini to the Volkswagen Group where it was placed under the control of the group's Audi division.\n\nLamborghini produces sports cars and V12 engines for offshore powerboat racing. Lamborghini currently produces the V12-powered Aventador and the V10-powered Huracán.",
			Models: []ModelSeed{
				{
					Name:        "HURACÁN",
					Image:       "Huracan.jpg",
					Description: "Lamborghini’s Super Trofeo championship has been described as the fastest one-make race series in the world. From 2015, it’s set to get even quicker. Racing in a separate class to the outgoing Gallardo, the new rear-drive **Huracán LP 620-2 Super Trofeo** features less weight, greater rigidity and more power.\n\nThe single most striking difference between Huracán road and race car is that the latter is rear-drive. Lamborghini argues this will make the Super Trofeo series more attractive to drivers seeking to progress up to GT3. It also cuts weight, while the traction advantage of all-wheel drive is less significant on track and with racing tyres. The chassis is modified to accept a larger race radiator up front and the racing gearbox at the rear, while the bodywork itself is mounted on fast fittings for pit-lane repairs. Aerodynamics have been tuned for greater downforce and drag efficiency; the rear wing has ten different downforce settings while the front air intake ducts are also adjustable.\n\nPower is provided by a 620 hp version of Lamborghini’s V10, driving through a three-disc clutch and XTrac sequential transmission. Power to weight ratio is an impressive 2.05 kilograms per horsepower (compared with 2.33 kg/hp for the roadgoing all-wheel drive model).\n\nInside the cockpit, luxuries are few but the sophisticated MOTEC M182 electronic control unit governs data, display and gearchanges while the Bosch Motorsport ABS offers 12 different traction settings that the driver can select via a control on the steering wheel. Adjustable anti-roll bars and dampers will enable drivers to set up the LP 620-2 Super Trofeo to their precise requirements, whatever the circuit or conditions.",
					EngineVol:   5.2,
					MaxSpeed:    325,
					Votes:       123,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Hear the ROARRR",
						},
					},
				},
				{
					Name:        "AVENTADOR",
					Image:       "aventador.jpg",
					Description: "A car like the **Lamborghini Aventador** is awesome enough in its own right that it really doesn’t need an aftermarket program. Still, there are those who prefer to give it a tuning program just to see how much faster and more powerful the car can go. That’s perfectly understandable. But every so often, you chance upon a car that is the recipient of not one, but two aftermarket programs.\n\nThis Lamborghini Aventador is the perfect example of that. By and large, this is an Aventador, albeit a special edition one that was released in limited quantities by Oakley Design back in 2012. It was nicknamed the “Dragon Edition” and only 10 units were made, each coming with an aero kit made up of a fixed carbon fiber rear spoiler, redesigned front bumper, and rear air vents. It also had individually-numbered plaques and embroidered dragons on the doors, and most importantly, a massive performance bump that gave the Aventador’s 6.5-liter V-12 engine 760 horsepower 550 pound-feet of torque to play with.",
					EngineVol:   6.5,
					MaxSpeed:    350,
					Votes:       543,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Just Aventador. No more words needed.",
						},
					},
				},
				{
					Name:        "GALLARDO",
					Image:       "Gallardo.jpg",
					Description: "Soon to be replaced by the Huracán, this baby bull still turns heads. The **Gallardo**’s sharply creased styling and V-10 prove irresistible to drivers who want to be both seen and heard. \nSpyder and coupe variants include the rear-wheel-drive LP550-2 and all-wheel-drive LP560-4. \nOnly the rear-drive coupe offers a six-speed manual; others feature a six-speed automated manual. \n\nSeveral special editions, each with an Italian name like Squadra Corse or Superleggera Edizione Tecnica, are available.",
					EngineVol:   5.2,
					MaxSpeed:    324,
					Votes:       353,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Gallardo...",
						},
					},
				},
				{
					Name:        "Veneno",
					Image:       "veneno.jpg",
					Description: "Lamborghini is evolving its styling language, and it’s more evident than ever in the **Veneno** shown at the Geneva auto show. Based on the Aventador LP700-4, it will be built in exactly three units, plus the company's demonstrator car. What is the reason for showing another supercar, given that Lamborghini has not yet delivered its ultra-low-volume Sesto Elemento to customers? It's the company's 50th birthday, which it celebrates in May. And the Veneno—named after \"one of the strongest and most aggressive fighting bulls ever,\" as Lamborghini informs us—presents the perfect way to celebrate.\n\nThe fissured skin of the Veneno hides the Aventador's carbon-fiber monocoque, plus aluminum front and rear subframes. A pushrod suspension with horizontal spring-damper units betrays its racing aspirations. The interior is largely carried over from the Aventador and is clad in carbon fiber. The Veneno is fitted with Pirelli P Zero tires on 20-inch wheels up front and 22-inch wheels in the rear. Center-locking hubs allow for quicker changes—and they look great.",
					EngineVol:   6.5,
					MaxSpeed:    355,
					Votes:       543,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "An absolute beauty",
						},
					},
				},
				{
					Name:        "Reventón",
					Image:       "reventon.jpg",
					Description: "If the greatest dream in all gearheaddom is of greater and greater speed, then the fullest realization of that dream must trade wheels for wings. If you can watch Top Gun without cursing your slow reflexes, bad back, or coke-bottle glasses, then maybe you should stick to collecting stamps.\n\nSo Lamborghini, builder of dreams and fulfiller of fantasies, has decided if its customers can't fly Maverick's F-14 Tomcat, then at least they can drive something that looks like the ground-bound equivalent.\n\nIf you're thinking you've seen this before, that's because what looks like a Murciélago in costume is in fact a Murciélago in costume. According to Stephan Winkelmann, president and CEO of Automobili Lamborghini, the company \"took the technical base of the Murciélago LP640 and compressed and intensified its DNA, its genetic code.\" In other words, Lamborghini took the already over-the-top Murciélago and went so high above it that Luke Skywalker could mistake the **Reventón** for an enemy combatant.",
					EngineVol:   6.5,
					MaxSpeed:    330,
					Votes:       543,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "No more words needed",
						},
					},
				},
				{
					Name:        "Murciélago",
					Image:       "Murcielago.jpg",
					Description: "The **Murciélago** has been in Lamborghini's 16 U.S. dealerships since December 2001. Since then, 200 have been sold. When a dealer places an order, the car is air-freighted from Italy in a sealed container and can be disrupting American schoolchildren in as few as 10 days. There are two options only: pearlescent paint ($2500) and a nav system ($3500). \n\nLamborghinis are famous for being as fragile as spring ice, so it was of some concern that our test car showed 15,500 miles-as much mule as bull. \"It's survived 35 road tests,\" asserted Lamborghini tech adviser Ken McCay, who is not Italian. \"You're the only guys who broke it [last summer, when a universal joint pulled free in the shift linkage].\" We broke it this time, too, when the whole shift lever snapped off at the root. The car came with the license \"AL 147,\" a reference to \"Automobili Lamborghini engineering project No. 147.\" Maybe project No. 148 will be devoted to shift-linkage reinforcement. What you notice first about the Murciélago is that its left-front wheel intrudes some eight inches into prime footwell territory, skewing your feet to the right. Your left foot searches for a place to relax-under the clutch is about the only comfortable spot.\n\nWhat you notice next is the accelerator pedal juts out of a small black box, like a paddle raised in a canoe. Your heel rests on the front of this box, and you bend your toes forward to move the throttle. You can duplicate the sensation by walking around with a box of Tic Tacs in your shoe.",
					EngineVol:   6.5,
					MaxSpeed:    342,
					Votes:       543,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "What can be better?!",
						},
					},
				},
				{
					Name:        "Diablo",
					Image:       "Diablo.jpg",
					Description: "This last **Diablo** in a decade of flamboyant—some daresay outrageous—sports cars first introduced in 1990 is certainly the best Diablo.\n\nThe VT is, curiously, a very high-powered four-wheel-drive beast. Situated amidships is the biggest, baddest V-12 engine we've ever tested in a Diablo. For this 2000 model, Lamborghini has lengthened the stroke by 0.16 inch, lightened the crankshaft, used lighter and stronger titanium connecting rods, and updated the old 16-bit engine-control system to a more powerful 32-bit unit. As a result, peak engine output has been promoted to 543 horsepower at 7100 rpm, 20 more than found in the last Diablo. Torque is up 11 pound-feet to 457 at 5800 rpm. \n\nHorsepower freaks take note: *This latest Diablo now has more horsepower than four four-cylinder Toyota Camrys.*",
					EngineVol:   6,
					MaxSpeed:    25,
					Votes:       999,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "The best car in the world!",
						},
					},
				},
				{
					Name:        "Countach",
					Image:       "countach.jpg",
					Description: "The **Countach** made its public debut at the 1971 Geneva motor show. The design of the ultra low two seater sports car took the world by surprise. Its most captivating parts were of course the scissor doors, swinging up and forward. Over the years these famous doors have become Lamborghini’s trade mark right up to the latest Murcielago. \n\nThis particular Countach is powered by a V-12 engine with an output of 375 HP and mated to a 5-speed manual transmission. The model got its \"Periscopica\" name from the periscope mounted on the roof.",
					EngineVol:   5.17,
					MaxSpeed:    315,
					Votes:       444,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Just wait till you see Diablo!",
						},
					},
				},
				{
					Name:        "Miura",
					Image:       "Miura.jpg",
					Description: "The **Miura** is Lamborghini's latest bid to out-Ferrari Ferrari. For starters, \"Miura\" is a breed of Spanish fighting bull, and the wild Miura, with its 430-horsepower, V-12 engine mounted transversely in the rear, is so bold, individualistic and unconventional that it's hard to imagine it fitting into anybody's arbitrary standards, safety or otherwise. \n\nDifficult though the challenge may seem, we believe the normal Lamborghini 350 GT (C/D, March 1966) and perhaps even the Miura will be modified to conform to any standards the National Traffic Safety Bureau will impose. That's the kind of commercial competitor Ferruccio Lamborghini is. As a self-made man—a successful manufacturer of tractors, heating units, and air conditioners—Lamborghini's sole objective in the automobile business (expense be damned) is to surpass Ferrari, first in sales, and then in reputation. He'll do what he needs to do to sell cars in the rich American market. If his two American importer-distributors agree that there should be different headlights or an optional automatic transmission, they let Lamborghini know, and behold—the necessary modifications are made with a rapidity that would astound those familiar with the normal reaction times of Italian auto‑makers, especially the smaller firms.\n\nThere are only two other cars that can be mentioned in the same breath with the Miura. They are the Ford Mark III (C/D, June 1967) and the Ferrari 275/LM (C/D, May 1965). Both are available in road versions, but were originally built for the track, where they will always be more at home. The Miura, on the other hand, was engineered with racing as a future possibility, but was developed solely for the street. As such, it is superior in performance, space and furnishings—in everything, in short—to the other twenty-grand two-seaters.",
					EngineVol:   3.9,
					MaxSpeed:    280,
					Votes:       444,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "A nice-looking old lady who knows how to speed.",
						},
					},
				},
			},
		},
		{
			Name:        "Alfa Romeo",
			Image:       "AR-logo.jpg",
			Description: "The company that became **Alfa Romeo** was founded as Società Anonima Italiana Darracq (SAID) in 1906 by the French automobile firm of Alexandre Darracq, with some Italian investors. In late 1909, the Italian Darracq cars were selling slowly and the Italian partners of the company hired Giuseppe Merosi to design new cars. On June 24, 1910, a new company was founded named A.L.F.A., initially still in partnership with Darracq. The first non-Darracq car produced by the company was the 1910 24 HP, designed by Merosi. A.L.F.A. ventured into motor racing, with drivers Franchini and Ronzoni competing in the 1911 Targa Florio with two 24-hp models. In August 1915, the company came under the direction of Neapolitan entrepreneur Nicola Romeo, who converted the factory to produce military hardware for the Italian and Allied war efforts. In 1920, the name of the company was changed to Alfa Romeo with the Torpedo 20-30 HP the first car to be so badged.\n\nAlfa Romeo has competed successfully in many different categories of motorsport, including Grand Prix motor racing, Formula One, sportscar racing, touring car racing, and rallies. It has competed both as a constructor and an engine supplier, via works entries (usually under the name Alfa Corse or Autodelta), and private entries. The first racing car was made in 1913, three years after the foundation of the company, and Alfa Romeo won the inaugural world championship for Grand Prix cars in 1925. The company gained a good name in motorsport, which gave a sporty image to the whole marque. Enzo Ferrari founded the Scuderia Ferrari racing team in 1929 as an Alfa Romeo racing team, before becoming independent in 1939. It holds the world's title of the most wins of any marque in the world.",
			Models: []ModelSeed{
				{
					Name:        "Giulietta",
					Image:       "giulietta.jpg",
					Description: "Derived from racing experience, Alfa D.N.A. is the exclusive Alfa Romeo driving selector which, by acting on engine, brakes, steering, gearbox, suspension and accelerator, perfectly adapts the vehicle's performance to suit the driver's style and the road conditions.\n\nIt offers simple, intuitive use:\n * In sporty Dynamic mode, the electronic traction control systems cut in more discreetly, the engine and brakes are more prompt and reactive while the steering is more direct and sporty.\n * When the selector is moved to Natural urban position, the setting is more neutral and fuel consumption is optimized.\n * When road surface grip is limited, All Weather mode may be selected. This has been designed to make all the electronic safety devices cut in at an earlier stage.\n\nThe Giulietta features an all turbo petrol engine lineup. The 1.4 Turbo 88kW / 215Nm engine in the Progression is the ideal engine for those in search of a car that can handle city traffic effortlessly while keeping running costs to a minimum. The turbo engine ensures a prompt response even at low speeds while the Start&Stop system makes it possible to cut consumption and harmful emission levels drastically, without compromising comfort and on-board safety. While the 1.4 Turbo MultiAir 125kW / 250Nm engine in the Distinctive offers the ultimate in terms of performance, fuel consumption and emissions.",
					EngineVol:   1.74,
					MaxSpeed:    242,
					Votes:       500,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Such a cutie!",
						},
						{
							DatePosted: time.Date(2016, 07, 11, 11, 0, 0, 0, time.UTC),
							Comment:    "A very city-friendly car!",
						},
					},
				},
				{
					Name:        "4c Spider",
					Image:       "spider.jpg",
					Description: "The Alfa Romeo **4C Spider** is built with the same materials you’ll find in a Formula 1 car. The 4C Spider’s unique alloy wheels are milled from just one piece of aluminium. And finally, SMC, an innovative low-density composite material is used for the outer body. The result is 0-100km/h in just 4.5 seconds and onto a top speed of over 258km/h.\n\nThe 4C's 1750 cc turbocharged power plant is based on state-of-the-art engine technology. 4 cylinders, an aluminium block, a new generation turbocharger, ultra-high pressure direct fuel injection, two continuously variable valve timing units, scavenging technology and a dual clutch enable this 1750 cc unit to deliver unprecedented sporting performance. The Alfa Romeo 4C reaches 100 km/h in 4.5 seconds and develops a top speed of 258 km/h. Despite this ballistic performance, the emissions of the Alfa Romeo 4C fall well within the strict limits of Euro 6.",
					EngineVol:   1.75,
					MaxSpeed:    260,
					Votes:       333,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "I'd love to have one!",
						},
						{
							DatePosted: time.Date(2016, 07, 11, 11, 0, 0, 0, time.UTC),
							Comment:    "Worth its money definitely",
						},
					},
				},
				{
					Name:        "Mito",
					Image:       "mito.jpg",
					Description: "The style of the *Mito* lives up to the latest news from the Alfa Romeo family in design, onboard experience and in its performance too.\n\nIn this way, the smallest Alfa Romeo sporty car hasn’t just grown-up: it has become grandiose.\n\nStrength and agility are Mito’s strong points. To provide them, Alfa Romeo engineers worked hard on every element of the car, to ensure only the best performance.",
					EngineVol:   1.4,
					MaxSpeed:    219,
					Votes:       500,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Such a cutie!",
						},
						{
							DatePosted: time.Date(2016, 07, 11, 11, 0, 0, 0, time.UTC),
							Comment:    "A very city-friendly car!",
						},
					},
				},
				{
					Name:        "Guilia Quadrifoglio",
					Image:       "giulia.png",
					Description: "New Giulia Quadrifoglio represents more than the most powerful Alfa Romeo ever created for street use. It represents a convergence of engineering and emotion that can only belong to a brand as fabled as Alfa Romeo.\n\nHere's to a sports badge born 106 years ago that still stands for something totally original today: a passion for motoring unlike any other. Visceral. Energetic. Technological. Crafted.\n\nThe technology behind New Giulia Quadrifoglio is created to enhance performance and to give great driving sensations. The human/machine relationship is always at the core of all innovation to create a car that is a natural extension of the body, the mind and the heart of the driver.",
					EngineVol:   2.9,
					MaxSpeed:    307,
					Votes:       700,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 12, 11, 0, 0, 0, time.UTC),
							Comment:    "Lovely!",
						},
					},
				},
			},
		},
		{
			Name:        "Pagani",
			Image:       "Pagani-Logo.jpg",
			Description: "**Pagani Automobili S.p.A.** is an Italian manufacturer of Supercars and carbon fibre. The company was founded in 1992 by the Argentinian Horacio Pagani, and is based in San Cesario sul Panaro, near Modena, Italy.",
			Models: []ModelSeed{
				{
					Name:        "Zonda",
					Image:       "zonda.jpg",
					Description: "The evolution of the species, the revolution in the concept of art applied to pure speed. \n\nThe **Pagani Zonda Revolucion** is the apex of the celebration of performance, technology and art applied to a track car. Horacio Pagani and his team have created a car designed to amaze both on the track and in a car collection.\n\nThe central monocoque is carbon-titanium, the needle on the scale stops at 1070kg. The AMG Mercedes engine is an evolution of the Zonda R powerplant. The 6.0-liter V12 now develops an output of 800 hp and 730 Nm of torque, resulting in a power to weight ratio of 748 hp per tonne.\n\nThe 6 speed magnesium transversal and sequential gearbox changes gears in 20ms. The traction control developed by Bosch with 12 different settings and the renewed ABS system, allows the driver to adapt the behavior of the car to his driving style.",
					EngineVol:   6,
					MaxSpeed:    350,
					Votes:       777,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 18, 11, 0, 0, 0, time.UTC),
							Comment:    "The star",
						},
					},
				},
				{
					Name:        "Huayra",
					Image:       "Huayra.jpg",
					Description: "The **Huayra** is made of more than 4000 components (engine and gearbox not included). To create them and put them together requires creativity, patience and passion that I shared with a fantastic young team and with the most competent partner in all sectors.\n\nIn defining the size, I immediately thought of a car that would be longer than the Zonda, a track increased by 70 mm, a cabin position shifted 40 mm to the back and even more spacious. The silhouette should be soft, easy to read and form itself from lean and sleek lines that have a clear beginning and an end.\n\nMercedes-AMG has created a truly unique and lightweight engine, a twin turbo with 730 HP and 1000 Nm of torque that perfectly complements the car giving a feeling that has motivated our research: that of the brute force of an airplane taking off.\n\nThe main focus was in a power delivery that was linear and not with lag, because that could potentially create safety problems or a continuous operation of electronic driving aids. \n\nMercedes-AMG already 5 years ago thought about the coming environmental restrictions by creating a 12-cylinder engine that is at the peak of efficiency in terms of CO2 and consumption.",
					EngineVol:   5.98,
					MaxSpeed:    383,
					Votes:       553,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 18, 12, 0, 0, 0, time.UTC),
							Comment:    "Lovely, isn't she?",
						},
					},
				},
			},
		},
		{
			Name:        "Bugatti",
			Image:       "bugatti-logo.png",
			Description: "**Automobiles Ettore Bugatti** was a French car manufacturer of high-performance automobiles, founded in 1909 in the then German city of Molsheim, Alsace by Italian-born Ettore Bugatti. Bugatti cars were known for their design beauty (Ettore Bugatti was from a family of artists and considered himself to be both an artist and constructor) and for their many race victories. Famous Bugattis include the Type 35 Grand Prix cars, the Type 41 \"Royale\", the Type 57 \"Atlantic\" and the Type 55 sports car.\n\nBugatti Automobiles S.A.S. is a French high-performance luxury automobiles manufacturer and a subsidiary of Volkswagen AG, with its head office and assembly plant in Molsheim, Alsace, France. Volkswagen purchased the Bugatti trademark in June 1998 and incorporated Bugatti Automobiles S.A.S. in 1999.\n\nBugatti presented several concept cars between 1998 and 2000 before commencing development of its first production model, the Veyron 16.4, delivering the first Veyron to a customer in 2005.",
			Models: []ModelSeed{
				{
					Name:        "Veyron",
					Image:       "veyron.jpg",
					Description: "Since its launch in 2005, the Bugatti **Veyron** has been regarded as a supercar of superlative quality. It was a real challenge for developers to fulfil the specifications that the new supercar was supposed to meet: over 1,000 hp, a top speed of over 400 km/h and the ability to accelerate from 0 to 100 in under three seconds. \n\n Even experts thought it was impossible to achieve these performance specs on the road. But that was not all. The same year, the Super Sport fulfilled the strict requirements of Guinness World Records and set a new world speed record for road cars of 431.072 km/h. Despite numerous attempts to dethrone the Super Sport from its status as the fastest production supercar, the Bugatti remains unbeaten to this day.",
					EngineVol:   8,
					MaxSpeed:    431,
					Votes:       234,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "A very good car.",
						},
						{
							DatePosted: time.Date(2016, 07, 11, 11, 0, 0, 0, time.UTC),
							Comment:    "Love it!",
						},
					},
				},
				{
					Name:        "Chiron",
					Image:       "chiron.jpg",
					Description: "The **Chiron** is the most modern interpretation of Bugatti’s brand DNA and embodies our new design language. The styling accentuates the performance aspect of the super sports car. The motto adopted by the Bugatti designers for the Chiron was “Form follows Performance”. Inspired through Bugatti Type 57SC Atlantic the new design language is characterised by extremely generous surfaces, which are demarcated by pronounced lines in the case of the Chiron. Thereby most of these elements have a technical background and have been designed to fully accentuate the growing performance requirements of the Chiron.\n\nIn order to achieve a 25 percent increase in performance compared to its predecessor, almost every single part of the engine was looked at and newly developed. This feat of engineering resulted in the W16 engine of the Bugatti Chiron being able to develop an unbelievable 1,103 kW (1,500 bps) from its 8 litres of displacement. The engine reaches its maximum torque of 1,600 Nm thanks to the turbocharger which Bugatti actually even developed itself. The 4 turbochargers are now double-powered and already guarantee maximum torque at 2,000 rpm, and the torque is maintained at this level all the way up to 6,000 rpm. The result is unbelievable acceleration which only comes to an end in the twilight zone somewhere beyond the 400 km/h mark.",
					EngineVol:   7.99,
					MaxSpeed:    420,
					Votes:       150,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "Looks amazing.",
						},
						{
							DatePosted: time.Date(2016, 07, 11, 11, 0, 0, 0, time.UTC),
							Comment:    "Mmmm... Just hear the beast roar!",
						},
					},
				},
			},
		},
		{
			Name:        "Lancia",
			Image:       "lancia-logo.jpg",
			Description: "**Lancia** is an Italian automobile manufacturer founded in 1906 by Vincenzo Lancia as Lancia & C.. It became part of the Fiat Group in 1969; the current company, Lancia Automobiles S.p.A., was established in 2007.\n\nThe company has a strong rally heritage and is noted for using letters of the Greek alphabet for its model names.\nLancia vehicles are no longer sold outside of Italy, and comprise only the Ypsilon supermini range, as Fiat CEO Sergio Marchionne foreshadowed in January 2014",
			Models: []ModelSeed{
				{
					Name:        "Delta",
					Image:       "lancia-delta.jpg",
					Description: "If your a crazy Italian Slalom, Giant Slalom and Super G skier, 5 time Olympic medal winner (3 gold), what car would you buy? Well Alberto Tomba was all of those and he picked this exhilarating machine, a 1992 **Lancia Delta Integrale HF Evo** 1 that was as equally happy on the snow as he was. \n\nThis particular model was a bespoke car built for him with many of the Evo2 upgrades that were available and a limited edition numbered 96 of 486. It has covered just 11469kms from new and is like new. A true drivers machine and a very angry little car!",
					EngineVol:   2,
					MaxSpeed:    220,
					Votes:       543,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "The legend",
						},
					},
				},
				{
					Name:        "Ypsilon",
					Image:       "lancia-ypsilon.jpg",
					Description: "Distinctive and intriguing new supermini, but Twinair engine is at odds with **Ypsilon**'s luxury positioning. \n\nThe Twinair, while impressively torquey and incredibly willing to work seamlessly around to the rev limiter, also has a very vocal, thrumming soundtrack. \nThere’s no doubt that this remarkable little motor rather encourages the driver to press-on, swapping up and down the ’box to exploit the performance. \nTo that end, the shift action is pretty (although not completely) clean and the lever very well-placed.",
					EngineVol:   1.25,
					MaxSpeed:    183,
					Votes:       103,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "A small funny car",
						},
					},
				},
				{
					Name:        "Rally 037",
					Image:       "Lancia-rally.jpg",
					Description: "Before the Ford Focus, before the Subaru WRX, and before the Mitsubishi EVO, there was the **Lancia 037 Stradale**. \nnThis vehicle is arguably one of the greatest rally cars ever created, despite winning only a single manufacturer’s title in the 1983 season of the World Rally Championship. \nYou see, the Lancia 037 accomplished that feat as a mid-engine, rear-wheel-drive platform running against the seemingly indomitable Audi Quattro. Even as the beast from Ingolstadt kicked off the sport’s inevitable mass migration to all-wheel-drive grip, the Lancia 037 somehow clawed its way to victory over the mighty German competitor. \n\nThe pitched battles fought between these two titans has become the stuff of rally legend, and now, the Lancia 037 sits as the final rear-wheel-drive car to win a WRC manufacturer’s championship.",
					EngineVol:   2,
					MaxSpeed:    220,
					Votes:       658,
					Comments: []CommentSeed{
						{
							DatePosted: time.Date(2016, 07, 10, 11, 0, 0, 0, time.UTC),
							Comment:    "One of the greatest rally cars ever",
						},
					},
				},
				{
					Name:        "Stratos",
					Image:       "lancia-stratos.jpg",
					Description: "Lancia Stratos Description",
					EngineVol:   4.3,
					MaxSpeed:    232,
					Votes:       123,
				},
			},
		},
	}
}
