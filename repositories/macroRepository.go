package repositories

import (
	"database/sql"

	"github.com/senowijayanto/apis/infrastructures"
	"github.com/senowijayanto/apis/models"
	log "github.com/sirupsen/logrus"
)

type IMacroRepository interface {
	GetListMacro() (models.MacroResponse, error)
}

type MacroRepository struct {
	DB infrastructures.IMYSQLConnection
}

func (r *MacroRepository) GetListMacro() (macroResponse models.MacroResponse, err error) {
	db := r.DB.MarketDataRead()
	defer db.Close()

	rows, err := db.Query(`SELECT type_name, economy_type, economy_code, economy_name, economy_last, economy_previous, (economy_last - economy_previous) AS ch,
												economy_change, economy_3month_avg FROM macro_economy_summary a LEFT JOIN macro_economy_type b ON b.type_id = a.economy_type
												WHERE a.economy_status = 1 ORDER BY economy_type,economy_code ASC`)
	if err == sql.ErrNoRows {
		err = nil
	}
	defer rows.Close()

	for rows.Next() {
		var macro models.Macro
		if err := rows.Scan(
			&macro.TypeTable,
			&macro.IdTable,
			&macro.Code,
			&macro.Name,
			&macro.Last,
			&macro.Previous,
			&macro.Change,
			&macro.ChangeP,
			&macro.ThreeMonthAVG,
		); err != nil {
			log.WithFields(log.Fields{
				"event": "get_list_macro",
			}).Error(err)
		}
		macroResponse.Data = append(macroResponse.Data, macro)
	}

	rows, err = db.Query(`SELECT COUNT(*) FROM macro_economy_summary a LEFT JOIN macro_economy_type b ON b.type_id = a.economy_type WHERE a.economy_status = 1 ORDER BY economy_type,economy_code ASC`)
	if err == sql.ErrNoRows {
		err = nil
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			log.WithFields(log.Fields{
				"event": "get_list_macro",
			}).Error(err)
		}
		macroResponse.Total = count
	}

	if err := rows.Err(); err != nil {
		log.WithFields(log.Fields{
			"event": "get_list_macro",
		}).Error(err)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"event": "get_list_macro",
		}).Error(err)
	}

	return
}
