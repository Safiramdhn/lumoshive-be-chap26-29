CREATE TABLE destination (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    description TEXT
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
    user_id INT,
    status VARCHAR(50),
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert dummy data into the destination table
INSERT INTO destination (name, location, description)
VALUES
('Paris', 'France', 'The capital city of France, known for its art, fashion, and culture.'),
('Tokyo', 'Japan', 'A vibrant city known for its skyscrapers, shopping, and technology.'),
('New York', 'USA', 'A bustling metropolis, home to landmarks such as Times Square and Central Park.'),
('Sydney', 'Australia', 'Famous for its stunning harbor, opera house, and beaches.'),
('London', 'UK', 'The capital city of the UK, rich in history and culture, with landmarks like Big Ben.'),
('Rome', 'Italy', 'Known for its ancient ruins, such as the Colosseum and the Roman Forum.'),
('Dubai', 'UAE', 'A luxurious city known for its modern architecture, shopping, and skyscrapers.'),
('Bangkok', 'Thailand', 'Famous for its vibrant street life, temples, and delicious food.'),
('Barcelona', 'Spain', 'A city with rich history, known for its unique architecture and Mediterranean beaches.'),
('Rio de Janeiro', 'Brazil', 'Known for its Carnival festival, beaches, and the Christ the Redeemer statue.'),
('Moscow', 'Russia', 'Capital city of Russia, famous for Red Square, the Kremlin, and its history.'),
('Los Angeles', 'USA', 'Known for Hollywood, entertainment, and its beaches.'),
('Cairo', 'Egypt', 'Home to the Pyramids of Giza and a long history of ancient civilization.'),
('Cape Town', 'South Africa', 'A coastal city known for its beaches, mountains, and vineyards.'),
('Istanbul', 'Turkey', 'A city that straddles Europe and Asia, rich in culture, history, and architecture.');

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
INSERT INTO transaction (event_id, user_id, status, price)
VALUES
(1, 101, 'Completed', 120.00),
(3, 102, 'Pending', 250.00),
(5, 103, 'Completed', 90.00),
(7, 104, 'Cancelled', 300.00),
(2, 105, 'Completed', 150.00),
(4, 106, 'Pending', 200.00),
(6, 107, 'Completed', 180.00),
(8, 108, 'Cancelled', 120.00),
(10, 109, 'Completed', 350.00),
(9, 110, 'Pending', 250.00);
