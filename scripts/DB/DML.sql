INSERT INTO test."genders_enum" ("id", "description") VALUES
  (0, 'Male'),
  (1, 'Female');

INSERT INTO test."species_enum" ("id", "description") VALUES
  (0, 'Dog'),
  (1, 'Cat');

INSERT INTO test."pets" ("name", "species", "gender", "birthday") VALUES
  ('Faro', 0, 0, '2023-01-01'),
  ('Estaca', 0, 1, '2022-01-01'),
  ('Firulais', 0, 0, '2021-01-01'),
  ('Max', 1, 0, '2021-01-01'),
  ('Funesto', 1, 1, '2024-01-01');
