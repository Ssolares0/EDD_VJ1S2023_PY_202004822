import React, { useState, useEffect } from 'react';

export const Factura = () => {
    const [Fecha, setFecha] = useState(0);
    const [Empleado, setEmpleado] = useState(0);
    const [Usuario, setUsuario] = useState(0);
    const [Pago, setPago] = useState(0);

    const handleSubmit = async(e) => {
        e.preventDefault();
        await fetch('http://localhost:3001/GenerarFactura', {
            method: 'POST',
            body: JSON.stringify({
                Fecha: Fecha,
                Empleado: Empleado,
                Usuario: Usuario,
                Pago: Pago,
            }),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          .then(res => res.json())

          .then(data => {
            // Aquí puedes trabajar con la respuesta JSON recibida
            if (data.Admin ===true){
                window.open('http://localhost:3000/Admin','_self');
            } else if (data.Admin ===false){
                window.alert('Bienvenido Empleado ');
                window.open('http://localhost:3000/Empleado','_self');
            }
          })
          .catch(error => {
            console.log('Error:', error);
            // Manejo de errores
          });
        

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
                        <input type="date" className="login__input" placeholder="Fecha" required onChange={e => setFecha(e.target.value)}/>
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="number" className="login__input" placeholder="Empleado cobrador" required onChange={e =>setEmpleado(e.target.value)} />
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="number" className="login__input" placeholder="Usuario" required onChange={e =>setUsuario(e.target.value)} />
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="number" className="login__input" placeholder="Pago" required onChange={e =>setPago(e.target.value)} />
                    </div>
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