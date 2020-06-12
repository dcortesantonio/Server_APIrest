package Controllers

import (
	"GoProject/ModelsAPI"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB


//Documentation: https://www.cockroachlabs.com/docs/v20.1/build-a-go-app-with-cockroachdb.html#step-1-install-the-go-pq-driver

// This method gives the history of domains searched.
// Returns: the list of domains searched.
//     Responses:
//       ServersConsulted
func getListServersDB() ModelsAPI.ServersConsulted {

	fmt.Println("Init getListServersDB")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")

	//Init Consult DB: Table of domains consulted.
	statement, err := db.Prepare(`SELECT domainSearched , id 
							FROM SERVERSDB.DOMAIN
							ORDER BY id DESC
							LIMIT 15;`)
	if err != nil {
		log.Fatal("[Error]: Query SELECT list of servers ", err)
	}
	defer statement.Close()
	//End Consult DB.
	table, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	history := ModelsAPI.ServersConsulted{}

	for table.Next() {
		var idDomain int64
		item := ModelsAPI.Item{}
		err = table.Scan(&item.Domain, &idDomain)
		if err != nil {
			log.Fatal(err)
		}
		item.Info = getInfoDomain(idDomain)
		history.Items = append(history.Items, item)

	}
	return history
}

// This method gives information of a domain.
// Returns: the information of a domain.
//     Responses:
//       Server
func getInfoDomain(idDomain int64) ModelsAPI.Server {

	fmt.Println(" Init getInfoDomain Server")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")

	//Init Consult DB: Table of information of server by a domain consulted.
	statement, err := db.Prepare(`SELECT serversChanged, MIN_SSLGrade, PREV_SSLGrade, logo, title, isDown, id
										FROM SERVERSDB.SERVER WHERE domainId = $1;`)
	if err != nil {
		log.Fatal("[Error]: Query SELECT information of a domain ", err)
	}
	defer statement.Close()
	//End Consult DB.

	var serverInfo ModelsAPI.Server
	var serverInfo_id int64
	err = statement.QueryRow(idDomain).Scan( &serverInfo.Servers_Changed, &serverInfo.Min_SSL_Grade, &serverInfo.Previous_SSL_Grade, &serverInfo.Logo, &serverInfo.Title, &serverInfo.Is_Down,&serverInfo_id)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	fmt.Println(" Init getInfoDomain Items")
	fmt.Println("DB ")
	db, err = sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")

	//Init Consult DB: Table of servers by a domain consulted.
	sqlStm, err := db.Prepare(`SELECT address, SSL_Grade, country,owner
									FROM SERVERSDB.SERVERITEM WHERE serverId = $1;`)
	if err != nil {
		log.Fatal("[Error]: Query SELECT information of a domain ", err)
	}
	defer sqlStm.Close()

	serverItems := []ModelsAPI.ServerItem{}
	rows, err := sqlStm.Query(serverInfo_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
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

// This method insert a domain in the Data Base.
// Returns: if the insert has successfully or not.
//     Responses:
//       bool
func insertDomain(domain string, information  ModelsAPI.Server)  (bool){
	//Open Data base SERVERDB.
	fmt.Println(" Init insertDomain  ")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")
	sqlStm, err := db.Prepare("INSERT INTO SERVERSDB.DOMAIN (domainSearched, timeSearched) VALUES ($1, NOW()) RETURNING id;")
	if err != nil {
		log.Fatal("[Error] : Inserting TABLE DOMAIN ", err)
	}
	defer sqlStm.Close()
	var idDomain int64
	err = sqlStm.QueryRow(domain).Scan(&idDomain)
	db.Close()
	if err != nil {
		log.Fatal("[Error] : Query Inserting TABLE DOMAIN ", err)
		return false
	} else {
		return insertServer(idDomain, information)
	}
	return true
}

// This method insert a server in the Data Base.
// Returns: if the insert has successfully or not.
//     Responses:
//       bool
func insertServer(idDomain int64,  information ModelsAPI.Server) bool{
	//Open Data base SERVERDB.
	fmt.Println(" Init insertServer ")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")
	//Documentation: https://blog.friendsofgo.tech/posts/empezando-con-cockroachdb/
	sqlStm, err := db.Prepare("INSERT INTO SERVERSDB.SERVER (serversChanged, MIN_SSLGrade, PREV_SSLGrade, logo, title, isDown, domainId) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
	if err != nil {
		log.Fatal("[Error] : Inserting TABLE SERVER ", err)
	}
	defer sqlStm.Close()
	var serverId int64
	var succes bool
	err = sqlStm.QueryRow(information.Servers_Changed, information.Min_SSL_Grade, information.Previous_SSL_Grade, information.Logo, information.Title, information.Is_Down, idDomain).Scan(&serverId)
	db.Close()
	if err != nil {
		log.Fatal("[Error] : Query Inserting TABLE SERVER ", err)
		return false
	} else {
		succes = insertServerItem(serverId, information.Servers)
	}
	return succes
}

// This method insert servers informations in the Data Base.
// Returns: if the insert has successfully or not.
//     Responses:
//       bool
func insertServerItem(serverID int64,  serverItems []ModelsAPI.ServerItem) bool {
	//Open Data base SERVERDB.
	fmt.Println(" Init insertServerItem")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")
	for _, server := range serverItems {
		sqlStm, err := db.Prepare("INSERT INTO SERVERSDB.SERVERITEM (address, ssl_grade, country, owner, serverId) VALUES ($1, $2, $3, $4, $5);")
		if err != nil {
			log.Fatal("[Error] :  Inserting TABLE SERVERITEM ", err)
			return false
		}
		defer sqlStm.Close()
		_, err = sqlStm.Exec(server.Address, server.SSL_Grade, server.Country, server.Owner, serverID)
		if err != nil {
			log.Fatal("[Error] : Exec TABLE SERVERITEM ", err)
		}
	}
	db.Close()
	return true
}

// This method gives previous information of servers.
// Returns: the list of servers.
//     Responses:
//       ServerItem
func getPreviousServersItems(name string) []ModelsAPI.ServerItem {

	//Open Data base SERVERDB.
	fmt.Println(" Init getPreviousServersItems ")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")
	stmt, err := db.Prepare(`SELECT SERVERSDB.SERVERITEM.address, SERVERSDB.SERVERITEM.ssl_grade, SERVERSDB.SERVERITEM.country, SERVERSDB.SERVERITEM.owner 
							FROM SERVERSDB.DOMAIN, SERVERSDB.SERVER, SERVERSDB.SERVERITEM
							WHERE domainSearched = $1 AND timeSearched < NOW() - INTERVAL '1 hour'
							AND SERVERSDB.SERVER.domainId = SERVERSDB.domain.id AND SERVERSDB.SERVER.id = SERVERSDB.SERVERITEM.serverId ;`)

	if err != nil {
		log.Fatal("[Error] Query SELECT previous information of servers Items ", err)
	}

	fmt.Println("SELECT")

	defer stmt.Close()
	rows, err := stmt.Query(name)
	fmt.Println("QUERY")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	servers := []ModelsAPI.ServerItem{}
	for rows.Next() {
		var server ModelsAPI.ServerItem
		err = rows.Scan(&server.Address, &server.SSL_Grade, &server.Country, &server.Owner)
		fmt.Println("add" + server.Address+"-" + server.SSL_Grade)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}
	db.Close()
	return servers
}

// This method gives previous SSL of a domain.
// Returns: the list of servers.
//     Responses:
//       string
func getPrevious_SSL_Grade(name string) string {

	//Open Data base SERVERDB.
	fmt.Println(" Init getPrevious_SSL_Grade ")
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("[Error] : Connecting to the database SERVERSDB ", err)
	}
	fmt.Println("DB OK")
	stmt, err := db.Prepare(`SELECT min_sslgrade FROM SERVERSDB.DOMAIN, SERVERSDB.SERVER
									WHERE domainsearched = $1 AND timesearched < NOW() - INTERVAL '1 hour'
									AND domainid=SERVERSDB.domain.id ORDER BY timesearched DESC LIMIT 1;`)
	if err != nil {
		log.Fatal("[Error] Query SELECT previous SSL Grade ", err)
	}
	defer stmt.Close()
	var ssl string
	log.Print(" \n *qUERY ROW:::")
	err = stmt.QueryRow(name).Scan(&ssl)

	fmt.Print(" **** Dom :"+name +" ,del SLL->" + ssl + "<-")
	return ssl

}
