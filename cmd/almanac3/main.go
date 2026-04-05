package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-almanac3/internal/server";"github.com/stockyard-dev/stockyard-almanac3/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./almanac3-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("almanac3: %v",err)};defer db.Close();srv:=server.New(db,server.DefaultLimits())
fmt.Printf("\n  Almanac III — Self-hosted data dictionary and schema registry\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n  Questions? hello@stockyard.dev — I read every message\n\n",port,port)
log.Printf("almanac3: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
