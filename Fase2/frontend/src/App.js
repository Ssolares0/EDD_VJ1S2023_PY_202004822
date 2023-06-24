import logo from './logo.svg';
import './css/Style.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { Login } from './components/Login';
import { Admin } from './components/Admin';
import { Empleado } from './components/Empleado';
import { Reportes } from './components/Reportes';

function App() {
  /*return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );*/
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/Admin" element={<Admin />}></Route>
        <Route path="/Empleado" element={<Empleado/>}></Route>
        <Route path="/Reportes" element={<Reportes />} />
      </Routes>  
    </Router>

  );
}

export default App;
