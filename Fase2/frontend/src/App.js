import logo from './logo.svg';
import './css/Style.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { Login } from './components/Login';
import { Admin } from './components/Admin';

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
      </Routes>  
    </Router>

  );
}

export default App;
