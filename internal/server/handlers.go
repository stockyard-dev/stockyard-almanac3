package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-almanac3/internal/store")
func(s *Server)handleListTables(w http.ResponseWriter,r *http.Request){db:=r.URL.Query().Get("database");list,_:=s.db.ListTables(db);if list==nil{list=[]store.Table{}};writeJSON(w,200,list)}
func(s *Server)handleCreateTable(w http.ResponseWriter,r *http.Request){var t store.Table;json.NewDecoder(r.Body).Decode(&t);if t.Name==""||t.Database==""{writeError(w,400,"name and database required");return};s.db.CreateTable(&t);writeJSON(w,201,t)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleListColumns(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListColumns(id);if list==nil{list=[]store.Column{}};writeJSON(w,200,list)}
func(s *Server)handleAddColumn(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var c store.Column;json.NewDecoder(r.Body).Decode(&c);c.TableID=id;if c.Name==""{writeError(w,400,"name required");return};s.db.AddColumn(&c);writeJSON(w,201,c)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
