package Controllers

import (
	"GoProject/ModelsAPI"
	"database/sql"
	"log"
)

var db *sql.DB

func FindServersRecords() ModelsAPI.ServersConsulted {
	//Call to DATA BASE.
	//Query: FROM DOMAIN
	//		ORDER BY MAX(consulted_time)
	//		LIMIT 100;
	stmt, err := db.Prepare(`SELECT DISTINCT ON (host) host, id
							FROM DOMAIN
							GROUP BY host, id
							ORDER BY host, MAX(consulted_time), id
							LIMIT 100;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	record := ModelsAPI.ServersConsulted{}
	items := []ModelsAPI.Item{}

	for rows.Next() {
		var item ModelsAPI.Item
		var id int64
		err = rows.Scan(&item.Domain, &id)
		if err != nil {
			log.Fatal(err)
		}
		//	item.Info = FindHost(id)
		items = append(items, item)
	}
	record.Items = items
	return record
}
