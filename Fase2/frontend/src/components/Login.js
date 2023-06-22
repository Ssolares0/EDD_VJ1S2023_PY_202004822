import React, { useState, useEffect } from 'react';

export const Login = () => {
    const [Username, setUsername] = useState(0);
    const [Password, setPassword] = useState(0);

    const handleSubmit = async(e) => {
        e.preventDefault();
        await fetch('http://localhost:3001/Login', {
            method: 'POST',
            body: JSON.stringify({
              Username: Username,
              Password: Password,
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
                window.alert('Pare');
            }
          })
          .catch(error => {
            console.log('Error:', error);
            // Manejo de errores
          });
        

    };
    

    return (
        
        <div className="container">
            
        <div className="screen">
            
            <div className="screen__content">
                
                
                <form onSubmit={handleSubmit} className="login">
                <h3 className='letra'>Log In</h3>

                    
                    
                    <div className="login__field">
                        
                        <i className="login__icon fas fa-user"></i>
                        <input type="text" className="login__input" placeholder="User name / Email" required onChange={e => setUsername(e.target.value)}/>
                    </div>
                    <div className="login__field">
                        <i className="login__icon fas fa-lock"></i>
                        <input type="password" className="login__input" placeholder="Password" required onChange={e =>setPassword(e.target.value)} />
                    </div>
                    <button className="button login__submit"  value="Iniciar sesión"type ="submit"id="submit" >
                        <span className="button__text">Log In Now</span>
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