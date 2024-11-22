package main

import (
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// Abrir o banco de dados Pebble
	dbPebble, err := pebble.Open("./chaindata", &pebble.Options{})
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados Pebble: %v", err)
	}
	defer dbPebble.Close()

	// Abrir o banco de dados LevelDB
	dbLevelDB, err := leveldb.OpenFile("./leveldb_data", nil)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados LevelDB: %v", err)
	}
	defer dbLevelDB.Close()

	// Iterar sobre as chaves e valores no banco de dados Pebble
	iterPebble, err := dbPebble.NewIter(nil)
	defer iterPebble.Close()

	for iterPebble.First(); iterPebble.Valid(); iterPebble.Next() {
		key := iterPebble.Key()
		value := iterPebble.Value()

		// Gravar a chave e o valor no banco LevelDB
		err := dbLevelDB.Put(key, value, nil)
		if err != nil {
			log.Printf("Erro ao escrever no LevelDB: %v", err)
			continue
		}

		// Imprimir chave e valor exportados
		fmt.Printf("Chave: %s, Valor: %s\n", key, value)
	}

	// Verificar erros do iterador
	if err := iterPebble.Error(); err != nil {
		log.Fatalf("Erro ao iterar no banco Pebble: %v", err)
	}

	fmt.Println("Exportação concluída com sucesso!")
}
