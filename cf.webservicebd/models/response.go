package models

import "net/http"
import "encoding/json"
import "fmt"

//Response es la estructura de la respuesta que vamos a enviar
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func createDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

func (response *Response) notFound() {
	response.Status = http.StatusNotFound
	response.Message = "No se encontro lo solicitado"
}

func SendNotFound(w http.ResponseWriter) {
	response := createDefaultResponse(w)
	response.notFound()
	response.send()
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := createDefaultResponse(w)
	response.unprocessableEntity()
	response.send()
}

func (response *Response) unprocessableEntity() {
	response.Status = http.StatusUnprocessableEntity
	response.Message = "No se pudo procesar la solicitud"
}

func (response *Response) noContent() {
	response.Status = http.StatusNoContent
	response.Message = "Se elimino el contenido"
}

func SendNotContent(w http.ResponseWriter) {
	response := createDefaultResponse(w)
	response.unprocessableEntity()
	response.send()
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := createDefaultResponse(w)
	response.Data = data
	response.send()
}

func (response *Response) send() {
	response.writer.Header().Set("Content-Type", response.contentType)
	response.writer.WriteHeader(response.Status)
	output, _ := json.Marshal(&response)
	fmt.Fprintf(response.writer, string(output))
}
