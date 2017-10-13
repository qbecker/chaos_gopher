package database

const (
	GetContainerRunStatement    = "SELECT id, dockerfile, script, environmentVariables FROM ContainerRun WHERE ID=?"
	InsertContainerRunStatement = "INSERT INTO ContainerRun(id, testID, dockerfile, script, environmentVariables) VALUES (? ,?, ?, ?, ?)"
	DeleteContainerRunStatement = "DELETE FROM ContainerRun WHERE id=?"
)
