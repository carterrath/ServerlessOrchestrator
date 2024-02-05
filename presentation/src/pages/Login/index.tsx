import './../../assets/styles/Login.css';
import { useNavigate } from "react-router-dom";
export function Login() {

  const navigate = useNavigate();
  
  const handleLogin = () => {
    navigate('/DeveloperOptions');
  };

  return (
    <div className="App">
      <header className="App-header">
      <h1>Serverless Orchestrator</h1>
      <div className="login-container">
      <div className="login developer-login">
      <h2>Developer</h2>
      <input type="text" placeholder="Username" />
      <input type="password" placeholder="Password" />
      <button onClick={handleLogin}>Login</button>
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