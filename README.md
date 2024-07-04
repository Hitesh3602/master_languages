SQL Documentation : master_languages

Introduction:

The master_languages table is designed to store information about different languages. Its structure includes:

id: An integer serving as the primary key, uniquely identifying each language entry.
name: A varchar field for the full name of the language (e.g., "French", "English").
short_code: A varchar field for the language's short code (e.g., "FR", "EN").
is_active: A boolean field indicating whether the language is currently active in the system.
created_at and updated_at: Timestamp fields to track when the entry was created and last updated.

The table supports basic CRUD (Create, Read, Update, Delete) operations

Create Table:
CREATE TABLE master_languages (
 id BIGSERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL UNIQUE,
 short_code VARCHAR(255) NOT NULL UNIQUE,
 is_active BOOLEAN NOT NULL,
 created_at TIMESTAMP,
 updated_at TIMESTAMP
);


Insert Values:
INSERT INTO master_languages (id, name, short_code, is_active, created_at, updated_at)
VALUES (1, 'French', 'FR', true, now(), now()),
      (2, 'English', 'EN', true, now(), now()),
      (3, 'Spanish', 'ES', true, now(), now());
INSERT INTO master_languages (id, name, short_code, is_active, created_at, updated_at)
VALUES (4, 'Hindi', 'HI', false, now(), now()),
      (5, 'Marathi', 'MA', false, now(), now()),
      (6, 'Sindhi', 'SI', true, now(), now());


GET/rms/master_languages: Retrieves a list of all master configurations.
SELECT * FROM master_languages ORDER BY id;


POST /rms/master_languages: Creates a new master configuration.
INSERT INTO master_languages (id, name, short_code, is_active, created_at, updated_at)
VALUES ('7', 'Tamil', 'TA', false, now(), now());


GET /rms/master_languages/{id}: Retrieves a specific master configuration by its ID.
SELECT * FROM master_languages WHERE id = 1;


GET /rms/master_languages/{language_name): Retrieves a specific master configuration by its language_name
SELECT * FROM master_languages WHERE name = 'Hindi';


PUT/rms/master_languages/{id}: Updates an existing master configuration.
UPDATE master_languages
SET name = $2, short_code = $3, is_active = $4, updated_at = now()
WHERE id = $1;


DELETE/rms/master_languages/{id): Deletes a specific master configuration.
DELETE FROM master_languages WHERE id = 1;


Combination/rms/master_languages ( ID / Name)
SELECT * FROM master_languages WHERE id = # OR name = #;


