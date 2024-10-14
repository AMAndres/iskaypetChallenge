package infraestructure

import (
	"context"
	"database/sql"

	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
)

type PostgresPetRepository struct {
	db *sql.DB
}

func NewPostgresPetRepository(db *sql.DB) *PostgresPetRepository {
	return &PostgresPetRepository{
		db: db,
	}
}

func (pr *PostgresPetRepository) GetPetById(ctx context.Context, id int64) (*petDomain.Pet, error) {

	var pet petDomain.Pet

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	row := tx.QueryRowContext(ctx, "SELECT * FROM test.pets WHERE id = $1", id)

	err = row.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Gender, &pet.Birthday)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErrors.NewNotFoundError("pet not found")
		}
		return nil, err
	}

	return &pet, nil
}

func (pr *PostgresPetRepository) GetAllPets(ctx context.Context) (*[]petDomain.Pet, error) {

	var pets []petDomain.Pet

	// Transaction management
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// DB query
	rows, err := tx.QueryContext(ctx, "SELECT * FROM test.pets")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErrors.NewNotFoundError("pets not found")
		}
		return nil, err
	}

	defer rows.Close()

	// Reads rows retrieved
	for rows.Next() {
		var pet petDomain.Pet
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Gender, &pet.Birthday)
		if err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pets, nil
}

func (pr *PostgresPetRepository) GetAllPetsBySpecies(ctx context.Context, species string) (*[]petDomain.Pet, error) {

	var pets []petDomain.Pet

	// Transaction management
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// DB query
	rows, err := tx.QueryContext(ctx, "SELECT p.id, p.name, p.species, p.gender, p.birthday FROM test.pets p inner join test.species_enum s on s.id = p.species WHERE s.description = $1;", species)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErrors.NewNotFoundError("pets not found")
		}
		return nil, err
	}

	defer rows.Close()

	// Reads rows retrieved
	for rows.Next() {
		var pet petDomain.Pet
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Gender, &pet.Birthday)
		if err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pets, nil
}

func (pr *PostgresPetRepository) AddPet(ctx context.Context, pet *petDomain.Pet) (*petDomain.Pet, error) {

	// Transaction management
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Query
	query := "INSERT INTO test.pets (name, species, gender, birthday) VALUES ($1, $2, $3, $4) RETURNING id"
	lastId := 0
	tx.QueryRowContext(ctx, query,
		pet.Name, pet.Species, pet.Gender, pet.Birthday).Scan(&lastId)

	pet.ID = int64(lastId)

	return pet, nil
}

func (pr *PostgresPetRepository) MostNumerousSpecies(ctx context.Context) (*apiModels.KpiMostNumerousSpecies, error) {

	var response apiModels.KpiMostNumerousSpecies

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := "SELECT p.species, s.description, COUNT(*) AS total_pets FROM test.pets p inner join test.species_enum s on s.id = p.species GROUP BY p.species, s.description LIMIT 1;"
	row := tx.QueryRowContext(ctx, query)

	err = row.Scan(&response.MostNumerousSpecies, &response.MostNumerousSpeciesDescription, &response.MostNumerousSpeciesQty)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErrors.NewNotFoundError("species not found")
		}
		return nil, err
	}

	return &response, nil
}
