package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-almanac3/internal/server";"github.com/stockyard-dev/stockyard-almanac3/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="10140"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./almanac3-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("almanac3: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Almanac III — data dictionary and schema registry\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("almanac3: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
