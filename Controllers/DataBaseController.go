package Controllers

import (
	"GoProject/ModelsAPI"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"strconv"
)

var db *sql.DB
func openDB() {
	log.Print(" OpenDB: ")
	/*db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/serverinformation?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	log.Print(" Open Succes ")
	defer db.Close()*/

}

//Documentation: https://www.cockroachlabs.com/docs/v20.1/build-a-go-app-with-cockroachdb.html#step-1-install-the-go-pq-driver

//Methods of List of Domains Searched
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
	sqlStm, err := db.Prepare(consultServerItems)
	if err != nil {
		log.Fatal("[Error]: Query SELECT ", err)
	}

	defer sqlStm.Close()
	rows, err := sqlStm.Query(serverInfo_id)
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
//End of Methods of List of Domains Searched

func insertDomain(domain string, information  ModelsAPI.Server)  (bool){
	openDB()
	sqlStm, err := db.Prepare("INSERT INTO DOMAIN (domain, timeSearched) VALUES (domain, NOW()) RETURNING id;")
	if err != nil {
		log.Fatal("[Error] : Inserting TABLE DOMAIN: ", err)
	}
	defer sqlStm.Close()
	var idDomain int64
	err = sqlStm.QueryRow(domain).Scan(&idDomain)
	db.Close()
	if err != nil {
		log.Fatal("Error Query Inserting TABLE DOMAIN: ", err)
		return false
	} else {
		return insertServer(idDomain, information)
	}
	return true
}
func insertServer(idDomain int64,  information ModelsAPI.Server) bool{
	openDB()
	//Documentation: https://blog.friendsofgo.tech/posts/empezando-con-cockroachdb/
	sqlStm, err := db.Prepare("INSERT INTO SERVER (serversChanged, MIN_SSLGrade, PREV_SSLGrade, logo, title, isDown, domainId) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
	if err != nil {
		log.Fatal("Error Inserting TABLE SERVER:", err)
	}
	defer sqlStm.Close()
	var serverId int64
	var succes bool
	err = sqlStm.QueryRow(information.Servers_Changed, information.Min_SSL_Grade, information.Previous_SSL_Grade, information.Logo, information.Title, information.Is_Down, idDomain).Scan(&serverId)
	db.Close()
	if err != nil {
		log.Fatal("Error Query Inserting TABLE SERVER:", err)
		return false
	} else {
		succes = insertServerItem(serverId, information.Servers)
	}
	return succes
}

func insertServerItem(serverID int64,  serverItems []ModelsAPI.ServerItem) bool {
	openDB()
	for _, server := range serverItems {
		sqlStm, err := db.Prepare("INSERT INTO SERVERITEM (address, ssl_grade, country, owner, serverId) VALUES ($1, $2, $3, $4, $5);")
		if err != nil {
			log.Fatal("Error Inserting TABLE SERVERITEM: ", err)
			return false
		}
		defer sqlStm.Close()
		_, err = sqlStm.Exec(server.Address, server.SSL_Grade, server.Country, server.Owner, serverID)
		if err != nil {
			log.Fatal("Error Exec TABLE SERVERITEM: ", err)
		}
	}
	db.Close()
	return true
}
func InsertInfoDomain(idDomain int64,  infoDomain ModelsAPI.Server){
	openDB()
	//Documentation: https://blog.friendsofgo.tech/posts/empezando-con-cockroachdb/
	sqlStm, err := db.Prepare("INSERT INTO SERVER (serversChanged, MIN_SSLGrade, PREV_SSLGrade, logo, title, isDown, domainId) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer sqlStm.Close()
	var idServer int64
	err = sqlStm.QueryRow(infoDomain.Servers_Changed, infoDomain.Min_SSL_Grade, infoDomain.Previous_SSL_Grade, infoDomain.Logo, infoDomain.Title, infoDomain.Is_Down, idDomain).Scan(&idServer)

	if err != nil {
		log.Fatal(err)
	} else {
		openDB()
		for _, server := range infoDomain.Servers {
			sqlStm, err := db.Prepare("INSERT INTO SERVERITEM (address, ssl_grade, country, owner, serverId) VALUES ($1, $2, $3, $4, $5);")
			if err != nil {
				log.Fatal("Error inserting s: ", err)
			}
			defer sqlStm.Close()
			_, err = sqlStm.Exec(server.Address, server.SSL_Grade, server.Country, server.Owner, idServer)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	db.Close()
}


func getPreviousServersItems(name string) []ModelsAPI.ServerItem {
	openDB()
	stmt, err := db.Prepare(`SELECT SERVERITEM.address, SERVERITEM.SSL_Grade, SERVERITEM.country, SERVERITEM.owner 
							FROM DOMAIN, SERVER, SERVERITEM
							WHERE domainSearched = $1 AND domain.timeSearched < NOW() - INTERVAL '1 hour'
							AND SERVERITEM.serverId = SERVER.id AND SERVER.domainId = domain.id;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	servers := []ModelsAPI.ServerItem{}
	for rows.Next() {
		var server ModelsAPI.ServerItem
		err = rows.Scan(&server.Address, &server.SSL_Grade, &server.Country, &server.Owner)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}
	db.Close()
	return servers
}



func FindServers(id int64) []ModelsAPI.ServerItem {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT address, ssl_grade, country, owner
							FROM SERVER
							WHERE serverId = $1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	servers := []ModelsAPI.ServerItem{}

	for rows.Next() {
		var server ModelsAPI.ServerItem
		err = rows.Scan(&server.Address, &server.SSL_Grade, &server.Country, &server.Owner)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}
	db.Close()
	return servers
}
func PreviousGrade(tx *sql.Tx,name string) string {
	var ssl string
	log.Print(" \n ***PREVIOUS")
	if err := tx.QueryRow(
		"SELECT server.min_sslgrade " +
			"FROM DOMAIN, SERVER" +
			"WHERE domainsearched = $1 AND timesearched < NOW() - INTERVAL '1 hour' AND server.domainid=domain.id" +
			"ORDER BY timesearched DESC LIMIT 1", name).Scan(&ssl); err != nil {

		log.Print(" \n ***QUERY" + ssl)
		return ssl
	}

	return ssl
}

func getPrevious_SSL_Grade(name string) string {

	log.Print(" \n *FIND-PREVIOUS")
	db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/serverinformation?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	stmt, err := db.Prepare(`SELECT server.min_sslgrade FROM DOMAIN, SERVER
									WHERE domainsearched = $1 AND timesearched < NOW() - INTERVAL '1 hour'
									AND server.domainid=domain.id ORDER BY timesearched DESC LIMIT 1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	var ssl string
	log.Print(" \n *qUERY ROW:::")
	//err = stmt.QueryRow(name).Scan(&ssl)
	rows, err := stmt.Query(name)
	if err != nil {
		log.Print(" \n *Nill VACÍo")
		ssl = ""
	}
	defer rows.Close()
	for rows.Next() {
		var ssl2 string
		err = rows.Scan(&ssl2)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(" FOR SLL2->" + ssl2 + "<-")
		ssl = ssl2
	}

	db.Close()
	fmt.Print(" **** POr faaa SLL->" + ssl + "<-")
	return ssl

	/*
	ssl := ""
	log.Print(" \n *AFTER EXE -PREVIOUS")
	// Run a transfer in a transaction.
	err = crdb.ExecuteTx(context.Background(), db, nil, func(tx *sql.Tx) error {
		ssl = PreviousGrade(tx,name)
		fmt.Print(" \n *INTO FUNCTION->"+ssl+"<----")
		return err
	})
	fmt.Print(" \n *BEF EXE -PREVIOUS->"+ssl+"<----")

	return ssl*/
}


func ConectionDB() {
	log.Print(" Creating DATABASES ")
	/*db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/serverinformation?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}
	log.Print(" Creating accounts tab ")
	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS DOMAIN (" +
			"  id SERIAL," +
			" timeSearched TIMESTAMP WITHOUT TIME ZONE NOT NULL," +
			"domainSearched TEXT NOT NULL," +
			"  PRIMARY KEY (id))"); err != nil {
		log.Fatal(err)
	}

	log.Print(" Creating domain tab ")

	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS SERVER (" +
			" id SERIAL, serversChanged BOOLEAN NOT NULL," +
			" MIN_SSLGrade VARCHAR(3) NOT NULL, " +
			"PREV_SSLGrade VARCHAR(3) NOT NULL," +
			"logo TEXT NOT NULL,title TEXT NOT NULL, " +
			"isDown BOOLEAN NOT NULL," +
			"domainId INT NOT NULL," +
			" 	PRIMARY KEY (id)," +
			"	INDEX domain_ID (domainId ASC)," +
			" 	CONSTRAINT domainId   " +
			"	FOREIGN KEY (domainId)   " +
			"	REFERENCES domain (id)" +
			"	ON DELETE CASCADE" +
			"	ON UPDATE CASCADE)"); err != nil {
		log.Fatal(err)
	}
	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS SERVERITEM (" +
			" id SERIAL," +
			" address TEXT NOT NULL," +
			" SSL_Grade VARCHAR(3) NOT NULL, " +
			"country VARCHAR(5) NOT NULL," +
			" owner TEXT NOT NULL," +
			" serverId INT NOT NULL," +
			" PRIMARY KEY (id)," +
			"  INDEX server_ID (serverId ASC)," +
			"  CONSTRAINT serverId" +
			"    FOREIGN KEY (serverId)" +
			"  REFERENCES server (id)" +
			"    ON DELETE CASCADE" +
			"   ON UPDATE CASCADE)"); err != nil {
		log.Fatal(err)
	}*/
}