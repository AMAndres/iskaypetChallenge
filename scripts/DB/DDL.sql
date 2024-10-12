CREATE DATABASE iskaypetChallenge;

CREATE SCHEMA test;

CREATE TABLE test.pets(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    species SMALLINT NOT NULL,
    gender SMALLINT NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE test.species_enum (
    id SMALLSERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE test.genders_enum (
    id SMALLSERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

ALTER TABLE test.pets
    ADD CONSTRAINT fk_pets_species
    FOREIGN KEY(species)
    REFERENCES test.species_enum(id);

ALTER TABLE test.pets
    ADD CONSTRAINT fk_pets_gender
    FOREIGN KEY(gender)
    REFERENCES test.genders_enum(id);