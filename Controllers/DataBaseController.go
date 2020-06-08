package Controllers

import (
	"GoProject/ModelsAPI"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

var db *sql.DB

//Documentation: https://www.cockroachlabs.com/docs/v20.1/build-a-go-app-with-cockroachdb.html#step-1-install-the-go-pq-driver

func getListServersDB() ModelsAPI.ServersConsulted {
	openDB()
	//Init Consult DB: Table of domains consulted.
	statement, err := db.Prepare(`SELECT DISTINCT ON (domainSearched) domainSearched, id
							FROM DOMAIN
							GROUP BY domainSearched, id
							ORDER BY domainSearched, MAX(timeSearched), id
							LIMIT 100;`)
	if err != nil {
		log.Fatal("[Error]: Query SELECT ", err)
	}
	defer statement.Close()
	//End Consult DB.
	table, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	history := ModelsAPI.ServersConsulted{}
	item := ModelsAPI.Item{}

	for table.Next() {
		var idDomain int64
		err = table.Scan(&item.Domain, &idDomain)
		if err != nil {
			log.Fatal(err)
		}
		item.Info = getInfoDomain(idDomain)
		history.Items = append(history.Items, item)
		//items = append(items, item)
	}
	//history.Items = items
	return history
}
func getInfoDomain(idDomain int64) ModelsAPI.Server {
	openDB()
	consultServer := "SELECT id, serversChanged, MIN_SSLGrade, PREV_SSLGrade, logo, title, isDown FROM SERVER WHERE domainId =" + strconv.Itoa(int(idDomain)) + ";"
	//Init Consult DB: Table of information of server by a domain consulted.
	statement, err := db.Prepare(consultServer)
	if err != nil {
		log.Fatal("[Error]: Query SELECT ", err)
	}
	defer statement.Close()
	//End Consult DB.
	table, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	var serverInfo ModelsAPI.Server
	var serverInfo_id int64
	err = statement.QueryRow(idDomain).Scan(&serverInfo_id, &serverInfo.Servers_Changed, &serverInfo.Min_SSL_Grade, &serverInfo.Previous_SSL_Grade, &serverInfo.Logo, &serverInfo.Title, &serverInfo.Is_Down)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	openDB()
	consultServerItems := "SELECT address, ssl_grade, country, owner FROM SERVER WHERE infoserver_id =" + strconv.Itoa(int(serverInfo_id)) + ";"
	//Init Consult DB: Table of servers by a domain consulted.
	stmt, err := db.Prepare(consultServerItems)
	if err != nil {
		log.Fatal("[Error]: Query SELECT ", err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(serverInfo_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	serverItems := []ModelsAPI.ServerItem{}

	for rows.Next() {
		var server ModelsAPI.ServerItem
		err = rows.Scan(&server.Address, &server.SSL_Grade, &server.Country, &server.Owner)
		if err != nil {
			log.Fatal(err)
		}
		serverItems = append(serverItems, server)
	}
	db.Close()

	serverInfo.Servers = serverItems
	return serverInfo
}
func openDB() {
	db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/bank?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")
	if err != nil {
		log.Fatal("[Error]: Can not connecting to the database: ", err)
	}
	defer db.Close()

}
func getQuery(stml *DB) *sql.table {

	table, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()
	return table
}
