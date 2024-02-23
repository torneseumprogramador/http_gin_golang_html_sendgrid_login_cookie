-- Criar a base de dados se não existir
CREATE DATABASE IF NOT EXISTS desafio_go;
USE desafio_go;

-- Criar a tabela de Administradores se não existir
CREATE TABLE IF NOT EXISTS Administradores (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Senha VARCHAR(255) NOT NULL,
    Super tinyint NOT NULL default(0),
    PRIMARY KEY (Id)
);

-- Criar a tabela de Donos se não existir
CREATE TABLE IF NOT EXISTS Donos (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    Telefone VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

-- Criar a tabela de Pets se não existir
CREATE TABLE IF NOT EXISTS Pets (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    DonoId VARCHAR(255) NOT NULL,
    Tipo int,
    PRIMARY KEY (Id),
    FOREIGN KEY (DonoId) REFERENCES Donos(Id)
);

CREATE TABLE IF NOT EXISTS fornecedores (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS tokens (
    Id VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);
