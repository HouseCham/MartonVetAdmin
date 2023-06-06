const nombreInput = document.getElementById('nombre');
const apellidoPInput = document.getElementById('apellidoP');
const apellidoMInput = document.getElementById('apellidoM');
const telefonoInput = document.getElementById('telefono');
const emailInput = document.getElementById('correo');
const calleInput = document.getElementById('calle');
const numeroExtInput = document.getElementById('numExt');
const numeroIntInput = document.getElementById('numInt');
const coloniaInput = document.getElementById('colonia');
const ciudadInput = document.getElementById('ciudad');
const estadoInput = document.getElementById('estado');
const codigoPostalInput = document.getElementById('codigoPostal');

const btnRegistrar = document.getElementById('registrar_cliente_btn');

btnRegistrar.addEventListener('click', (e) => {
    btnRegistrar.disabled = true;
    console.log(codigoPostalInput.value);
    $.ajax({
        url: '/insertar_cliente',
        method: 'POST',
        data: {
            nombre: nombreInput.value,
            apellidoP: apellidoPInput.value,
            apellidoM: apellidoMInput.value,
            telefono: telefonoInput.value,
            email: emailInput.value,
            calle: calleInput.value,
            numeroExt: numeroExtInput.value,
            numeroInt: numeroIntInput.value ? numeroIntInput.value : '',
            colonia: coloniaInput.value,
            codigoPostal: codigoPostalInput.value,
            ciudad: ciudadInput.value,
            estado: estadoInput.value
        },
        success: function (response) {
            Toastify({
                text: "Cliente registrado correctamente",
                duration: 2000,
                gravity: "top", // Position: "top", "bottom", or "center"
                positionLeft: false, // Position the toast on the left side
                backgroundColor: "linear-gradient(to right, #00b09b, #96c93d)", // Custom background color
                stopOnFocus: true, // Stop timer when the toast is hovered over
              }).showToast();
            setInputBlank();
            setTimeout(function() {
                btnRegistrar.disabled = false;
                window.location.href = "./ver_clientes";
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

              btnRegistrar.disabled = false;
        }
    });
});

function setInputBlank() {
    nombreInput.value = '';
    apellidoPInput.value = '';
    apellidoMInput.value = '';
    telefonoInput.value = '';
    emailInput.value = '';
    calleInput.value = '';
    numeroExtInput.value = '';
    numeroIntInput.value = '';
    coloniaInput.value = '';
    ciudadInput.value = '';
    estadoInput.value = '';
    codigoPostalInput.value = '';
}

  