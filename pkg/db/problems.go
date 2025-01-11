package db

import (
    "fmt"
    "context"
    "github.com/arthurasanaliev/math-olymp-platform/pkg/models"
)

// GetProblemByID retrieves a problem by its ID 
func (db *DB) GetProblemByID(id int) (*models.Problem, error) {
    query := `SELECT id, title, difficulty, tags, statement, tutorial FROM problems WHERE id = $1`

    var problem models.Problem
    err := db.Conn.QueryRow(context.Background(), query, id).Scan(&problem.ID, &problem.Title,
        &problem.Difficulty, &problem.Tags,
        &problem.Statement, &problem.Tutorial)

    if err != nil {
        return nil, fmt.Errorf("failed to fetch problem: %v", err)
    }

    return &problem, nil
}

// GetAllProblems retrieves all problems
func (db *DB) GetAllProblems() ([]models.Problem, error) {
    query := `SELECT id, title, difficulty, tags, statement, tutorial FROM problems`

    rows, err := db.Conn.Query(context.Background(), query)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch problems: %v", err)
    }
    defer rows.Close()

    var problems []models.Problem
    for rows.Next() {
        var problem models.Problem
        if err := rows.Scan(&problem.ID, &problem.Title, &problem.Difficulty, &problem.Tags, 
            &problem.Statement, &problem.Tutorial); err != nil {
            return nil, fmt.Errorf("failed to scan problem: %v", err)
        }
        problems = append(problems, problem)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("failed to read rows: %v", err)
    }

    return problems, nil
}
