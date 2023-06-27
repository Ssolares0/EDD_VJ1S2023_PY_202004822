import React, { useState } from 'react';
export const Filtros = () => {
    const [Filtro, setTipo] = useState(0);
    const handleSubmit = async(e) => {
        
        e.preventDefault();
        
        
    }

    const EspejoX = async(e) => {
        e.preventDefault();

        await fetch('http://localhost:3001/Filtros', {
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify({
                Filtro: "EspejoX"

            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'

            }
        })
        
        
    }
    const EspejoY = async(e) => {
        e.preventDefault();

        await fetch('http://localhost:3001/Filtros', {
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify({
                Filtro: "EspejoY"

            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'

            }
        })
    }
    const Negativo = async(e) => {
        e.preventDefault();

        await fetch('http://localhost:3001/Filtros', {
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify({
                Filtro: "Negativo"

            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'

            }
        })
        
        
    }
    const EscalaGrises = async(e) => {
        e.preventDefault();

        await fetch('http://localhost:3001/Filtros', {
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify({
                Filtro: "EscalaGrises"

            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'

            }
        })
    }

    const Regresar = async(e) => {
        window.open('http://localhost:3000/Empleado', '_self');
        
    }

    return (
        <div className="container">
        <div className="screenFiltro">
            <div className="screen__content">
            
                
                
                <form  onSubmit={handleSubmit} className="login">
                <h3 className='letra'>Aplicar Filtros</h3>
                    <button className="button login__submit"  value="negativo" type='button' onClick={Negativo}>
                        <span className="button__text">Aplicar negativo</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="EscalaGrises" type ="button"onClick={EscalaGrises} >
                        <span className="button__text">Aplicar Escala de grises</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="EspejoX" type ="button"onClick={EspejoX}>
                        <span className="button__text">Aplicar Espejo X </span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="EspejoY"type ="button" onClick={EspejoY}>
                        <span className="button__text">Aplicar Espejo Y </span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit" type ="button" value="ambosEspejos">
                        <span className="button__text">Aplicar ambos espejos </span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit" type ="button" value="GenerarCon">
                        <span className="button__text">Generar imagen con</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>	
                    <button className="button login__submit"  value="Regresar" type ="button" onClick={Regresar}>
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