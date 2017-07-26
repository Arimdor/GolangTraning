$(document).ready(function () {
    var username
    var final_conexion
    create_conexion()

    function create_conexion() {
        $("#registro").hide()
        $("#container_chat").show()
        var conexion = new WebSocket("ws://localhost/ws")
        final_conexion = conexion

        conexion.onopen = function (response) {
            conexion.onmessage = function (response) {
                var obj = jQuery.parseJSON(response.data);
                val = $("#chat_area").val()
                if(obj.sender === undefined){
                    $("#chat_area").val(val + "\n"+ obj.content)
                }
                else{
                    $("#chat_area").val(val + "\n" + obj.sender + ": " + obj.content)
                } 
            }
        }
    }
    $("#form_message").on("submit", function (e) {
        e.preventDefault();
        mensaje = $("#msg").val()
        if (mensaje == "") {
            return
        }
        final_conexion.send(mensaje)
        $("#msg").val("")
    })
});