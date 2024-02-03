import './assets/styles/App.css'
import { Router } from './router'
import { useNavigate } from 'react-router-dom';


function App() {
  
  const navigate = useNavigate();
  
  const handleDeveloperLogin = () => {
    navigate('/developeroptions');
  };
  
  return (
    <div className="App">
      <Router/>
      // This is your home page layout with login options
      <header className="App-header">
      <h1>Serverless Orchestrator</h1>
      <div className="login-container">
      <div className="login developer-login">
      <h2>Developer</h2>
      <input type="text" placeholder="Username" />
      <input type="password" placeholder="Password" />
      <button onClick={handleDeveloperLogin}>Login</button>
      </div>
      <div className="login user-login">
      <h2>Consumer</h2>
      <input type="text" placeholder="Username" />
      <input type="password" placeholder="Password" />
      <button>Login</button>
      </div>
      </div>
      </header>
    </div>
    );
  }
  
  export default App;