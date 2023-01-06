package repositories

import "todo/bin/repositories/models"

var ModelTables []interface{} = []interface{}{
	&models.Activities{},
	&models.Todos{},
}
