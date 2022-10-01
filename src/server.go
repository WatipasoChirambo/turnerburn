package main

import (
	"fmt"
	"log"
	"net/http"
)

var pumpState = false;
const MAX_WATER_LEVEL=10;
var water_level = 0;


func pumpHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm();err != nil{
		fmt.Fprintf(w, "ParseForm() err:%v",err)
		return
	}
	fmt.Fprintf(w, "Pump State\n")
	on := r.FormValue("pump");
	fmt.Fprintf(w, "Pump = %s\n", on)
	if on == "on" && MAX_WATER_LEVEL != water_level {
		for i := 0; i < MAX_WATER_LEVEL; i++ {
			fmt.Println("Watering Plants")
			water_level++;
		};
		if water_level==MAX_WATER_LEVEL {
				pumpState= false;
				fmt.Println("Enough Water")
				fmt.Println("Pump Closed")
			};
	}else{
		return
	}
	fmt.Fprintf(w, "Water Level = %d\n", water_level)

}

func main(){
	var fileServer = http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/pump",pumpHandler)

	fmt.Printf("Starting server at port 8080\n");

	if err := http.ListenAndServe(":8000",nil); err != nil {
		log.Fatal(err)
	}
}
