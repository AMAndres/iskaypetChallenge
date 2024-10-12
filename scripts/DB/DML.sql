INSERT INTO test."genders_enum" ("id", "description") VALUES
  (0, 'Male'),
  (1, 'Female');

INSERT INTO test."species_enum" ("id", "description") VALUES
  (0, 'Dog'),
  (1, 'Cat');

INSERT INTO test."pets" ("id", "name", "species", "gender", "birthday") VALUES
  (0, 'Faro', 1, 1, '2023-01-01'),
  (1, 'Estaca', 1, 2, '2022-05-15'),
  (2, 'Max', 2, 1, '2021-10-08'),
  (3, 'Funesto', 2, 2, '2024-03-20');
