package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Schema struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Database string `json:"database"`
	TableName string `json:"table_name"`
	Columns string `json:"columns"`
	Description string `json:"description"`
	Owner string `json:"owner"`
	Version int `json:"version"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"almanac3.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS schemas(id TEXT PRIMARY KEY,name TEXT NOT NULL,database TEXT DEFAULT '',table_name TEXT DEFAULT '',columns TEXT DEFAULT '[]',description TEXT DEFAULT '',owner TEXT DEFAULT '',version INTEGER DEFAULT 1,created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Schema)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO schemas(id,name,database,table_name,columns,description,owner,version,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Database,e.TableName,e.Columns,e.Description,e.Owner,e.Version,e.CreatedAt);return err}
func(d *DB)Get(id string)*Schema{var e Schema;if d.db.QueryRow(`SELECT id,name,database,table_name,columns,description,owner,version,created_at FROM schemas WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Database,&e.TableName,&e.Columns,&e.Description,&e.Owner,&e.Version,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Schema{rows,_:=d.db.Query(`SELECT id,name,database,table_name,columns,description,owner,version,created_at FROM schemas ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Schema;for rows.Next(){var e Schema;rows.Scan(&e.ID,&e.Name,&e.Database,&e.TableName,&e.Columns,&e.Description,&e.Owner,&e.Version,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Schema)error{_,err:=d.db.Exec(`UPDATE schemas SET name=?,database=?,table_name=?,columns=?,description=?,owner=?,version=? WHERE id=?`,e.Name,e.Database,e.TableName,e.Columns,e.Description,e.Owner,e.Version,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM schemas WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM schemas`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Schema{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (name LIKE ? OR description LIKE ?)"
        args=append(args,"%"+q+"%");args=append(args,"%"+q+"%");
    }
    rows,_:=d.db.Query(`SELECT id,name,database,table_name,columns,description,owner,version,created_at FROM schemas WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Schema;for rows.Next(){var e Schema;rows.Scan(&e.ID,&e.Name,&e.Database,&e.TableName,&e.Columns,&e.Description,&e.Owner,&e.Version,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    return m
}
