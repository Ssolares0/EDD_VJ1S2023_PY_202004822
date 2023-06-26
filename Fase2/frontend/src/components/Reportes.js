import React, {useState, useEffect} from 'react';

export const Reportes = () => {
    const [imagen, setImagen] = useState();
    
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
        
        <div className="form-signin">
        <div className="text-center">
            <form className="card card-body">
                <h1 className="h3 mb-3 fw-normal">Arbol AVL</h1>
                <label htmlFor="inputEmail" className="visually-hidden">Valor</label>
                
        
            
                <br/>
                <button className="w-100 btn btn-lg btn-primary" onClick={pedirReporte}>Ver Reporte</button>
                <br/>
                <img src={imagen} width="250" height="250" alt='some value' />
            </form>
        </div>
      </div>
        
    );

};