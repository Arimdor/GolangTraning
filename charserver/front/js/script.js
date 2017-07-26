$(document).ready(function () {
    var username
    var final_conexion
    console.log(final_conexion)
    $("#form_registro").on("submit", function (e) {
        e.preventDefault();
        username = $("input[name=username]").val()
        $.ajax({
            type: "POST",
            url: "http://localhost/validate",
            data: {
                "username": username
            },
            success: function (data) {
                result(data)
                console.log("ajax")
            }
        })
    })

    function result(data) {
        console.log("El servidor envio algo " + username)
        obj = JSON.parse(data)
        console.log(obj)
        if (obj.isvalid === true) {
            console.log("Se creo conexion")
            create_conexion()
        } else {
            alert("El nickname ya esta siendo utilizado")
        }
    }

    function create_conexion() {
        console.log("entro conexion")
        $("#registro").hide()
        $("#container_chat").show()
        var conexion = new WebSocket("ws://localhost/chat/" + username)
        final_conexion = conexion

        conexion.onopen = function (response) {
            conexion.onmessage = function (response) {
                console.log(response.data)
                val = $("#chat_area").val()
                $("#chat_area").val(val + "\n" + response.data)
            }
        }
    }
    $("#form_message").on("submit", function (e) {
        e.preventDefault();
        mensaje = $("#msg").val()
        final_conexion.send(mensaje)
        $("#msg").val("")
    })
});