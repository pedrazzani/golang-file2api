package handle

import(
	"html/template"
    "net/http"

	"github.com/pedrazzani/golang-file2api/usercase/product"
)

func ProductHandles(productService *product.Service) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
        listProducts,_ := productService.List()
        tmp := template.Must(template.ParseFiles("resource/template.html"))
        tmp.Execute(w, listProducts)
    })
}