package main
import "fmt"

func main(){
	var anos int
	var dias int
	fmt.Println("Insira o numero de anos que deseja converter:")
	fmt.Scan(&anos)
	dias = anos*365
	if(dias<0){
		dias=dias*(-1)
	}
	fmt.Println("São",dias,"dias!!!")
}
