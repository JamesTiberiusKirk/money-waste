package models

// GetAllModels returns a list of model instances for gorm migration. Any models
// added to this package that need to be migrated to the database should be
// included in this list.
func GetAllModels() []any {
	return []any{
		&User{},
		&Transaction{},
	}
}
