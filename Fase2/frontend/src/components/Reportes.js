import React, {useState, useEffect} from 'react';

export const Reportes = () => {
    const [imagen, setImagen] = useState('https://yakurefu.com/wp-//2020/02/Chi_by_wallabby.jpg');
    
    const validar = (data) =>{
        console.log(data)
        setImagen(data.ImageBase64)
        /*if (data.mensaje === "OK"){
            swal({
                title: "Session Iniciada",
                text: "Datos Correctos",
                icon: "success",
                button: "aceptar"
            }).then(respuesta => {
                if(respuesta){
                    console.log(respuesta);
                    localStorage.setItem('current',respuesta.data);
                    window.open("/inicio","_self");
                }
            })
            
        }else{
            swal({
                title: "Error en Credenciales",
                text: "Su usuario o contraseña son incorrectos",
                icon: "error",
                timer: "4000",
                buttons: false
            })
        }*/
    }

    const pedirReporte = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/ReporteTree',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }
    

    

    return (
        
        <div className="container">
            <button className="button login__submit"  value="Ver reporte" onClick={pedirReporte}>
                        <span className="button__text">Enviar Datos</span>
                        <i className="button__icon fas fa-chevron-right"></i>
            </button>	
            <img src={imagen}  />
       
        </div>
        
    );

};