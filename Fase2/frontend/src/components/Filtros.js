
export const Filtros = () => {
    const handleSubmit = async(e) => {
        e.preventDefault();
        
        
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
                    <button className="button login__submit"  value="negativo">
                        <span className="button__text">Aplicar negativo</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="EscalaGrises">
                        <span className="button__text">Aplicar Escala de grises</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="EspejoX">
                        <span className="button__text">Aplicar Espejo X </span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="ambosEspejos">
                        <span className="button__text">Aplicar ambos espejos </span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>
                    <button className="button login__submit"  value="GenerarCon">
                        <span className="button__text">Generar imagen con</span>
                        <i className="button__icon fas fa-chevron-right"></i>
                    </button>	
                    <button className="button login__submit"  value="Regresar" onClick={Regresar}>
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