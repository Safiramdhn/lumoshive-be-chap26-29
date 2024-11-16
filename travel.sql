CREATE TABLE destination (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    description TEXT,
	map_url TEXT
);

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    destination_id INT REFERENCES destination(id),
    date DATE NOT NULL,
    title VARCHAR(100),
    description TEXT,
    photo_url VARCHAR,
    price decimal(10,2)
);

CREATE TABLE review (
    id SERIAL PRIMARY KEY,
    event_id INT REFERENCES event(id),
    rating DECIMAL(2, 1),
    comment TEXT
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    event_id INT REFERENCES event(id),
    status VARCHAR(50),
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert dummy data into the destination table
INSERT INTO destination (name, location, description, map_url)
VALUES
('Paris', 'France', 'The capital city of France, known for its art, fashion, and culture.', 'https://www.google.com/maps?q=48.8566,2.3522'),
('Tokyo', 'Japan', 'A vibrant city known for its skyscrapers, shopping, and technology.', 'https://www.google.com/maps?q=35.6895,139.6917'),
('New York', 'USA', 'A bustling metropolis, home to landmarks such as Times Square and Central Park.', 'https://www.google.com/maps?q=40.7128,-74.0060'),
('Sydney', 'Australia', 'Famous for its stunning harbor, opera house, and beaches.', 'https://www.google.com/maps?q=-33.8688,151.2093'),
('London', 'UK', 'The capital city of the UK, rich in history and culture, with landmarks like Big Ben.', 'https://www.google.com/maps?q=51.5074,-0.1278'),
('Rome', 'Italy', 'Known for its ancient ruins, such as the Colosseum and the Roman Forum.', 'https://www.google.com/maps?q=41.9028,12.4964'),
('Dubai', 'UAE', 'A luxurious city known for its modern architecture, shopping, and skyscrapers.', 'https://www.google.com/maps?q=25.2048,55.2708'),
('Bangkok', 'Thailand', 'Famous for its vibrant street life, temples, and delicious food.', 'https://www.google.com/maps?q=13.7563,100.5018'),
('Barcelona', 'Spain', 'A city with rich history, known for its unique architecture and Mediterranean beaches.', 'https://www.google.com/maps?q=41.3851,2.1734'),
('Rio de Janeiro', 'Brazil', 'Known for its Carnival festival, beaches, and the Christ the Redeemer statue.', 'https://www.google.com/maps?q=-22.9068,-43.1729'),
('Moscow', 'Russia', 'Capital city of Russia, famous for Red Square, the Kremlin, and its history.', 'https://www.google.com/maps?q=55.7558,37.6176'),
('Los Angeles', 'USA', 'Known for Hollywood, entertainment, and its beaches.', 'https://www.google.com/maps?q=34.0522,-118.2437'),
('Cairo', 'Egypt', 'Home to the Pyramids of Giza and a long history of ancient civilization.', 'https://www.google.com/maps?q=30.0444,31.2357'),
('Cape Town', 'South Africa', 'A coastal city known for its beaches, mountains, and vineyards.', 'https://www.google.com/maps?q=-33.9249,18.4241'),
('Istanbul', 'Turkey', 'A city that straddles Europe and Asia, rich in culture, history, and architecture.', 'https://www.google.com/maps?q=41.0082,28.9784');

-- Insert dummy data into the event table
INSERT INTO event (destination_id, date, title, description, price)
VALUES
(1, '2024-12-01', 'Paris Fashion Week', 'A world-renowned fashion event showcasing the latest trends.', 1000.00),
(2, '2024-11-15', 'Tokyo Tech Expo', 'An exhibition of the latest technology and innovations from Tokyo.', 200.00),
(3, '2024-12-10', 'New York Christmas Parade', 'A festive event celebrating the Christmas season in New York.', 200.00),
(4, '2024-11-20', 'Sydney Opera Performance', 'A spectacular opera performance at the Sydney Opera House.', 600.00),
(5, '2024-12-05', 'London Winter Wonderland', 'A holiday-themed event with rides, ice skating, and Christmas markets.', 750.00),
(6, '2024-11-25', 'Rome Ancient Ruins Tour', 'A guided tour through Rome’s ancient historical sites.', 1000.00),
(7, '2024-12-12', 'Dubai Luxury Shopping Festival', 'A shopping festival featuring top luxury brands in Dubai.', 500.00),
(8, '2024-11-30', 'Bangkok Street Food Festival', 'A festival dedicated to Bangkok’s famous street food.', 850.00),
(9, '2024-12-15', 'Barcelona Beach Party', 'A beach party with live music and dancing by the sea.', 800.00),
(10, '2024-11-28', 'Rio Carnival Parade', 'An annual carnival parade filled with music, dancing, and vibrant costumes.', 1200.00),
(11, '2024-12-05', 'Moscow Winter Market', 'A traditional winter market offering food, drinks, and gifts.', 900.00),
(12, '2024-12-18', 'Los Angeles Film Festival', 'A festival celebrating independent films and filmmakers in Los Angeles.', 900.00),
(13, '2024-11-22', 'Cairo Pyramids Light Show', 'A spectacular light and sound show at the Pyramids of Giza.', 890.00),
(14, '2024-12-03', 'Cape Town Wine Tasting', 'A wine-tasting event showcasing South African wines in Cape Town.', 1500.00),
(15, '2024-11-17', 'Istanbul Cultural Fest', 'A celebration of Turkish culture with music, dance, and food.', 1300.00);

-- Insert dummy data into the review table
-- Insert dummy review data with random event_id based on transaction
INSERT INTO review (transaction_id, rating, comment)
VALUES
(1, 4.3, 'Great event, but the venue could have been better.'),
(2, 4.8, 'Amazing experience, highly recommend!'),
(3, 3.9, 'Nice event but it lacked some features I expected.'),
(4, 4.7, 'Perfect, the event was well-organized and fun.'),
(5, 4.5, 'Loved the experience, but the food could be improved.'),
(6, 4.0, 'The event was good, but I had some issues with the schedule.'),
(7, 4.6, 'Great show and excellent performances throughout.'),
(8, 4.2, 'Good experience, but it felt a bit too crowded for my liking.'),
(9, 5.0, 'Absolutely amazing! A must-visit event in Rio.'),
(10, 4.4, 'The event was fun, but I wish there were more activities.');



-- Insert dummy transaction data with random event_id
INSERT INTO transaction (event_id, customer_id, status, price)
VALUES
(1, 1, 'Completed', 120.00),
(3, 2, 'Pending', 250.00),
(5, 3, 'Completed', 90.00),
(7, 4, 'Cancelled', 300.00),
(2, 5, 'Completed', 150.00),
(4, 6, 'Pending', 200.00),
(6, 7, 'Completed', 180.00),
(8, 8, 'Cancelled', 120.00),
(10, 9, 'Completed', 350.00),
(9, 10, 'Pending', 250.00);

CREATE TABLE customer (
	id SERIAL PRIMARY KEY,
	email varchar not null,
	phone_number varchar not null
)

ALTER TABLE customer ADD COLUMN name varchar not null

ALTER TABLE transaction
ADD COLUMN customer_id INT REFERENCES customer(id);

ALTER TABLE transaction
ADD COLUMN number_of_ticket INT

INSERT INTO customer (name, email, phone_number)
VALUES
('Alice', 'alice@example.com', '1234567890'),
('Bob', 'bob@example.com', '0987654321'),
('Charlie', 'charlie@example.com', '1112223333'),
('Diana', 'diana@example.com', '4445556666'),
('Eve', 'eve@example.com', '7778889999'),
('Frank', 'frank@example.com', '1010101010'),
('Grace', 'grace@example.com', '1212121212'),
('Heidi', 'heidi@example.com', '1313131313'),
('Ivan', 'ivan@example.com', '1414141414'),
('Judy', 'judy@example.com', '1515151515');

ALTER TABLE destination ADD COLUMN map TEXT

CREATE TABLE tour_plan (
	id SERIAL PRIMARY KEY,
	event_id INTEGER REFERENCES event(id),
	activities TEXT[],
	description VARCHAR
);

-- Insert sample data into the tour_plan table
INSERT INTO tour_plan (event_id, activities, description)
VALUES
-- Paris Fashion Week
(1, ARRAY['Arrival in Paris', 'Check-in at 5-Star Hotel', 'Welcome Dinner'], 'Day 1: Arrival and check-in. Enjoy a welcome dinner.'),
(1, ARRAY['Paris Fashion Week Attendance', 'Private Fashion Boutique Tour'], 'Day 2: Attend Paris Fashion Week shows and explore exclusive boutiques.'),
(1, ARRAY['Guided Louvre Museum Tour', 'Eiffel Tower Evening Visit'], 'Day 3: Discover the Louvre’s masterpieces and visit the Eiffel Tower at night.'),

-- Tokyo Tech Expo
(2, ARRAY['Arrival in Tokyo', 'Check-in at luxury hotel'], 'Day 1: Arrival and hotel check-in in Tokyo.'),
(2, ARRAY['Tokyo Tech Expo - Day 1', 'Networking Lunch'], 'Day 2: Attend the first day of the Tokyo Tech Expo.'),
(2, ARRAY['Tokyo Tech Expo - Day 2', 'Dinner at Robot Restaurant'], 'Day 3: Attend the second day of the Expo and enjoy a unique dining experience.'),

-- New York Christmas Parade
(3, ARRAY['Arrival in New York', 'Check-in at Times Square Hotel'], 'Day 1: Arrival and check-in at a hotel in Times Square.'),
(3, ARRAY['Christmas Parade Viewing', 'Rockefeller Center Ice Skating'], 'Day 2: Watch the Christmas Parade and go ice skating at Rockefeller Center.'),
(3, ARRAY['Broadway Show', 'Farewell Dinner'], 'Day 3: Enjoy a Broadway show and farewell dinner.'),

-- Sydney Opera Performance
(4, ARRAY['Arrival in Sydney', 'Check-in at Sydney Harbor Hotel'], 'Day 1: Arrival and check-in with a view of Sydney Harbor.'),
(4, ARRAY['Sydney Opera House Tour', 'Opera Performance'], 'Day 2: Explore the Opera House and attend an opera performance.'),
(4, ARRAY['Bondi Beach Visit', 'Shopping in The Rocks'], 'Day 3: Relax at Bondi Beach and shop in the historic Rocks area.'),

-- London Winter Wonderland
(5, ARRAY['Arrival in London', 'Check-in at Central London Hotel'], 'Day 1: Arrival in London and hotel check-in.'),
(5, ARRAY['Winter Wonderland Visit', 'Ice Skating and Christmas Markets'], 'Day 2: Enjoy the Winter Wonderland with skating and shopping at Christmas markets.'),
(5, ARRAY['London Sightseeing Tour', 'Afternoon Tea Experience'], 'Day 3: Sightseeing around London and traditional afternoon tea.'),

-- Rome Ancient Ruins Tour
(6, ARRAY['Arrival in Rome', 'Check-in at historic hotel'], 'Day 1: Arrival in Rome and check-in at a historic hotel.'),
(6, ARRAY['Ancient Ruins Guided Tour', 'Visit to the Colosseum'], 'Day 2: Explore the ancient ruins of Rome, including the Colosseum.'),
(6, ARRAY['Vatican City Tour', 'Dinner at local Italian restaurant'], 'Day 3: Visit Vatican City and enjoy an authentic Italian dinner.'),

-- Dubai Luxury Shopping Festival
(7, ARRAY['Arrival in Dubai', 'Check-in at luxury hotel'], 'Day 1: Arrival and luxury hotel check-in in Dubai.'),
(7, ARRAY['Shopping at Dubai Mall', 'Visit to Burj Khalifa'], 'Day 2: Shop at Dubai Mall and visit the Burj Khalifa.'),
(7, ARRAY['Luxury Desert Safari', 'Dinner at desert camp'], 'Day 3: Enjoy a luxury desert safari and dinner under the stars.'),

-- Bangkok Street Food Festival
(8, ARRAY['Arrival in Bangkok', 'Check-in at city center hotel'], 'Day 1: Arrival and check-in at a central Bangkok hotel.'),
(8, ARRAY['Street Food Festival', 'Temple Visit'], 'Day 2: Attend the Street Food Festival and visit famous temples.'),
(8, ARRAY['Floating Market Tour', 'Thai Cooking Class'], 'Day 3: Visit a floating market and take a Thai cooking class.'),

-- Barcelona Beach Party
(9, ARRAY['Arrival in Barcelona', 'Check-in at beachside hotel'], 'Day 1: Arrival in Barcelona and check-in at a beachside hotel.'),
(9, ARRAY['Beach Party', 'Live Music and Dance'], 'Day 2: Join the beach party with live music and dancing by the sea.'),
(9, ARRAY['City Tour', 'Sagrada Familia Visit'], 'Day 3: Explore the city and visit the famous Sagrada Familia.'),

-- Rio Carnival Parade
(10, ARRAY['Arrival in Rio', 'Check-in at hotel near Copacabana Beach'], 'Day 1: Arrival in Rio and check-in at a hotel near Copacabana Beach.'),
(10, ARRAY['Attend Rio Carnival Parade', 'Samba Dance Class'], 'Day 2: Watch the Rio Carnival Parade and learn Samba dancing.'),
(10, ARRAY['Christ the Redeemer Tour', 'Sugarloaf Mountain Cable Car'], 'Day 3: Visit Christ the Redeemer and take a cable car up Sugarloaf Mountain.'),

-- Moscow Winter Market
(11, ARRAY['Arrival in Moscow', 'Check-in at city center hotel'], 'Day 1: Arrival and check-in at a central Moscow hotel.'),
(11, ARRAY['Visit Moscow Winter Market', 'Kremlin Tour'], 'Day 2: Explore the Moscow Winter Market and take a Kremlin tour.'),
(11, ARRAY['Red Square Tour', 'Traditional Russian Dinner'], 'Day 3: Visit Red Square and enjoy a traditional Russian dinner.'),

-- Los Angeles Film Festival
(12, ARRAY['Arrival in Los Angeles', 'Check-in at hotel near Hollywood'], 'Day 1: Arrival in LA and check-in near Hollywood.'),
(12, ARRAY['Film Festival Opening', 'Q&A with Filmmakers'], 'Day 2: Attend the film festival opening and Q&A session.'),
(12, ARRAY['Hollywood Tour', 'Walk of Fame Visit'], 'Day 3: Tour Hollywood and visit the Walk of Fame.'),

-- Cairo Pyramids Light Show
(13, ARRAY['Arrival in Cairo', 'Check-in at Giza hotel'], 'Day 1: Arrival in Cairo and check-in near the Pyramids of Giza.'),
(13, ARRAY['Pyramids Light Show', 'Dinner with Nile View'], 'Day 2: Watch the Pyramids Light Show and enjoy a Nile-view dinner.'),
(13, ARRAY['Museum of Egyptian Antiquities', 'Shopping at Khan el-Khalili Bazaar'], 'Day 3: Visit the museum and shop at the historic bazaar.'),

-- Cape Town Wine Tasting
(14, ARRAY['Arrival in Cape Town', 'Check-in at vineyard hotel'], 'Day 1: Arrival and check-in at a vineyard hotel in Cape Town.'),
(14, ARRAY['Wine Tasting Tour', 'Cape Peninsula Tour'], 'Day 2: Take a wine-tasting tour and explore the Cape Peninsula.'),
(14, ARRAY['Table Mountain Visit', 'Dinner at seaside restaurant'], 'Day 3: Visit Table Mountain and have dinner by the sea.'),

-- Istanbul Cultural Fest
(15, ARRAY['Arrival in Istanbul', 'Check-in at historic hotel'], 'Day 1: Arrival in Istanbul and check-in at a historic hotel.'),
(15, ARRAY['Cultural Festival Activities', 'Turkish Cuisine Tasting'], 'Day 2: Participate in cultural activities and taste Turkish cuisine.'),
(15, ARRAY['Bosphorus Cruise', 'Shopping at Grand Bazaar'], 'Day 3: Take a Bosphorus cruise and shop at the Grand Bazaar.');
