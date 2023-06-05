const idInput = document.getElementById('clientId');

const btnEliminar = document.getElementById('eliminar_cliente_btn');

btnEliminar.addEventListener('click', (e) => {
    btnEliminar.disabled = true;
    $.ajax({
        url: '../eliminar_cliente/' + idInput.value,
        method: 'DELETE',
        data: {},
        success: function (response) {
            Toastify({
                text: "Cliente eliminado correctamente",
                duration: 2000,
                gravity: "top", // Position: "top", "bottom", or "center"
                positionLeft: false, // Position the toast on the left side
                backgroundColor: "linear-gradient(to right, #00b09b, #96c93d)", // Custom background color
                stopOnFocus: true, // Stop timer when the toast is hovered over
              }).showToast();
              setTimeout(function() {
                btnEliminar.disabled = false;
                window.location.href = "../ver_clientes";
              }, 2000);
        },
        error: function (error) {
            console.log(error);
            Toastify({
                text: error.responseText,
                duration: 3000,
                gravity: "top", // Position: "top", "bottom", or "center"
                positionLeft: false, // Position the toast on the left side
                backgroundColor: "linear-gradient(to right, #c44008, #f0113d)", // Custom background color
                stopOnFocus: true, // Stop timer when the toast is hovered over
              }).showToast();

              btnEliminar.disabled = false;
        }
    });
});