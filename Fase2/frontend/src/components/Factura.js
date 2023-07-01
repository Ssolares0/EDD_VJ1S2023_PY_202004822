import React, { useState, useEffect } from 'react';

export const Factura = () => {

    const [Pago, setPago] = useState(0);
    const [Fecha, setFecha] = useState(0);
    const [Empleado, setEmpleado] = useState(0);
    const [Cliente, setCliente] = useState(0);


    const handleSubmit = async(e) => {
        e.preventDefault();
        await fetch('http://localhost:3001/GenerarFactura', {
            method: 'POST',
            body: JSON.stringify({
                Biller: Empleado,
                Customer:Cliente,
                Payment:Pago,
                Timestamp:Fecha
            }),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          .then(res => res.json())

          .then(data => {
            // Aquí puedes trabajar con la respuesta JSON recibida

          })
          .catch(error => {
            console.log('Error:', error);
            // Manejo de errores
          });
        

    };

    const validar = (data) =>{
        console.log(data)
        setFecha(data.Fecha)
        setEmpleado(data.Id_Empleado)
        setCliente(data.Id_Cliente)


    };

    const Reporte2 = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/MostrarDatosFactura',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    };


    const regresar = async(e) => {
        window.open('http://localhost:3000/Empleado', '_self');
    }

    
    

    return (
        
        <div className="container">
            
        <div className="screenFiltro">
            
            <div className="screen__content">
                
                
                <form onSubmit={handleSubmit} className="login">
                <h3 className='letra'>Generar Factura</h3>

                    
                    
                    <div className="login__field">
                        
                        <i className="login__icon fas fa-user"></i>
                        <input  className="login__input" placeholder="Fecha" value ={Fecha} />
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="number" className="login__input" placeholder="Empleado cobrador" value ={Empleado} />
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="number" className="login__input" placeholder="Usuario" value={Cliente} />
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="text" className="login__input" placeholder="Pago"  />
                    </div>
                    <button className="button login__submit"  value="mostrardata Pago"type ="button"onClick={Reporte2} > 
                        <span className="button__text">Mostrar Data</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="Realizar Pago"type ="submit"id="submit" >
                        <span className="button__text">Realizar Pago</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>		
                    <button className="button login__submit"  value="Realizar Pago"type ="button"id="regresar" onClick={regresar} >
                        <span className="button__text">Regresar</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>		
                </form>
               
            </div>
            <div className="screen__background">
                <span className="screen__background__shape screen__background__shape4"></span>
                <span className="screen__background__shape screen__background__shape3"></span>		
                <span className="screen__background__shape screen__background__shape2"></span>
                <span className="screen__background__shape screen__background__shape1"></span>
            </div>		
        </div>
    </div>
    );

};