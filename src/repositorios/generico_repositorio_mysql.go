package repositorios

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

// Definição de GenericoRepositorioMySql usando generics.
type GenericoRepositorioMySql[T any] struct {
	DB *sql.DB
}

// getTableName retorna o nome da tabela a partir da struct genérica T.
func (gs *GenericoRepositorioMySql[T]) getTableName() string {
	// Cria uma instância temporária do tipo T para acessar suas propriedades via reflexão.
	var temp T
	t := reflect.TypeOf(temp)

	// Se T for um ponteiro, obtém o tipo ao qual ele aponta.
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Itera pelos campos da struct em busca do campo 'TableName' e extrai o valor da tag 'db'.
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == "tableName" {
			tagValue := field.Tag.Get("table")
			if tagValue != "" {
				return tagValue
			}
		}
	}

	// Se o nome da tabela não for encontrado, retorna uma string vazia ou lança um erro.
	// A abordagem específica depende de como você deseja lidar com essa situação.
	return ""
}

func (gs *GenericoRepositorioMySql[T]) Where(condicoes map[string]string) ([]T, error) {
	var whereClauses []string
	var valores []interface{}

	for chave, valor := range condicoes {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", chave))
		valores = append(valores, valor)
	}

	queryWhere := ""
	if len(whereClauses) > 0 {
		queryWhere = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query := fmt.Sprintf("SELECT * FROM %s%s", gs.getTableName(), queryWhere)

	rows, err := gs.DB.Query(query, valores...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []T
	for rows.Next() {
		var entidade T
		err := scanEntity(rows, &entidade)
		if err != nil {
			log.Println("Erro ao escanear a entidade:", err)
			continue
		}
		result = append(result, entidade)
	}

	return result, nil
}

// Lista implementa a operação de listar entidades do tipo T.
func (gs *GenericoRepositorioMySql[T]) Lista() ([]T, error) {
	rows, err := gs.DB.Query(fmt.Sprintf("SELECT * FROM %s", gs.getTableName()))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []T
	for rows.Next() {
		var entidade T
		err := scanEntity(rows, &entidade)

		if err != nil {
			log.Println("Erro ao escanear a entidade:", err)
			continue
		}
		result = append(result, entidade)
	}
	return result, nil
}

func (gs *GenericoRepositorioMySql[T]) BuscaPorId(id string) (*T, error) {
	rows, err := gs.DB.Query(fmt.Sprintf("SELECT * FROM %s where Id = ?", gs.getTableName()), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entidade T

	for rows.Next() {
		err := scanEntity(rows, &entidade)

		if err != nil {
			log.Println("Erro ao escanear a entidade:", err)
			continue
		}

		break
	}

	return &entidade, nil
}

func scanEntity[T any](rows *sql.Rows, entidade *T) error {
	entType := reflect.TypeOf(entidade).Elem()
	if entType.Kind() != reflect.Struct {
		return fmt.Errorf("scanEntity: Tipo fornecido T não é uma struct")
	}

	var vals []interface{}

	for i := 0; i < entType.NumField(); i++ {
		structField := entType.Field(i)
		dbTag := structField.Tag.Get("field")
		if dbTag == "" {
			continue // Ignora campos sem a tag `field`
		}

		fieldValue := reflect.ValueOf(entidade).Elem().Field(i)
		if !fieldValue.CanSet() {
			log.Printf("O campo %s não é exportado e será ignorado\n", dbTag)
			continue
		}
		vals = append(vals, fieldValue.Addr().Interface())
	}

	if err := rows.Scan(vals...); err != nil {
		return fmt.Errorf("scanEntity: falha no scan dos campos da entidade: %w", err)
	}

	return nil
}

func (gs *GenericoRepositorioMySql[T]) Adicionar(entidade T) error {
	val := reflect.ValueOf(entidade)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Adicionar: Tipo fornecido T não é uma struct ou ponteiro para struct")
	}

	var campos []string
	var placeholders []string
	var valores []interface{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		dbTag := field.Tag.Get("field")
		if dbTag == "" {
			continue // Ignora campos sem a tag `field`
		}

		// Verifica se o campo é o Id e se está vazio, então preenche com um novo UUID
		if dbTag == "id" {
			idField := val.Field(i)
			if idField.Kind() == reflect.String && idField.String() == "" {
				newUUID := uuid.New().String()
				valores = append(valores, newUUID)
			} else {
				valores = append(valores, idField.Interface())
			}
		} else {
			valores = append(valores, val.Field(i).Interface())
		}

		campos = append(campos, dbTag)
		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		gs.getTableName(),
		strings.Join(campos, ", "),
		strings.Join(placeholders, ", "))

	_, err := gs.DB.Exec(query, valores...)
	if err != nil {
		log.Printf("Erro ao adicionar entidade: %v\n", err)
		return err
	}

	return nil
}

func (gs *GenericoRepositorioMySql[T]) Alterar(entidade T) error {
	val := reflect.ValueOf(entidade)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Alterar: Tipo fornecido T não é uma struct ou ponteiro para struct")
	}

	var campos []string
	var valores []interface{}
	var idValue interface{}
	idFieldFound := false

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		dbTag := field.Tag.Get("field")
		if dbTag == "" {
			continue // Ignora campos sem a tag `field`
		}

		fieldValue := val.Field(i).Interface()

		// Verifica se o campo é o ID
		if dbTag == "id" {
			idValue = fieldValue
			idFieldFound = true
			continue // Não inclui o ID nos campos a serem atualizados
		}

		campos = append(campos, fmt.Sprintf("%s = ?", dbTag))
		valores = append(valores, fieldValue)
	}

	if !idFieldFound || idValue == nil {
		return fmt.Errorf("Alterar: ID não encontrado ou nulo")
	}

	// Adiciona o ID ao final da lista de valores para a cláusula WHERE
	valores = append(valores, idValue)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?",
		gs.getTableName(),
		strings.Join(campos, ", "))

	_, err := gs.DB.Exec(query, valores...)
	if err != nil {
		log.Printf("Erro ao alterar entidade: %v\n", err)
		return err
	}

	return nil
}

func (gs *GenericoRepositorioMySql[T]) ApagarPorId(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", gs.getTableName())

	_, err := gs.DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao remover entidade: %v\n", err)
		return err
	}

	return nil
}
