package models

type Van struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Price       int    `json:"Price"`
	Description string `json:"Description"`
	ImageUrl    string `json:"ImageUrl"`
	VanType     string `json:"VanType"`
}

var Vans = []Van{
	{
		Id: 1, Name: "Modest Explorer", Price: 60,
		Description: "The Modest Explorer is a van designed " +
			"to get you out of the house and into nature. " +
			"This beauty is equipped with solar panels, " +
			"a composting toilet, a water tank and kitchenette. " +
			"The idea is that you can pack up your home and escape for a weekend or even longer!",
		ImageUrl: "http://localhost:8080/van1.png",
		VanType:  "simple",
	},
	{
		Id: 2, Name: "Beach Bum", Price: 80,
		Description:  "Beach Bum is a van inspired by " +
			"surfers and travelers. It was created to be a " +
			"portable home away from home, but with some cool " +
			"features in it you won't find in an ordinary camper.",
		ImageUrl: "http://localhost:8080/van2.png",
		VanType: "rugged",
	},
	{
		Id: 3, Name: "Reliable Red", Price: 100,
		Description: "Reliable Red is a van that was made for travelling. " +
			"The inside is comfortable and cozy, with plenty of space to stretch out in. " +
			"There's a small kitchen, so you can cook if you need to. You'll feel like " +
			"home as soon as you step out of it.",
		ImageUrl: "http://localhost:8080/van3.png",
		VanType: "luxury",
	},
	{
		Id: 4, Name: "Dreamfinder", Price: 65,
		Description: "Dreamfinder is the perfect van to travel in and experience. " +
			"With a ceiling height of 2.1m, you can stand up in this van and there is great " +
			"head room. The floor is a beautiful glass-reinforced plastic (GRP) which is easy to " +
			"clean and very hard wearing. A large rear window and large side windows make it really " +
			"light inside and keep it well ventilated.",
		ImageUrl: "http://localhost:8080/van4.png",
		VanType: "simple",
	},
	{
		Id: 5, Name: "The Cruiser", Price: 120,
		Description: "The Cruiser is a van for those who love to travel in comfort " +
			"and luxury. With its many windows, spacious interior and ample storage space, " +
			"the Cruiser offers a beautiful view wherever you go.",
		ImageUrl: "http://localhost:8080/van5.png",
		VanType: "luxury",
	},
	{
		Id: 6, Name: "Green Wonder", Price: 70,
		Description: "With this van, you can take your travel " +
			"life to the next level. The Green Wonder is a sustainable " +
			"vehicle that's perfect for people who are looking for a stylish, " +
			"eco-friendly mode of transport that can go anywhere.",
		ImageUrl: "http://localhost:8080/van6.png",
		VanType: "rugged",
	},
}


