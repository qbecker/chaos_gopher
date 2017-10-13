package database

const Schema = `CREATE TABLE IF NOT EXISTS 'ContainerRun' (
	'id'	INTEGER NOT NULL UNIQUE,
	'dockerfile'	TEXT NOT NULL,
	'script'	TEXT NOT NULL,
	'environmentVariables'	TEXT NOT NULL DEFAULT "",
	PRIMARY KEY('id')
);`
