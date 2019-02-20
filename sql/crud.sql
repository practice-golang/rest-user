-- Insert / Crud
INSERT INTO users ("username", "username", "email", "password", "auth") VALUES ('john', 'John Doe', 'john@mail.com', '1234', '0');
INSERT INTO users ("username", "username", "email", "password", "auth") VALUES ('jane', 'Jane Eyre', 'jane@mail.com', '5678', '0');
INSERT INTO users ("username", "username", "email", "password", "auth") VALUES ('robert', 'Robert Shin', 'robert@mail.com', 'qwer', '1');

-- Select / cRud
SELECT * FROM "users" ORDER BY "_id" ASC LIMIT 100;
SELECT * FROM "users" WHERE "username"='Master' and "password"='12345';

-- Update / crUd
UPDATE "users" SET "email"='bob@mail.com', WHERE "username"='robert';

-- Delete / cruD
Delete FROM "users" WHERE "_id"=1;
Delete FROM "users" WHERE "username"='jane';

-- Find
SELECT * FROM "users" WHERE "email" LIKE '%mail.com%' ORDER BY "_id" ASC LIMIT 100;

